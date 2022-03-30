// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"uint8\"},{\"name\":\"_tokenAddresses\",\"type\":\"address\"},{\"name\":\"_tokenNumbers\",\"type\":\"uint256\"},{\"name\":\"_startingPrices\",\"type\":\"uint256\"},{\"name\":\"_endingPrices\",\"type\":\"uint256\"},{\"name\":\"_exchangeTokens\",\"type\":\"address\"},{\"name\":\"_durations\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _listingIndex) returns()
func (_Marketplace *MarketplaceTransactor) CancelAuction(opts *bind.TransactOpts, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelAuction", _listingIndex)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _listingIndex) returns()
func (_Marketplace *MarketplaceSession) CancelAuction(_listingIndex *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelAuction(&_Marketplace.TransactOpts, _listingIndex)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _listingIndex) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelAuction(_listingIndex *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelAuction(&_Marketplace.TransactOpts, _listingIndex)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x5c1f8e16.
//
// Solidity: function createAuction(uint8 _tokenTypes, address _tokenAddresses, uint256 _tokenNumbers, uint256 _startingPrices, uint256 _endingPrices, address _exchangeTokens, uint256 _durations) returns()
func (_Marketplace *MarketplaceTransactor) CreateAuction(opts *bind.TransactOpts, _tokenTypes uint8, _tokenAddresses common.Address, _tokenNumbers *big.Int, _startingPrices *big.Int, _endingPrices *big.Int, _exchangeTokens common.Address, _durations *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createAuction", _tokenTypes, _tokenAddresses, _tokenNumbers, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x5c1f8e16.
//
// Solidity: function createAuction(uint8 _tokenTypes, address _tokenAddresses, uint256 _tokenNumbers, uint256 _startingPrices, uint256 _endingPrices, address _exchangeTokens, uint256 _durations) returns()
func (_Marketplace *MarketplaceSession) CreateAuction(_tokenTypes uint8, _tokenAddresses common.Address, _tokenNumbers *big.Int, _startingPrices *big.Int, _endingPrices *big.Int, _exchangeTokens common.Address, _durations *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateAuction(&_Marketplace.TransactOpts, _tokenTypes, _tokenAddresses, _tokenNumbers, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x5c1f8e16.
//
// Solidity: function createAuction(uint8 _tokenTypes, address _tokenAddresses, uint256 _tokenNumbers, uint256 _startingPrices, uint256 _endingPrices, address _exchangeTokens, uint256 _durations) returns()
func (_Marketplace *MarketplaceTransactorSession) CreateAuction(_tokenTypes uint8, _tokenAddresses common.Address, _tokenNumbers *big.Int, _startingPrices *big.Int, _endingPrices *big.Int, _exchangeTokens common.Address, _durations *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateAuction(&_Marketplace.TransactOpts, _tokenTypes, _tokenAddresses, _tokenNumbers, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}
