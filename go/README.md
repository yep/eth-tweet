# EtherTweet API Server

This directory contains native [go](https://golang.org) bindings for the EtherTweet [Solidity](https://solidity.readthedocs.org) code.

Using these bindings, a simple JSON HTTP API is implemented.

See the Ethereum Wiki on page [Native-DApps](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts) for more info.


## Setup on OSX

Install `go`, for example using the [homebrew package manager](http://brew.sh):

    brew install go
    
[Install the solidity compiler](http://www.ethdocs.org/en/latest/ethereum-clients/cpp-ethereum/index.html#installing-and-building). The `solc` command has to be available below.

Get the source code of `EtherTweet`. Set `GOPATH` environment variable. Change to subdirectory `go` of the `eth-tweet` directory. Run update script. The commands are:

    git clone https://github.com/yep/eth-tweet.git
    cd eth-tweet/go
    export GOPATH=$(pwd)
    ./update-go-bindings.sh

Assuming you have set the GOPATH as described above, you can run the `ethertweet` binary now by typing:

     ethertweet
     
Server should then be running at `http://localhost:8080`

Run some tests on a simulated blockchain:

    go test -v ethertweet.net/ethertweet
    
Output should be similar to:

    TweetAccount address: 0x5dd637cb09774c3e946966db2dace2b73bd21082
    TweetAccount getNumberOfTweets: 1
    TweetAccount getLatestTweet: hello world

With the GOPATH set as above and after having run the setup script once, you can recompile the `ethertweet` binary like this:

    go install ethertweet.net/ethertweet


## Setup in a Virtual Machine

Setup on a fresh Ubuntu Server 16.04 Virtual Machine (VM). You can have a look at the commands used for install in the `Vagrantfile`.

 - Install [Virtualbox](http://virtualbox.org) and [Vagrant](http://vagrantup.com)
 - Run `vagrant up`


## Blockchain rsync

Downloading the blockchain takes a long time, even using `geth --fast`. Copy existing blockchain data from a linux host with rsync over ssh:

    rsync -avz -e 'ssh -i /home/USER/.ssh/KEY' /PATH/.ethereum/ USER@HOST:/PATH/.ethereum

