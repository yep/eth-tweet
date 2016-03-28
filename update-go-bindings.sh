#!/bin/sh

# update native go bindings
# see https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts

abigen --sol TweetAccount.sol --pkg main --out go/src/ethertweet.net/ethertweet/TweetAccount.go
abigen --sol TweetRegistry.sol --pkg main --out go/src/ethertweet.net/ethertweet/TweetRegistry.go
