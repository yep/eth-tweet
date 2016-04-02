# Go Bindings for EtherTweet

Generate native [go](https://golang.org) bindings for the `EtherTweet` [Solidity](https://solidity.readthedocs.org) code.

See the Ethereum Wiki on page [Native-DApps](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts) for more info.

## Setup

Install `go`. On OSX, for example using the [homebrew package manager](http://brew.sh):

    brew install go

Get the source code of `EtherTweet`. Set `GOPATH` environment variable. Change to subdirectory `go` of the `eth-tweet` directory. Run update script. The commands are:

    git clone https://github.com/yep/eth-tweet.git
    cd eth-tweet/go
    export GOPATH=$(pwd)
    ./update-go-bindings.sh

The `ethertweet` binary should now be available at `eth-tweet/go/bin/ethertweet`.

Assuming you have set the GOPATH as described above, you can run the `ethertweet` binary now by typing:

     ethertweet

It should run some tests on a simulated blockchain with an output similar to:

    TweetAccount address: 0xcf7309e18faa8c166019496de2a831217df3ad1c
    TweetAccount getNumberOfTweets: 1
    TweetAccount getLatestTweet: hello world

With the GOPATH set as above and after having run the setup script once, you can recompile the `ethertweet` binary like this later:

    go install ethertweet.net/ethertweet

