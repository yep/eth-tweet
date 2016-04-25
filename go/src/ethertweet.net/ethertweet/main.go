package main

// json http api for ether tweet
// see https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"ethertweet.net/ethertweet/lib"

	"fmt"
	"log"
	"net/http"
	"os/user"
	"path/filepath"
	"encoding/json"
	"errors"
	"strconv"
	"math/big"
	"runtime"
)

var rpcBackend = createRpcBackend()
var tweetRegistry = createRegistry()

func main() {
	router := httprouter.New()
	router.GET("/status", statusHandler)
	router.GET("/registry", registryHandler)
	router.GET("/registry/:accountId", registryAccountIdHandler)
	router.GET("/account/:accountName", accountHandler)
	router.GET("/account/:accountName/:tweetNumber", accountTweetNumberHandler)
	router.GET("/address/:accountAddress", addressHandler)
	router.GET("/address/:accountAddress/:tweetNumber", addressTweetNumberHandler)
	router.GET("/", indexHandler)

	router.NotFound = http.HandlerFunc(notFound)
	router.MethodNotAllowed = http.HandlerFunc(notAllowed)
	router.PanicHandler = panicHandler

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// --- setup

// Get an IPC connection to geth
func createRpcBackend() bind.ContractBackend {
	conn, err := rpc.NewIPCClient(getIPCSocketPath())
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}

	return backends.NewRPCBackend(conn)
}

// Get Path to IPC socket depending on OS
func getIPCSocketPath() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get home directory of current user: %v", err)
	}

	path := ""
	if runtime.GOOS == "darwin" {
		path = filepath.Join(currentUser.HomeDir, "Library/Ethereum/geth.ipc")
	} else if runtime.GOOS == "linux" {
		path = filepath.Join(currentUser.HomeDir, ".ethereum/geth.ipc")
	} else {
		log.Fatal("unsupported operating system")
	}

	return path
}

// EtherTweet default account registry
func createRegistry() *ethertweet.TweetRegistry {
	tweetRegistry, err := ethertweet.NewTweetRegistry(common.HexToAddress("0xe0f278b72097e563b09d7dc94c6f75aff5e83298"), rpcBackend)
	if err != nil {
		log.Fatalf("Failed to access TweetRegistry: %v", err)
	}

	return tweetRegistry
}

// --- http router errors

func notFound(responseWriter http.ResponseWriter, request *http.Request) {
	printJsonError(responseWriter, request, 404, "Not found")
}

func notAllowed(responseWriter http.ResponseWriter, request *http.Request) {
	printJsonError(responseWriter, request, 405, "Not allowed")
}

func panicHandler(responseWriter http.ResponseWriter, request *http.Request, _ interface{}) {
	printJsonError(responseWriter, request, 500, "Internal server error")
}

// --- http router

func indexHandler(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	printJsonKeyValueResponse(responseWriter, request, "EtherTweet", "API")
}

func statusHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	// check if ipc connection is available
	_, err := tweetRegistry.GetNumberOfAccounts(nil)
	if err != nil {
		printJsonError(responseWriter, request, 503, "Service unavailable")
		return
	}

	printJsonKeyValueResponse(responseWriter, request, "Status", "Ok")
}

func registryHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	numberOfAccounts, err := tweetRegistry.GetNumberOfAccounts(nil)
	if err != nil {
		printJsonError(responseWriter, request, 500, "Failed to get number of accounts")
		return
	}

	printJsonKeyValueResponse(responseWriter, request, "NumberOfAccounts", numberOfAccounts.String())
}

func registryAccountIdHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	accountId, err := strconv.Atoi(parameters.ByName("accountId"))
	if err != nil {
		printJsonError(responseWriter, request, 400, "Invalid account id")
		return
	}

	accountAddress, err := tweetRegistry.GetAddressOfId(nil, big.NewInt(int64(accountId)))
	if err != nil {
		printJsonError(responseWriter, request, 500, "Failed to get address of id")
		return
	}

	accountName, err := tweetRegistry.GetNameOfAddress(nil, accountAddress)
	if err != nil {
		printJsonError(responseWriter, request, 500, "Failed to get name of address")
		return
	}
	if accountName == "" {
		printJsonError(responseWriter, request, 404, "Account not found")
		return
	}

	type result struct {
		AccountName string
		AccountAddress string
	}
	printJsonResponse(responseWriter, request, result{AccountName: accountName, AccountAddress:accountAddress.Hex()})
}

func accountHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	getLatestTweet(responseWriter, request, parameters.ByName("accountName"))
}

func accountTweetNumberHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	getTweet(responseWriter, request, parameters.ByName("accountName"), parameters.ByName("tweetNumber"))
}

func addressHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	accountName, err := tweetRegistry.GetNameOfAddress(nil, common.HexToAddress(parameters.ByName("accountAddress")))
	if err != nil {
		printJsonError(responseWriter, request, 500, "Failed to get account name of address")
		return
	}
	if accountName == "" {
		printJsonError(responseWriter, request, 404, "Account name not found")
		return
	}

	getLatestTweet(responseWriter, request, accountName)
}

func addressTweetNumberHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	accountName, err := tweetRegistry.GetNameOfAddress(nil, common.HexToAddress(parameters.ByName("accountAddress")))
	if err != nil {
		printJsonError(responseWriter, request, 500, "Failed to get account name of address")
		return
	}
	if accountName == "" {
		printJsonError(responseWriter, request, 404, "Account name not found")
		return
	}

	getTweet(responseWriter, request, accountName, parameters.ByName("tweetNumber"))
}

// --- helper functions

func getLatestTweet(responseWriter http.ResponseWriter, request *http.Request, accountName string) {
	tweetAccount, err := getAccountByName(accountName)
	if err != nil {
		printJsonError(responseWriter, request, 500, err.Error())
		return
	}

	getLatestTweetResult, err := tweetAccount.GetLatestTweet(nil)
	if err != nil {
		printJsonError(responseWriter, request, 500, "Failed to get latest tweet")
		return
	}

	printJsonResponse(responseWriter, request, getLatestTweetResult)
}

func getTweet(responseWriter http.ResponseWriter, request *http.Request, accountNameString string, tweetNumberString string) {
	tweetAccount, err := getAccountByName(accountNameString)
	if err != nil {
		printJsonError(responseWriter, request, 500, err.Error())
		return
	}

	tweetNumber, err := strconv.Atoi(tweetNumberString)
	if err != nil {
		printJsonError(responseWriter, request, 400, "Invalid tweet number")
		return
	}

	getTweetResult, err := tweetAccount.GetTweet(nil, big.NewInt(int64(tweetNumber)))
	if err != nil {
		printJsonError(responseWriter, request, 500, "Failed to get tweet")
		return
	}

	if getTweetResult.TweetString == "" && getTweetResult.Timestamp.Int64() == 0 {
		printJsonError(responseWriter, request, 404, "Tweet not found")
		return
	}

	printJsonResponse(responseWriter, request, getTweetResult)
}

func getAccountByName(accountName string) (*ethertweet.TweetAccount, error) {
	accountAddress, err := tweetRegistry.GetAddressOfName(nil, accountName)

	if err != nil {
		return nil, errors.New("Failed to get address of account name")
	}

	tweetAccount, err := ethertweet.NewTweetAccount(accountAddress, rpcBackend)
	if err != nil {
		return nil, errors.New("Failed to get TweetAccount")
	}

	return tweetAccount, nil
}

// --- response

func printJsonResponse(responseWriter http.ResponseWriter, request *http.Request, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		log.Println("500 " + request.URL.String())
		printJsonError(responseWriter, request, 500, "Json marshal failed")
	} else {
		log.Println("200 " + request.URL.String())
		fmt.Fprintf(responseWriter, "%s", json)
	}
}

func printJsonKeyValueResponse(responseWriter http.ResponseWriter, request *http.Request, key string, value string) {
	log.Println("200 " + request.URL.String())
	fmt.Fprintf(responseWriter, keyValueJson(key, value))
}

func printJsonError(responseWriter http.ResponseWriter, request *http.Request, statusCode int, errorMessage string) {
	log.Println(fmt.Sprintf("%d %s", statusCode, request.URL.String()))
	responseWriter.WriteHeader(statusCode)
	fmt.Fprintf(responseWriter, keyValueJson("Error", errorMessage))
}

func keyValueJson(key string, value string) string {
	return fmt.Sprintf(`{"` + key + `":"` + value + `"}`)
}
