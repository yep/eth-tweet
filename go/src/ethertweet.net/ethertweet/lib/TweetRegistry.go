// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package ethertweet

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TweetRegistryABI is the input ABI used to generate the binding from.
const TweetRegistryABI = `[{"constant":false,"inputs":[{"name":"name","type":"string"}],"name":"adminUnregister","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"name","type":"string"},{"name":"accountAddress","type":"address"}],"name":"register","outputs":[{"name":"result","type":"int256"}],"type":"function"},{"constant":true,"inputs":[],"name":"getNumberOfAccounts","outputs":[{"name":"numberOfAccounts","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[],"name":"adminRetrieveDonations","outputs":[],"type":"function"},{"constant":true,"inputs":[{"name":"name","type":"string"}],"name":"getAddressOfName","outputs":[{"name":"addr","type":"address"}],"type":"function"},{"constant":false,"inputs":[],"name":"adminDeleteRegistry","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"accountAdmin","type":"address"}],"name":"adminSetAccountAdministrator","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"registrationDisabled","type":"bool"}],"name":"adminSetRegistrationDisabled","outputs":[],"type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"getNameOfAddress","outputs":[{"name":"name","type":"string"}],"type":"function"},{"constant":false,"inputs":[],"name":"unregister","outputs":[{"name":"unregisteredAccountName","type":"string"}],"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getAddressOfId","outputs":[{"name":"addr","type":"address"}],"type":"function"},{"inputs":[],"type":"constructor"}]`

// TweetRegistryBin is the compiled bytecode used for deploying new contracts.
const TweetRegistryBin = `0x606060405260048054600160a060020a03199081163390811790925560058054600060035560a060020a60ff02199216909217169055610923806100436000396000f36060604052361561008d5760e060020a60003504630cc04b55811461008f5780631e59c529146101b9578063309e36ef14610268578063345e34161461027e57806338ec6ba8146102c25780633af41dc21461035957806349f0c90d146103825780639b6d86d6146103b5578063b2dad25d146103fd578063e79a198f1461047b578063ec43eeb6146104f6575b005b6040805160206004803580820135601f810184900484028501840190955284845261008d94919360249390929184019190819084018382808284375050955493955060009450505050600160a060020a039081163390911614806101025750600554600160a060020a0390811633909116145b156108ef57600260005082604051808280519060200190808383829060006004602084601f0104600302600f01f150909101848152604080519182900360209081018320548382018084526000808652600160a060020a03929092168083528284529382208054955181845292849020949a5098601f6000196001881615610100020190961604850192909204830196509450909250908390106108f357805160ff19168380011785555b5061089b92915061066a565b6040805160206004803580820135601f810184900484028501840190955284845261026c949193602493909291840191908190840183828082843750949650509335935050505060006000600160a060020a0316600260005084604051808280519060200190808383829060006004602084601f0104600302600f01f15090910193845250506040519182900360200190912054600160a060020a0316909114905061059d57506000196106fe565b6003545b60408051918252519081900360200190f35b61008d600454600160a060020a039081163390911614156102c057604051600454600160a060020a03908116916000913016319082818181858883f150505050505b565b6040805160206004803580820135601f81018490048402850184019095528484526105129491936024939092918401919081908401838280828437509496505050505050506000600260005082604051808280519060200190808383829060006004602084601f0104600302600f01f15090910193845250506040519182900360200190912054600160a060020a03169392505050565b61008d600454600160a060020a039081163390911614156102c057600454600160a060020a0316ff5b61008d600435600454600160a060020a039081163390911614156103fa5760058054600160a060020a0319168217905550565b61008d600435600454600160a060020a039081163390911614156103fa576005805474ff0000000000000000000000000000000000000000191660a060020a83021790555b50565b61052f60043560408051602081810183526000808352600160a060020a0385168152808252835190849020805460026001821615610100026000190190911604601f810184900484028301840190955284825292939092918301828280156107635780601f1061073857610100808354040283529160200191610763565b61052f6040805160208181018352600080835233600160a060020a03168152808252835190849020805460026001821615610100026000190190911604601f8101849004840283018401909552848252929390929183018282801561079a5780601f1061076f5761010080835404028352916020019161079a565b600435600090815260016020526040902054600160a060020a03165b60408051600160a060020a03929092168252519081900360200190f35b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561058f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600160a060020a03821660009081526020819052604081205460026001821615610100026000190190911604146105d757506001196106fe565b8251604090106105ea57506002196106fe565b60055460a060020a900460ff161561060557506003196106fe565b600160a060020a03821660009081526020818152604082208551815482855293839020919360026001821615610100026000190190911604601f90810184900483019391929188019083901061070457805160ff19168380011785555b5061067e9291505b80821115610734576000815560010161066a565b505081600260005084604051808280519060200190808383829060006004602084601f0104600302600f01f1509091019384525050604080519283900360209081019093208054600160a060020a031990811690951790556003805460009081526001948590529182208054909516871790945583549092019092559150505b92915050565b82800160010185558215610662579182015b82811115610662578251826000505591602001919060010190610716565b5090565b820191906000526020600020905b81548152906001019060200180831161074657829003601f168201915b50939695505050505050565b820191906000526020600020905b81548152906001019060200180831161077d57829003601f168201915b5050604080516020818101808452600080845233600160a060020a03168152808352938420805493518186529483902098995097601f6002600019600187161561010002019095169490940484019290920482019650929450925083901061081557805160ff19168380011785555b5061084592915061066a565b82800160010185558215610809579182015b82811115610809578251826000505591602001919060010190610827565b50506000600260005082604051808280519060200190808383829060006004602084601f0104600302600f01f1509091019384525050604051918290036020019091208054600160a060020a0319169055505090565b50506000600260005083604051808280519060200190808383829060006004602084601f0104600302600f01f1509091019384525050604051918290036020019091208054600160a060020a031916905550505b5050565b828001600101855582156101ad579182015b828111156101ad57825182600050559160200191906001019061090556`

// DeployTweetRegistry deploys a new Ethereum contract, binding an instance of TweetRegistry to it.
func DeployTweetRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TweetRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(TweetRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TweetRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TweetRegistry{TweetRegistryCaller: TweetRegistryCaller{contract: contract}, TweetRegistryTransactor: TweetRegistryTransactor{contract: contract}}, nil
}

// TweetRegistry is an auto generated Go binding around an Ethereum contract.
type TweetRegistry struct {
	TweetRegistryCaller     // Read-only binding to the contract
	TweetRegistryTransactor // Write-only binding to the contract
}

// TweetRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TweetRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TweetRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TweetRegistrySession struct {
	Contract     *TweetRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TweetRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TweetRegistryCallerSession struct {
	Contract *TweetRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TweetRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TweetRegistryTransactorSession struct {
	Contract     *TweetRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TweetRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TweetRegistryRaw struct {
	Contract *TweetRegistry // Generic contract binding to access the raw methods on
}

// TweetRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TweetRegistryCallerRaw struct {
	Contract *TweetRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TweetRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TweetRegistryTransactorRaw struct {
	Contract *TweetRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTweetRegistry creates a new instance of TweetRegistry, bound to a specific deployed contract.
func NewTweetRegistry(address common.Address, backend bind.ContractBackend) (*TweetRegistry, error) {
	contract, err := bindTweetRegistry(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &TweetRegistry{TweetRegistryCaller: TweetRegistryCaller{contract: contract}, TweetRegistryTransactor: TweetRegistryTransactor{contract: contract}}, nil
}

// NewTweetRegistryCaller creates a new read-only instance of TweetRegistry, bound to a specific deployed contract.
func NewTweetRegistryCaller(address common.Address, caller bind.ContractCaller) (*TweetRegistryCaller, error) {
	contract, err := bindTweetRegistry(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TweetRegistryCaller{contract: contract}, nil
}

// NewTweetRegistryTransactor creates a new write-only instance of TweetRegistry, bound to a specific deployed contract.
func NewTweetRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TweetRegistryTransactor, error) {
	contract, err := bindTweetRegistry(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TweetRegistryTransactor{contract: contract}, nil
}

// bindTweetRegistry binds a generic wrapper to an already deployed contract.
func bindTweetRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TweetRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TweetRegistry *TweetRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TweetRegistry.Contract.TweetRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TweetRegistry *TweetRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetRegistry.Contract.TweetRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TweetRegistry *TweetRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TweetRegistry.Contract.TweetRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TweetRegistry *TweetRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TweetRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TweetRegistry *TweetRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TweetRegistry *TweetRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TweetRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(id uint256) constant returns(addr address)
func (_TweetRegistry *TweetRegistryCaller) GetAddressOfId(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TweetRegistry.contract.Call(opts, out, "getAddressOfId", id)
	return *ret0, err
}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(id uint256) constant returns(addr address)
func (_TweetRegistry *TweetRegistrySession) GetAddressOfId(id *big.Int) (common.Address, error) {
	return _TweetRegistry.Contract.GetAddressOfId(&_TweetRegistry.CallOpts, id)
}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(id uint256) constant returns(addr address)
func (_TweetRegistry *TweetRegistryCallerSession) GetAddressOfId(id *big.Int) (common.Address, error) {
	return _TweetRegistry.Contract.GetAddressOfId(&_TweetRegistry.CallOpts, id)
}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(name string) constant returns(addr address)
func (_TweetRegistry *TweetRegistryCaller) GetAddressOfName(opts *bind.CallOpts, name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TweetRegistry.contract.Call(opts, out, "getAddressOfName", name)
	return *ret0, err
}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(name string) constant returns(addr address)
func (_TweetRegistry *TweetRegistrySession) GetAddressOfName(name string) (common.Address, error) {
	return _TweetRegistry.Contract.GetAddressOfName(&_TweetRegistry.CallOpts, name)
}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(name string) constant returns(addr address)
func (_TweetRegistry *TweetRegistryCallerSession) GetAddressOfName(name string) (common.Address, error) {
	return _TweetRegistry.Contract.GetAddressOfName(&_TweetRegistry.CallOpts, name)
}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(addr address) constant returns(name string)
func (_TweetRegistry *TweetRegistryCaller) GetNameOfAddress(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TweetRegistry.contract.Call(opts, out, "getNameOfAddress", addr)
	return *ret0, err
}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(addr address) constant returns(name string)
func (_TweetRegistry *TweetRegistrySession) GetNameOfAddress(addr common.Address) (string, error) {
	return _TweetRegistry.Contract.GetNameOfAddress(&_TweetRegistry.CallOpts, addr)
}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(addr address) constant returns(name string)
func (_TweetRegistry *TweetRegistryCallerSession) GetNameOfAddress(addr common.Address) (string, error) {
	return _TweetRegistry.Contract.GetNameOfAddress(&_TweetRegistry.CallOpts, addr)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(numberOfAccounts uint256)
func (_TweetRegistry *TweetRegistryCaller) GetNumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TweetRegistry.contract.Call(opts, out, "getNumberOfAccounts")
	return *ret0, err
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(numberOfAccounts uint256)
func (_TweetRegistry *TweetRegistrySession) GetNumberOfAccounts() (*big.Int, error) {
	return _TweetRegistry.Contract.GetNumberOfAccounts(&_TweetRegistry.CallOpts)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(numberOfAccounts uint256)
func (_TweetRegistry *TweetRegistryCallerSession) GetNumberOfAccounts() (*big.Int, error) {
	return _TweetRegistry.Contract.GetNumberOfAccounts(&_TweetRegistry.CallOpts)
}

// AdminDeleteRegistry is a paid mutator transaction binding the contract method 0x3af41dc2.
//
// Solidity: function adminDeleteRegistry() returns()
func (_TweetRegistry *TweetRegistryTransactor) AdminDeleteRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetRegistry.contract.Transact(opts, "adminDeleteRegistry")
}

// AdminDeleteRegistry is a paid mutator transaction binding the contract method 0x3af41dc2.
//
// Solidity: function adminDeleteRegistry() returns()
func (_TweetRegistry *TweetRegistrySession) AdminDeleteRegistry() (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminDeleteRegistry(&_TweetRegistry.TransactOpts)
}

// AdminDeleteRegistry is a paid mutator transaction binding the contract method 0x3af41dc2.
//
// Solidity: function adminDeleteRegistry() returns()
func (_TweetRegistry *TweetRegistryTransactorSession) AdminDeleteRegistry() (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminDeleteRegistry(&_TweetRegistry.TransactOpts)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_TweetRegistry *TweetRegistryTransactor) AdminRetrieveDonations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetRegistry.contract.Transact(opts, "adminRetrieveDonations")
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_TweetRegistry *TweetRegistrySession) AdminRetrieveDonations() (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminRetrieveDonations(&_TweetRegistry.TransactOpts)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_TweetRegistry *TweetRegistryTransactorSession) AdminRetrieveDonations() (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminRetrieveDonations(&_TweetRegistry.TransactOpts)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(accountAdmin address) returns()
func (_TweetRegistry *TweetRegistryTransactor) AdminSetAccountAdministrator(opts *bind.TransactOpts, accountAdmin common.Address) (*types.Transaction, error) {
	return _TweetRegistry.contract.Transact(opts, "adminSetAccountAdministrator", accountAdmin)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(accountAdmin address) returns()
func (_TweetRegistry *TweetRegistrySession) AdminSetAccountAdministrator(accountAdmin common.Address) (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminSetAccountAdministrator(&_TweetRegistry.TransactOpts, accountAdmin)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(accountAdmin address) returns()
func (_TweetRegistry *TweetRegistryTransactorSession) AdminSetAccountAdministrator(accountAdmin common.Address) (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminSetAccountAdministrator(&_TweetRegistry.TransactOpts, accountAdmin)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(registrationDisabled bool) returns()
func (_TweetRegistry *TweetRegistryTransactor) AdminSetRegistrationDisabled(opts *bind.TransactOpts, registrationDisabled bool) (*types.Transaction, error) {
	return _TweetRegistry.contract.Transact(opts, "adminSetRegistrationDisabled", registrationDisabled)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(registrationDisabled bool) returns()
func (_TweetRegistry *TweetRegistrySession) AdminSetRegistrationDisabled(registrationDisabled bool) (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminSetRegistrationDisabled(&_TweetRegistry.TransactOpts, registrationDisabled)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(registrationDisabled bool) returns()
func (_TweetRegistry *TweetRegistryTransactorSession) AdminSetRegistrationDisabled(registrationDisabled bool) (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminSetRegistrationDisabled(&_TweetRegistry.TransactOpts, registrationDisabled)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(name string) returns()
func (_TweetRegistry *TweetRegistryTransactor) AdminUnregister(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _TweetRegistry.contract.Transact(opts, "adminUnregister", name)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(name string) returns()
func (_TweetRegistry *TweetRegistrySession) AdminUnregister(name string) (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminUnregister(&_TweetRegistry.TransactOpts, name)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(name string) returns()
func (_TweetRegistry *TweetRegistryTransactorSession) AdminUnregister(name string) (*types.Transaction, error) {
	return _TweetRegistry.Contract.AdminUnregister(&_TweetRegistry.TransactOpts, name)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(name string, accountAddress address) returns(result int256)
func (_TweetRegistry *TweetRegistryTransactor) Register(opts *bind.TransactOpts, name string, accountAddress common.Address) (*types.Transaction, error) {
	return _TweetRegistry.contract.Transact(opts, "register", name, accountAddress)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(name string, accountAddress address) returns(result int256)
func (_TweetRegistry *TweetRegistrySession) Register(name string, accountAddress common.Address) (*types.Transaction, error) {
	return _TweetRegistry.Contract.Register(&_TweetRegistry.TransactOpts, name, accountAddress)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(name string, accountAddress address) returns(result int256)
func (_TweetRegistry *TweetRegistryTransactorSession) Register(name string, accountAddress common.Address) (*types.Transaction, error) {
	return _TweetRegistry.Contract.Register(&_TweetRegistry.TransactOpts, name, accountAddress)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns(unregisteredAccountName string)
func (_TweetRegistry *TweetRegistryTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetRegistry.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns(unregisteredAccountName string)
func (_TweetRegistry *TweetRegistrySession) Unregister() (*types.Transaction, error) {
	return _TweetRegistry.Contract.Unregister(&_TweetRegistry.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns(unregisteredAccountName string)
func (_TweetRegistry *TweetRegistryTransactorSession) Unregister() (*types.Transaction, error) {
	return _TweetRegistry.Contract.Unregister(&_TweetRegistry.TransactOpts)
}
