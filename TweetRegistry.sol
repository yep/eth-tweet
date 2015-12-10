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

// "class" TweetRegistry
contract TweetRegistry {
	
	// mappings to look up account names, account ids and addresses
	mapping (address => string) _addressToAccountName;
	mapping (uint => address) _accountIdToAccountAddress;
	mapping (string => address) _accountNameToAddress;
	
	// might be interesting to see how many people use the system
	uint _numberOfAccounts;
	
	// owner
	address _registryAdmin;
	
	// allowed to administrate accounts only, not everything
	address _accountAdmin;
	
	// if a newer version of this registry is available, force users to use it
	bool _registrationDisabled;

	function TweetRegistry() {
		_registryAdmin = msg.sender;
		_accountAdmin = msg.sender; // can be changed later
		_numberOfAccounts = 0;
		_registrationDisabled = false;
	}

	function register(string name, address accountAddress) returns (int result) {
		if (_accountNameToAddress[name] != address(0)) {
			// name already taken
			result = -1;
		} else if (bytes(_addressToAccountName[accountAddress]).length != 0) {
			// account address is already registered
			result = -2;
		} else if (bytes(name).length >= 64) {
			// name too long
			result = -3;
		} else if (_registrationDisabled){
			// registry is disabled because a newer version is available
			result = -4;
		} else {
			_addressToAccountName[accountAddress] = name;
			_accountNameToAddress[name] = accountAddress;
			_accountIdToAccountAddress[_numberOfAccounts] = accountAddress;
			_numberOfAccounts++;
			result = 0; // success
		}
	}
	
	function getNumberOfAccounts() constant returns (uint numberOfAccounts) {
		numberOfAccounts = _numberOfAccounts;
	}

	function getAddressOfName(string name) constant returns (address addr) {
		addr = _accountNameToAddress[name];
	}

	function getNameOfAddress(address addr) constant returns (string name) {
		name = _addressToAccountName[addr];
	}
	
	function getAddressOfId(uint id) constant returns (address addr) {
		addr = _accountIdToAccountAddress[id];
	}

	function unregister() returns (string unregisteredAccountName) {
		unregisteredAccountName = _addressToAccountName[msg.sender];
		_addressToAccountName[msg.sender] = "";
		_accountNameToAddress[unregisteredAccountName] = address(0);
		// _accountIdToAccountAddress is never deleted on purpose
	}
	
	function adminUnregister(string name) {
		if (msg.sender == _registryAdmin || msg.sender == _accountAdmin) {
			address addr = _accountNameToAddress[name];
			_addressToAccountName[addr] = "";
			_accountNameToAddress[name] = address(0);
			// _accountIdToAccountAddress is never deleted on purpose
		}
	}
	
	function adminSetRegistrationDisabled(bool registrationDisabled) {
		// currently, the code of the registry can not be updated once it is
		// deployed. if a newer version of the registry is available, account
		// registration can be disabled
		if (msg.sender == _registryAdmin) {
			_registrationDisabled = registrationDisabled;
		}
	}
	
	function adminSetAccountAdministrator(address accountAdmin) {
		if (msg.sender == _registryAdmin) {
			_accountAdmin = accountAdmin;
		}
	}
	
	function adminRetrieveDonations() {
		if (msg.sender == _registryAdmin) {
			_registryAdmin.send(this.balance);
		}
	}
			
	function adminDeleteRegistry() {
		if (msg.sender == _registryAdmin) {
			suicide(_registryAdmin); // this is a predefined function, it deletes the contract and returns all funds to the admin's address
		}
	}
}