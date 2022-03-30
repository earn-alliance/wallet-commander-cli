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

// SlpMetaData contains all meta data concerning the Slp contract.
var SlpMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"checkpoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SlpABI is the input ABI used to generate the binding from.
// Deprecated: Use SlpMetaData.ABI instead.
var SlpABI = SlpMetaData.ABI

// Slp is an auto generated Go binding around an Ethereum contract.
type Slp struct {
	SlpCaller     // Read-only binding to the contract
	SlpTransactor // Write-only binding to the contract
	SlpFilterer   // Log filterer for contract events
}

// SlpCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlpSession struct {
	Contract     *Slp              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlpCallerSession struct {
	Contract *SlpCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SlpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlpTransactorSession struct {
	Contract     *SlpTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlpRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlpRaw struct {
	Contract *Slp // Generic contract binding to access the raw methods on
}

// SlpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlpCallerRaw struct {
	Contract *SlpCaller // Generic read-only contract binding to access the raw methods on
}

// SlpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlpTransactorRaw struct {
	Contract *SlpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlp creates a new instance of Slp, bound to a specific deployed contract.
func NewSlp(address common.Address, backend bind.ContractBackend) (*Slp, error) {
	contract, err := bindSlp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Slp{SlpCaller: SlpCaller{contract: contract}, SlpTransactor: SlpTransactor{contract: contract}, SlpFilterer: SlpFilterer{contract: contract}}, nil
}

// NewSlpCaller creates a new read-only instance of Slp, bound to a specific deployed contract.
func NewSlpCaller(address common.Address, caller bind.ContractCaller) (*SlpCaller, error) {
	contract, err := bindSlp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlpCaller{contract: contract}, nil
}

// NewSlpTransactor creates a new write-only instance of Slp, bound to a specific deployed contract.
func NewSlpTransactor(address common.Address, transactor bind.ContractTransactor) (*SlpTransactor, error) {
	contract, err := bindSlp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlpTransactor{contract: contract}, nil
}

// NewSlpFilterer creates a new log filterer instance of Slp, bound to a specific deployed contract.
func NewSlpFilterer(address common.Address, filterer bind.ContractFilterer) (*SlpFilterer, error) {
	contract, err := bindSlp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlpFilterer{contract: contract}, nil
}

// bindSlp binds a generic wrapper to an already deployed contract.
func bindSlp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SlpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slp *SlpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slp.Contract.SlpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slp *SlpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slp.Contract.SlpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slp *SlpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slp.Contract.SlpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slp *SlpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slp *SlpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slp *SlpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slp.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Slp *SlpCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Slp.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Slp *SlpSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Slp.Contract.BalanceOf(&_Slp.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Slp *SlpCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Slp.Contract.BalanceOf(&_Slp.CallOpts, arg0)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xd3392ddf.
//
// Solidity: function checkpoint(address _owner, uint256 _amount, uint256 _createdAt, bytes _signature) returns(uint256 _balance)
func (_Slp *SlpTransactor) Checkpoint(opts *bind.TransactOpts, _owner common.Address, _amount *big.Int, _createdAt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Slp.contract.Transact(opts, "checkpoint", _owner, _amount, _createdAt, _signature)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xd3392ddf.
//
// Solidity: function checkpoint(address _owner, uint256 _amount, uint256 _createdAt, bytes _signature) returns(uint256 _balance)
func (_Slp *SlpSession) Checkpoint(_owner common.Address, _amount *big.Int, _createdAt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Slp.Contract.Checkpoint(&_Slp.TransactOpts, _owner, _amount, _createdAt, _signature)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xd3392ddf.
//
// Solidity: function checkpoint(address _owner, uint256 _amount, uint256 _createdAt, bytes _signature) returns(uint256 _balance)
func (_Slp *SlpTransactorSession) Checkpoint(_owner common.Address, _amount *big.Int, _createdAt *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Slp.Contract.Checkpoint(&_Slp.TransactOpts, _owner, _amount, _createdAt, _signature)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_Slp *SlpTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Slp.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_Slp *SlpSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Slp.Contract.Transfer(&_Slp.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_Slp *SlpTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Slp.Contract.Transfer(&_Slp.TransactOpts, _to, _value)
}
