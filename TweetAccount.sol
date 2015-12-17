/*
decentralized microblogging 
Copyright (C) 2015 Jahn Bertsch

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation in version 3.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// "class" TweetAccount
contract TweetAccount {
	
	// data structure of a single tweet
	struct Tweet {
		uint timestamp;
		string tweetString;
	}
	
	// "array" of all tweets of this account: maps the tweet id to the actual tweet
	mapping (uint => Tweet) _tweets;
	
	// total number of tweets in the above _tweets mapping
	uint _numberOfTweets;
	
	// "owner" of this account: only admin is allowed to tweet
	address _adminAddress;
	
	// constructor
	function TweetAccount() {
		_numberOfTweets = 0;
		_adminAddress = msg.sender;
	}
	
	// returns true if caller of function ("sender") is admin
	function isAdmin() constant returns (bool isAdmin) {
		return msg.sender == _adminAddress;
	}
	
	// create new tweet
	function tweet(string tweetString) constant returns (int result) {
		if (!isAdmin()) {
			// only owner is allowed to create tweets for this account
			result = -1;
		} else if (bytes(tweetString).length > 160) {
			// tweet contains more than 160 bytes
			result = -2;
		} else {
			_tweets[_numberOfTweets].timestamp = now;
			_tweets[_numberOfTweets].tweetString = tweetString;
			_numberOfTweets++;
			result = 0; // success
		}
	}
	
	function getTweet(uint tweetId) constant returns (string tweetString, uint timestamp) {
		// returns two values
		tweetString = _tweets[tweetId].tweetString;
		timestamp = _tweets[tweetId].timestamp;
	}
	
	function getLatestTweet() constant returns (string tweetString, uint timestamp, uint numberOfTweets) {
		// returns three values
		tweetString = _tweets[_numberOfTweets - 1].tweetString;
		timestamp = _tweets[_numberOfTweets - 1].timestamp;
		numberOfTweets = _numberOfTweets;
	}
	
	function getOwnerAddress() constant returns (address adminAddress) {
		return _adminAddress;
	}
	
	function getNumberOfTweets() constant returns (uint numberOfTweets) {
		return _numberOfTweets;
	}

	// other users can send donations to your account: use this function for donation withdrawal
	function adminRetrieveDonations(address receiver) {
		if (isAdmin()) {
			receiver.send(this.balance);
		}
	}
	
	function adminDeleteAccount() {
		if (isAdmin()) {
			suicide(_adminAddress); // this is a predefined function, it deletes the contract and returns all funds to the owner's address
		}
	}	
}
