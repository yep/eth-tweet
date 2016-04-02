#!/bin/bash -x

# update native go bindings
# see https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts

# change to directory this script is located at and set gopath to it
cd "$(dirname "$0")"
GOPATH=$(pwd)

go get github.com/tools/godep

# do not checkout complete git history and only the correct branch for a faster clone operation
mkdir -p src/github.com/ethereum
git clone --depth=1 --single-branch --branch develop https://github.com/ethereum/go-ethereum.git src/github.com/ethereum/go-ethereum

cd src/github.com/ethereum/go-ethereum/
godep go install ./cmd/abigen

cd $GOPATH
PACKAGE=ethertweet
abigen --sol ../TweetAccount.sol --pkg ${PACKAGE} --out src/ethertweet.net/ethertweet/lib/TweetAccount.go
abigen --sol ../TweetRegistry.sol --pkg ${PACKAGE} --out src/ethertweet.net/ethertweet/lib/TweetRegistry.go

# build
cd $GOPATH/src/ethertweet.net/ethertweet
go get -v .
go install

set +x
echo
echo successfully updated go bindings
echo ethertweet binary is available at:
echo $GOPATH/bin/ethertweet