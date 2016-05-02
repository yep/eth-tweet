#!/bin/bash -x

# update native go bindings
# see https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts

# change to directory this script is located at and set gopath to it
cd "$(dirname "$0")"
export GOPATH=$(pwd)

go get github.com/tools/godep

# do not checkout complete git history and only the correct branch for a faster clone operation
mkdir -p src/github.com/ethereum
git clone --depth=1 --single-branch --branch release/1.4 https://github.com/ethereum/go-ethereum.git src/github.com/ethereum/go-ethereum

cd src/github.com/ethereum/go-ethereum/
$GOPATH/bin/godep go install ./cmd/abigen
$GOPATH/bin/godep go install ./cmd/geth

cd $GOPATH
$GOPATH/bin/abigen --sol ../TweetAccount.sol --pkg ethertweet --out src/ethertweet.net/ethertweet/lib/TweetAccount.go
$GOPATH/bin/abigen --sol ../TweetRegistry.sol --pkg ethertweet --out src/ethertweet.net/ethertweet/lib/TweetRegistry.go

# build
cd $GOPATH/src/ethertweet.net/ethertweet
go get -v .
go install

set +x
echo
echo ethertweet binary is at:
echo $GOPATH/bin/ethertweet
