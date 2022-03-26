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

// AxieMetaData contains all meta data concerning the Axie contract.
var AxieMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sireId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"breedAxies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AxieABI is the input ABI used to generate the binding from.
// Deprecated: Use AxieMetaData.ABI instead.
var AxieABI = AxieMetaData.ABI

// Axie is an auto generated Go binding around an Ethereum contract.
type Axie struct {
	AxieCaller     // Read-only binding to the contract
	AxieTransactor // Write-only binding to the contract
	AxieFilterer   // Log filterer for contract events
}

// AxieCaller is an auto generated read-only Go binding around an Ethereum contract.
type AxieCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxieTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AxieTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxieFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AxieFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxieSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AxieSession struct {
	Contract     *Axie             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxieCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AxieCallerSession struct {
	Contract *AxieCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AxieTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AxieTransactorSession struct {
	Contract     *AxieTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxieRaw is an auto generated low-level Go binding around an Ethereum contract.
type AxieRaw struct {
	Contract *Axie // Generic contract binding to access the raw methods on
}

// AxieCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AxieCallerRaw struct {
	Contract *AxieCaller // Generic read-only contract binding to access the raw methods on
}

// AxieTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AxieTransactorRaw struct {
	Contract *AxieTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAxie creates a new instance of Axie, bound to a specific deployed contract.
func NewAxie(address common.Address, backend bind.ContractBackend) (*Axie, error) {
	contract, err := bindAxie(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Axie{AxieCaller: AxieCaller{contract: contract}, AxieTransactor: AxieTransactor{contract: contract}, AxieFilterer: AxieFilterer{contract: contract}}, nil
}

// NewAxieCaller creates a new read-only instance of Axie, bound to a specific deployed contract.
func NewAxieCaller(address common.Address, caller bind.ContractCaller) (*AxieCaller, error) {
	contract, err := bindAxie(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AxieCaller{contract: contract}, nil
}

// NewAxieTransactor creates a new write-only instance of Axie, bound to a specific deployed contract.
func NewAxieTransactor(address common.Address, transactor bind.ContractTransactor) (*AxieTransactor, error) {
	contract, err := bindAxie(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AxieTransactor{contract: contract}, nil
}

// NewAxieFilterer creates a new log filterer instance of Axie, bound to a specific deployed contract.
func NewAxieFilterer(address common.Address, filterer bind.ContractFilterer) (*AxieFilterer, error) {
	contract, err := bindAxie(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AxieFilterer{contract: contract}, nil
}

// bindAxie binds a generic wrapper to an already deployed contract.
func bindAxie(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AxieABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Axie *AxieRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Axie.Contract.AxieCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Axie *AxieRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axie.Contract.AxieTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Axie *AxieRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Axie.Contract.AxieTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Axie *AxieCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Axie.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Axie *AxieTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axie.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Axie *AxieTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Axie.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Axie *AxieCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Axie.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Axie *AxieSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Axie.Contract.BalanceOf(&_Axie.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Axie *AxieCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Axie.Contract.BalanceOf(&_Axie.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256)
func (_Axie *AxieCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Axie.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256)
func (_Axie *AxieSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Axie.Contract.TokenOfOwnerByIndex(&_Axie.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256)
func (_Axie *AxieCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Axie.Contract.TokenOfOwnerByIndex(&_Axie.CallOpts, _owner, _index)
}

// BreedAxies is a paid mutator transaction binding the contract method 0x8264f2c2.
//
// Solidity: function breedAxies(uint256 _sireId, uint256 _matronId) returns(bool _success)
func (_Axie *AxieTransactor) BreedAxies(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _Axie.contract.Transact(opts, "breedAxies", _sireId, _matronId)
}

// BreedAxies is a paid mutator transaction binding the contract method 0x8264f2c2.
//
// Solidity: function breedAxies(uint256 _sireId, uint256 _matronId) returns(bool _success)
func (_Axie *AxieSession) BreedAxies(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _Axie.Contract.BreedAxies(&_Axie.TransactOpts, _sireId, _matronId)
}

// BreedAxies is a paid mutator transaction binding the contract method 0x8264f2c2.
//
// Solidity: function breedAxies(uint256 _sireId, uint256 _matronId) returns(bool _success)
func (_Axie *AxieTransactorSession) BreedAxies(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _Axie.Contract.BreedAxies(&_Axie.TransactOpts, _sireId, _matronId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns(bool _success)
func (_Axie *AxieTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axie.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns(bool _success)
func (_Axie *AxieSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axie.Contract.SafeTransferFrom(&_Axie.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns(bool _success)
func (_Axie *AxieTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axie.Contract.SafeTransferFrom(&_Axie.TransactOpts, _from, _to, _tokenId)
}
