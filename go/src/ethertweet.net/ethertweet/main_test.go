package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"ethertweet.net/ethertweet/lib"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestAverage(t *testing.T) {
	// Generate a new random account and a funded simulator
	key := crypto.NewKey(rand.Reader)
	keyBytes, err := key.MarshalJSON()
	fmt.Println("private key: " + string(keyBytes))

	simulator := backends.NewSimulatedBackend(core.GenesisAccount{key.Address, big.NewInt(10000000000)})

	// Convert the tester key to an authorized transactor for ease of use
	auth := bind.NewKeyedTransactor(key)

	// Deploy a contract on the simulated blockchain
	tweetAccountAddress, _, tweetAccount, err := ethertweet.DeployTweetAccount(auth, simulator)
	if err != nil {
		t.Error("Failed to deploy TweetAccount contract: %v", err)
	}
	fmt.Printf("TweetAccount address: 0x%x\n", tweetAccountAddress)

	// Wrap the contract instance into a session
	tweetAccountSession := &ethertweet.TweetAccountSession{
		Contract: tweetAccount,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(3141592),
		},
	}

	_, err = tweetAccountSession.Tweet("hello world")
	if err != nil {
		t.Error("Failed to tweet: %v", err)
	}

	simulator.Commit()

	numberOfTweets, err := tweetAccount.GetNumberOfTweets(nil)
	if err != nil {
		t.Error("Failed to get number of tweets: %v", err)
	}
	fmt.Println("TweetAccount getNumberOfTweets:", numberOfTweets)

	getLatestTweetResult, err := tweetAccount.GetLatestTweet(nil)
	if err != nil {
		t.Error("Failed to read tweet: %v", err)
	}
	fmt.Println("TweetAccount getLatestTweet:", getLatestTweetResult.TweetString)
}