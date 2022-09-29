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

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Owner     common.Address
	Nonce     *big.Int
	ExpiredAt *big.Int
	Assets    []Struct0
	ExtraData []byte
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Erc      uint8
	Addr     common.Address
	Id       *big.Int
	Quantity *big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Kind      uint8
	Nonce     *big.Int
	Owner     common.Address
	Assets    []Struct0
	ExtraData []byte
}

// PortalMetaData contains all meta data concerning the Portal contract.
var PortalMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"owner\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"assets\",\"type\":\"tuple[]\",\"components\":[{\"internal_type\":\"\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internal_type\":\"\",\"name\":\"addr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"id\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"indexed\":false}],\"name\":\"Deposited\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"name\":\"Initialized\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"timeWindow\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MinWithdrawalPeriodUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"name\":\"Paused\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"usdc\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"withdrawalFeeInUsdc\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"paymentToken\",\"type\":\"address\",\"indexed\":false}],\"name\":\"PaymentUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true}],\"name\":\"RoleAdminChanged\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"sender\",\"type\":\"address\",\"indexed\":true}],\"name\":\"RoleGranted\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"sender\",\"type\":\"address\",\"indexed\":true}],\"name\":\"RoleRevoked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"routerContract\",\"type\":\"address\",\"indexed\":false}],\"name\":\"RouterContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"TreasuryUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"name\":\"Unpaused\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"owner\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"assets\",\"type\":\"tuple[]\",\"components\":[{\"internal_type\":\"\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internal_type\":\"\",\"name\":\"addr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"id\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"indexed\":false},{\"internal_type\":\"\",\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false}],\"name\":\"Withdrew\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"PORTAL_OPERATOR\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_assets\",\"type\":\"tuple[]\",\"components\":[{\"internal_type\":\"\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internal_type\":\"\",\"name\":\"addr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"id\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_paymentToken\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_path\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"calculateWithdrawalFee\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_paymentAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_paymentAmountInUsdc\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_assets\",\"type\":\"tuple[]\",\"components\":[{\"internal_type\":\"\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internal_type\":\"\",\"name\":\"addr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"id\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"indexed\":false}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_owner\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"getPortalInfo\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_info\",\"type\":\"tuple\",\"components\":[{\"internal_type\":\"\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internal_type\":\"\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"owner\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"assets\",\"type\":\"tuple[]\",\"components\":[{\"internal_type\":\"\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internal_type\":\"\",\"name\":\"addr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"id\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"quantity\",\"type\":\"uint256\"}]},{\"internal_type\":\"\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"getRoleMember\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"name\":\"grantRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"name\":\"hasRole\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_minWithdrawalPeriod\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_routerContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_usdc\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_withdrawalFeeInUsdc\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_paymentToken\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_treasury\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_owner\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_operators\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"lastWithdrawalTimestamp\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"minWithdrawalPeriod\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"nonce\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes\",\"indexed\":false}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes4\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes\",\"indexed\":false}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes4\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes\",\"indexed\":false}],\"name\":\"onERC721Received\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes4\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"name\":\"renounceRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"name\":\"revokeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"routerContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_timeWindow\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setMinWithdrawalPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_usdc\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_withdrawalFeeInUsdc\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_paymentToken\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_routerContract\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setRouterContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_treasury\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setTreasury\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"interfaceId\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"supportsInterface\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_withdrawal\",\"type\":\"tuple\",\"components\":[{\"internal_type\":\"\",\"name\":\"owner\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"assets\",\"type\":\"tuple[]\",\"components\":[{\"internal_type\":\"\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internal_type\":\"\",\"name\":\"addr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"id\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"quantity\",\"type\":\"uint256\"}]},{\"internal_type\":\"\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_signature\",\"type\":\"bytes\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_path\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":false}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawalFeeInUsdc\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false}]",
}

// PortalABI is the input ABI used to generate the binding from.
// Deprecated: Use PortalMetaData.ABI instead.
var PortalABI = PortalMetaData.ABI

// Portal is an auto generated Go binding around an Ethereum contract.
type Portal struct {
	PortalCaller     // Read-only binding to the contract
	PortalTransactor // Write-only binding to the contract
	PortalFilterer   // Log filterer for contract events
}

// PortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type PortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PortalSession struct {
	Contract     *Portal           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PortalCallerSession struct {
	Contract *PortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PortalTransactorSession struct {
	Contract     *PortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type PortalRaw struct {
	Contract *Portal // Generic contract binding to access the raw methods on
}

// PortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PortalCallerRaw struct {
	Contract *PortalCaller // Generic read-only contract binding to access the raw methods on
}

// PortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PortalTransactorRaw struct {
	Contract *PortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPortal creates a new instance of Portal, bound to a specific deployed contract.
func NewPortal(address common.Address, backend bind.ContractBackend) (*Portal, error) {
	contract, err := bindPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Portal{PortalCaller: PortalCaller{contract: contract}, PortalTransactor: PortalTransactor{contract: contract}, PortalFilterer: PortalFilterer{contract: contract}}, nil
}

// NewPortalCaller creates a new read-only instance of Portal, bound to a specific deployed contract.
func NewPortalCaller(address common.Address, caller bind.ContractCaller) (*PortalCaller, error) {
	contract, err := bindPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PortalCaller{contract: contract}, nil
}

// NewPortalTransactor creates a new write-only instance of Portal, bound to a specific deployed contract.
func NewPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*PortalTransactor, error) {
	contract, err := bindPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PortalTransactor{contract: contract}, nil
}

// NewPortalFilterer creates a new log filterer instance of Portal, bound to a specific deployed contract.
func NewPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*PortalFilterer, error) {
	contract, err := bindPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PortalFilterer{contract: contract}, nil
}

// bindPortal binds a generic wrapper to an already deployed contract.
func bindPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portal *PortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Portal.Contract.PortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portal *PortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.Contract.PortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portal *PortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portal.Contract.PortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portal *PortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Portal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portal *PortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portal *PortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portal.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Portal *PortalCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Portal *PortalSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Portal.Contract.DEFAULTADMINROLE(&_Portal.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Portal *PortalCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Portal.Contract.DEFAULTADMINROLE(&_Portal.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Portal *PortalCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Portal *PortalSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Portal.Contract.DOMAINSEPARATOR(&_Portal.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Portal *PortalCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Portal.Contract.DOMAINSEPARATOR(&_Portal.CallOpts)
}

// PORTALOPERATOR is a free data retrieval call binding the contract method 0xfb4b9637.
//
// Solidity: function PORTAL_OPERATOR() view returns(bytes32)
func (_Portal *PortalCaller) PORTALOPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "PORTAL_OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PORTALOPERATOR is a free data retrieval call binding the contract method 0xfb4b9637.
//
// Solidity: function PORTAL_OPERATOR() view returns(bytes32)
func (_Portal *PortalSession) PORTALOPERATOR() ([32]byte, error) {
	return _Portal.Contract.PORTALOPERATOR(&_Portal.CallOpts)
}

// PORTALOPERATOR is a free data retrieval call binding the contract method 0xfb4b9637.
//
// Solidity: function PORTAL_OPERATOR() view returns(bytes32)
func (_Portal *PortalCallerSession) PORTALOPERATOR() ([32]byte, error) {
	return _Portal.Contract.PORTALOPERATOR(&_Portal.CallOpts)
}

// CalculateWithdrawalFee is a free data retrieval call binding the contract method 0x017556f4.
//
// Solidity: function calculateWithdrawalFee((uint8,address,uint256,uint256)[] _assets, address _paymentToken, address[] _path) view returns(uint256 _paymentAmount, uint256 _paymentAmountInUsdc)
func (_Portal *PortalCaller) CalculateWithdrawalFee(opts *bind.CallOpts, _assets []Struct0, _paymentToken common.Address, _path []common.Address) (struct {
	PaymentAmount       *big.Int
	PaymentAmountInUsdc *big.Int
}, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "calculateWithdrawalFee", _assets, _paymentToken, _path)

	outstruct := new(struct {
		PaymentAmount       *big.Int
		PaymentAmountInUsdc *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PaymentAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PaymentAmountInUsdc = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateWithdrawalFee is a free data retrieval call binding the contract method 0x017556f4.
//
// Solidity: function calculateWithdrawalFee((uint8,address,uint256,uint256)[] _assets, address _paymentToken, address[] _path) view returns(uint256 _paymentAmount, uint256 _paymentAmountInUsdc)
func (_Portal *PortalSession) CalculateWithdrawalFee(_assets []Struct0, _paymentToken common.Address, _path []common.Address) (struct {
	PaymentAmount       *big.Int
	PaymentAmountInUsdc *big.Int
}, error) {
	return _Portal.Contract.CalculateWithdrawalFee(&_Portal.CallOpts, _assets, _paymentToken, _path)
}

// CalculateWithdrawalFee is a free data retrieval call binding the contract method 0x017556f4.
//
// Solidity: function calculateWithdrawalFee((uint8,address,uint256,uint256)[] _assets, address _paymentToken, address[] _path) view returns(uint256 _paymentAmount, uint256 _paymentAmountInUsdc)
func (_Portal *PortalCallerSession) CalculateWithdrawalFee(_assets []Struct0, _paymentToken common.Address, _path []common.Address) (struct {
	PaymentAmount       *big.Int
	PaymentAmountInUsdc *big.Int
}, error) {
	return _Portal.Contract.CalculateWithdrawalFee(&_Portal.CallOpts, _assets, _paymentToken, _path)
}

// GetPortalInfo is a free data retrieval call binding the contract method 0x09960a55.
//
// Solidity: function getPortalInfo(address _owner, uint256 _nonce) view returns((uint8,uint256,address,(uint8,address,uint256,uint256)[],bytes) _info)
func (_Portal *PortalCaller) GetPortalInfo(opts *bind.CallOpts, _owner common.Address, _nonce *big.Int) (Struct1, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "getPortalInfo", _owner, _nonce)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetPortalInfo is a free data retrieval call binding the contract method 0x09960a55.
//
// Solidity: function getPortalInfo(address _owner, uint256 _nonce) view returns((uint8,uint256,address,(uint8,address,uint256,uint256)[],bytes) _info)
func (_Portal *PortalSession) GetPortalInfo(_owner common.Address, _nonce *big.Int) (Struct1, error) {
	return _Portal.Contract.GetPortalInfo(&_Portal.CallOpts, _owner, _nonce)
}

// GetPortalInfo is a free data retrieval call binding the contract method 0x09960a55.
//
// Solidity: function getPortalInfo(address _owner, uint256 _nonce) view returns((uint8,uint256,address,(uint8,address,uint256,uint256)[],bytes) _info)
func (_Portal *PortalCallerSession) GetPortalInfo(_owner common.Address, _nonce *big.Int) (Struct1, error) {
	return _Portal.Contract.GetPortalInfo(&_Portal.CallOpts, _owner, _nonce)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Portal *PortalCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Portal *PortalSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Portal.Contract.GetRoleAdmin(&_Portal.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Portal *PortalCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Portal.Contract.GetRoleAdmin(&_Portal.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Portal *PortalCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Portal *PortalSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Portal.Contract.GetRoleMember(&_Portal.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Portal *PortalCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Portal.Contract.GetRoleMember(&_Portal.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Portal *PortalCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Portal *PortalSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Portal.Contract.GetRoleMemberCount(&_Portal.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Portal *PortalCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Portal.Contract.GetRoleMemberCount(&_Portal.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Portal *PortalCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Portal *PortalSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Portal.Contract.HasRole(&_Portal.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Portal *PortalCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Portal.Contract.HasRole(&_Portal.CallOpts, role, account)
}

// LastWithdrawalTimestamp is a free data retrieval call binding the contract method 0x2be50b89.
//
// Solidity: function lastWithdrawalTimestamp(address ) view returns(uint256)
func (_Portal *PortalCaller) LastWithdrawalTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "lastWithdrawalTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastWithdrawalTimestamp is a free data retrieval call binding the contract method 0x2be50b89.
//
// Solidity: function lastWithdrawalTimestamp(address ) view returns(uint256)
func (_Portal *PortalSession) LastWithdrawalTimestamp(arg0 common.Address) (*big.Int, error) {
	return _Portal.Contract.LastWithdrawalTimestamp(&_Portal.CallOpts, arg0)
}

// LastWithdrawalTimestamp is a free data retrieval call binding the contract method 0x2be50b89.
//
// Solidity: function lastWithdrawalTimestamp(address ) view returns(uint256)
func (_Portal *PortalCallerSession) LastWithdrawalTimestamp(arg0 common.Address) (*big.Int, error) {
	return _Portal.Contract.LastWithdrawalTimestamp(&_Portal.CallOpts, arg0)
}

// MinWithdrawalPeriod is a free data retrieval call binding the contract method 0x12b4af3e.
//
// Solidity: function minWithdrawalPeriod() view returns(uint256)
func (_Portal *PortalCaller) MinWithdrawalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "minWithdrawalPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawalPeriod is a free data retrieval call binding the contract method 0x12b4af3e.
//
// Solidity: function minWithdrawalPeriod() view returns(uint256)
func (_Portal *PortalSession) MinWithdrawalPeriod() (*big.Int, error) {
	return _Portal.Contract.MinWithdrawalPeriod(&_Portal.CallOpts)
}

// MinWithdrawalPeriod is a free data retrieval call binding the contract method 0x12b4af3e.
//
// Solidity: function minWithdrawalPeriod() view returns(uint256)
func (_Portal *PortalCallerSession) MinWithdrawalPeriod() (*big.Int, error) {
	return _Portal.Contract.MinWithdrawalPeriod(&_Portal.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Portal *PortalCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Portal *PortalSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Portal.Contract.Nonce(&_Portal.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Portal *PortalCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Portal.Contract.Nonce(&_Portal.CallOpts, arg0)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_Portal *PortalCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_Portal *PortalSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Portal.Contract.OnERC1155BatchReceived(&_Portal.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_Portal *PortalCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Portal.Contract.OnERC1155BatchReceived(&_Portal.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_Portal *PortalCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_Portal *PortalSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Portal.Contract.OnERC1155Received(&_Portal.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_Portal *PortalCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Portal.Contract.OnERC1155Received(&_Portal.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_Portal *PortalCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_Portal *PortalSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Portal.Contract.OnERC721Received(&_Portal.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_Portal *PortalCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Portal.Contract.OnERC721Received(&_Portal.CallOpts, arg0, arg1, arg2, arg3)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Portal *PortalCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Portal *PortalSession) Paused() (bool, error) {
	return _Portal.Contract.Paused(&_Portal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Portal *PortalCallerSession) Paused() (bool, error) {
	return _Portal.Contract.Paused(&_Portal.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Portal *PortalCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Portal *PortalSession) PaymentToken() (common.Address, error) {
	return _Portal.Contract.PaymentToken(&_Portal.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Portal *PortalCallerSession) PaymentToken() (common.Address, error) {
	return _Portal.Contract.PaymentToken(&_Portal.CallOpts)
}

// RouterContract is a free data retrieval call binding the contract method 0x4f72d846.
//
// Solidity: function routerContract() view returns(address)
func (_Portal *PortalCaller) RouterContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "routerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterContract is a free data retrieval call binding the contract method 0x4f72d846.
//
// Solidity: function routerContract() view returns(address)
func (_Portal *PortalSession) RouterContract() (common.Address, error) {
	return _Portal.Contract.RouterContract(&_Portal.CallOpts)
}

// RouterContract is a free data retrieval call binding the contract method 0x4f72d846.
//
// Solidity: function routerContract() view returns(address)
func (_Portal *PortalCallerSession) RouterContract() (common.Address, error) {
	return _Portal.Contract.RouterContract(&_Portal.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Portal *PortalCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Portal *PortalSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Portal.Contract.SupportsInterface(&_Portal.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Portal *PortalCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Portal.Contract.SupportsInterface(&_Portal.CallOpts, interfaceId)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Portal *PortalCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Portal *PortalSession) Treasury() (common.Address, error) {
	return _Portal.Contract.Treasury(&_Portal.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Portal *PortalCallerSession) Treasury() (common.Address, error) {
	return _Portal.Contract.Treasury(&_Portal.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Portal *PortalCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Portal *PortalSession) Usdc() (common.Address, error) {
	return _Portal.Contract.Usdc(&_Portal.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Portal *PortalCallerSession) Usdc() (common.Address, error) {
	return _Portal.Contract.Usdc(&_Portal.CallOpts)
}

// WithdrawalFeeInUsdc is a free data retrieval call binding the contract method 0x70306e37.
//
// Solidity: function withdrawalFeeInUsdc() view returns(uint256)
func (_Portal *PortalCaller) WithdrawalFeeInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "withdrawalFeeInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalFeeInUsdc is a free data retrieval call binding the contract method 0x70306e37.
//
// Solidity: function withdrawalFeeInUsdc() view returns(uint256)
func (_Portal *PortalSession) WithdrawalFeeInUsdc() (*big.Int, error) {
	return _Portal.Contract.WithdrawalFeeInUsdc(&_Portal.CallOpts)
}

// WithdrawalFeeInUsdc is a free data retrieval call binding the contract method 0x70306e37.
//
// Solidity: function withdrawalFeeInUsdc() view returns(uint256)
func (_Portal *PortalCallerSession) WithdrawalFeeInUsdc() (*big.Int, error) {
	return _Portal.Contract.WithdrawalFeeInUsdc(&_Portal.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x1253c921.
//
// Solidity: function deposit((uint8,address,uint256,uint256)[] _assets) returns()
func (_Portal *PortalTransactor) Deposit(opts *bind.TransactOpts, _assets []Struct0) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "deposit", _assets)
}

// Deposit is a paid mutator transaction binding the contract method 0x1253c921.
//
// Solidity: function deposit((uint8,address,uint256,uint256)[] _assets) returns()
func (_Portal *PortalSession) Deposit(_assets []Struct0) (*types.Transaction, error) {
	return _Portal.Contract.Deposit(&_Portal.TransactOpts, _assets)
}

// Deposit is a paid mutator transaction binding the contract method 0x1253c921.
//
// Solidity: function deposit((uint8,address,uint256,uint256)[] _assets) returns()
func (_Portal *PortalTransactorSession) Deposit(_assets []Struct0) (*types.Transaction, error) {
	return _Portal.Contract.Deposit(&_Portal.TransactOpts, _assets)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Portal *PortalTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Portal *PortalSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.Contract.GrantRole(&_Portal.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Portal *PortalTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.Contract.GrantRole(&_Portal.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xd8360cd4.
//
// Solidity: function initialize(uint256 _minWithdrawalPeriod, address _routerContract, address _usdc, uint256 _withdrawalFeeInUsdc, address _paymentToken, address _treasury, address _owner, address[] _operators) returns()
func (_Portal *PortalTransactor) Initialize(opts *bind.TransactOpts, _minWithdrawalPeriod *big.Int, _routerContract common.Address, _usdc common.Address, _withdrawalFeeInUsdc *big.Int, _paymentToken common.Address, _treasury common.Address, _owner common.Address, _operators []common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "initialize", _minWithdrawalPeriod, _routerContract, _usdc, _withdrawalFeeInUsdc, _paymentToken, _treasury, _owner, _operators)
}

// Initialize is a paid mutator transaction binding the contract method 0xd8360cd4.
//
// Solidity: function initialize(uint256 _minWithdrawalPeriod, address _routerContract, address _usdc, uint256 _withdrawalFeeInUsdc, address _paymentToken, address _treasury, address _owner, address[] _operators) returns()
func (_Portal *PortalSession) Initialize(_minWithdrawalPeriod *big.Int, _routerContract common.Address, _usdc common.Address, _withdrawalFeeInUsdc *big.Int, _paymentToken common.Address, _treasury common.Address, _owner common.Address, _operators []common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Initialize(&_Portal.TransactOpts, _minWithdrawalPeriod, _routerContract, _usdc, _withdrawalFeeInUsdc, _paymentToken, _treasury, _owner, _operators)
}

// Initialize is a paid mutator transaction binding the contract method 0xd8360cd4.
//
// Solidity: function initialize(uint256 _minWithdrawalPeriod, address _routerContract, address _usdc, uint256 _withdrawalFeeInUsdc, address _paymentToken, address _treasury, address _owner, address[] _operators) returns()
func (_Portal *PortalTransactorSession) Initialize(_minWithdrawalPeriod *big.Int, _routerContract common.Address, _usdc common.Address, _withdrawalFeeInUsdc *big.Int, _paymentToken common.Address, _treasury common.Address, _owner common.Address, _operators []common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Initialize(&_Portal.TransactOpts, _minWithdrawalPeriod, _routerContract, _usdc, _withdrawalFeeInUsdc, _paymentToken, _treasury, _owner, _operators)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Portal *PortalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Portal *PortalSession) Pause() (*types.Transaction, error) {
	return _Portal.Contract.Pause(&_Portal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Portal *PortalTransactorSession) Pause() (*types.Transaction, error) {
	return _Portal.Contract.Pause(&_Portal.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Portal *PortalTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Portal *PortalSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.Contract.RenounceRole(&_Portal.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Portal *PortalTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.Contract.RenounceRole(&_Portal.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Portal *PortalTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Portal *PortalSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.Contract.RevokeRole(&_Portal.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Portal *PortalTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Portal.Contract.RevokeRole(&_Portal.TransactOpts, role, account)
}

// SetMinWithdrawalPeriod is a paid mutator transaction binding the contract method 0x30608f6f.
//
// Solidity: function setMinWithdrawalPeriod(uint256 _timeWindow) returns()
func (_Portal *PortalTransactor) SetMinWithdrawalPeriod(opts *bind.TransactOpts, _timeWindow *big.Int) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "setMinWithdrawalPeriod", _timeWindow)
}

// SetMinWithdrawalPeriod is a paid mutator transaction binding the contract method 0x30608f6f.
//
// Solidity: function setMinWithdrawalPeriod(uint256 _timeWindow) returns()
func (_Portal *PortalSession) SetMinWithdrawalPeriod(_timeWindow *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.SetMinWithdrawalPeriod(&_Portal.TransactOpts, _timeWindow)
}

// SetMinWithdrawalPeriod is a paid mutator transaction binding the contract method 0x30608f6f.
//
// Solidity: function setMinWithdrawalPeriod(uint256 _timeWindow) returns()
func (_Portal *PortalTransactorSession) SetMinWithdrawalPeriod(_timeWindow *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.SetMinWithdrawalPeriod(&_Portal.TransactOpts, _timeWindow)
}

// SetPayment is a paid mutator transaction binding the contract method 0x8de63a56.
//
// Solidity: function setPayment(address _usdc, uint256 _withdrawalFeeInUsdc, address _paymentToken) returns()
func (_Portal *PortalTransactor) SetPayment(opts *bind.TransactOpts, _usdc common.Address, _withdrawalFeeInUsdc *big.Int, _paymentToken common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "setPayment", _usdc, _withdrawalFeeInUsdc, _paymentToken)
}

// SetPayment is a paid mutator transaction binding the contract method 0x8de63a56.
//
// Solidity: function setPayment(address _usdc, uint256 _withdrawalFeeInUsdc, address _paymentToken) returns()
func (_Portal *PortalSession) SetPayment(_usdc common.Address, _withdrawalFeeInUsdc *big.Int, _paymentToken common.Address) (*types.Transaction, error) {
	return _Portal.Contract.SetPayment(&_Portal.TransactOpts, _usdc, _withdrawalFeeInUsdc, _paymentToken)
}

// SetPayment is a paid mutator transaction binding the contract method 0x8de63a56.
//
// Solidity: function setPayment(address _usdc, uint256 _withdrawalFeeInUsdc, address _paymentToken) returns()
func (_Portal *PortalTransactorSession) SetPayment(_usdc common.Address, _withdrawalFeeInUsdc *big.Int, _paymentToken common.Address) (*types.Transaction, error) {
	return _Portal.Contract.SetPayment(&_Portal.TransactOpts, _usdc, _withdrawalFeeInUsdc, _paymentToken)
}

// SetRouterContract is a paid mutator transaction binding the contract method 0x6cc91bca.
//
// Solidity: function setRouterContract(address _routerContract) returns()
func (_Portal *PortalTransactor) SetRouterContract(opts *bind.TransactOpts, _routerContract common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "setRouterContract", _routerContract)
}

// SetRouterContract is a paid mutator transaction binding the contract method 0x6cc91bca.
//
// Solidity: function setRouterContract(address _routerContract) returns()
func (_Portal *PortalSession) SetRouterContract(_routerContract common.Address) (*types.Transaction, error) {
	return _Portal.Contract.SetRouterContract(&_Portal.TransactOpts, _routerContract)
}

// SetRouterContract is a paid mutator transaction binding the contract method 0x6cc91bca.
//
// Solidity: function setRouterContract(address _routerContract) returns()
func (_Portal *PortalTransactorSession) SetRouterContract(_routerContract common.Address) (*types.Transaction, error) {
	return _Portal.Contract.SetRouterContract(&_Portal.TransactOpts, _routerContract)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Portal *PortalTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Portal *PortalSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Portal.Contract.SetTreasury(&_Portal.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Portal *PortalTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Portal.Contract.SetTreasury(&_Portal.TransactOpts, _treasury)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Portal *PortalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Portal *PortalSession) Unpause() (*types.Transaction, error) {
	return _Portal.Contract.Unpause(&_Portal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Portal *PortalTransactorSession) Unpause() (*types.Transaction, error) {
	return _Portal.Contract.Unpause(&_Portal.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95fb7103.
//
// Solidity: function withdraw((address,uint256,uint256,(uint8,address,uint256,uint256)[],bytes) _withdrawal, bytes _signature, address[] _path) returns()
func (_Portal *PortalTransactor) Withdraw(opts *bind.TransactOpts, _withdrawal Struct2, _signature []byte, _path []common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "withdraw", _withdrawal, _signature, _path)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95fb7103.
//
// Solidity: function withdraw((address,uint256,uint256,(uint8,address,uint256,uint256)[],bytes) _withdrawal, bytes _signature, address[] _path) returns()
func (_Portal *PortalSession) Withdraw(_withdrawal Struct2, _signature []byte, _path []common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Withdraw(&_Portal.TransactOpts, _withdrawal, _signature, _path)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95fb7103.
//
// Solidity: function withdraw((address,uint256,uint256,(uint8,address,uint256,uint256)[],bytes) _withdrawal, bytes _signature, address[] _path) returns()
func (_Portal *PortalTransactorSession) Withdraw(_withdrawal Struct2, _signature []byte, _path []common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Withdraw(&_Portal.TransactOpts, _withdrawal, _signature, _path)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x5a18664c.
//
// Solidity: function withdrawNativeToken() returns()
func (_Portal *PortalTransactor) WithdrawNativeToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "withdrawNativeToken")
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x5a18664c.
//
// Solidity: function withdrawNativeToken() returns()
func (_Portal *PortalSession) WithdrawNativeToken() (*types.Transaction, error) {
	return _Portal.Contract.WithdrawNativeToken(&_Portal.TransactOpts)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x5a18664c.
//
// Solidity: function withdrawNativeToken() returns()
func (_Portal *PortalTransactorSession) WithdrawNativeToken() (*types.Transaction, error) {
	return _Portal.Contract.WithdrawNativeToken(&_Portal.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Portal *PortalTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "withdrawToken", _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Portal *PortalSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _Portal.Contract.WithdrawToken(&_Portal.TransactOpts, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Portal *PortalTransactorSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _Portal.Contract.WithdrawToken(&_Portal.TransactOpts, _token)
}

// PortalDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Portal contract.
type PortalDepositedIterator struct {
	Event *PortalDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalDeposited represents a Deposited event raised by the Portal contract.
type PortalDeposited struct {
	Nonce  *big.Int
	Owner  common.Address
	Assets []Struct0
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x08cb48d6c543fb34918502f44a07a8e07d7c4d097161fb4fad82e7ee57a9a915.
//
// Solidity: event Deposited(uint256 nonce, address owner, (uint8,address,uint256,uint256)[] assets)
func (_Portal *PortalFilterer) FilterDeposited(opts *bind.FilterOpts) (*PortalDepositedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &PortalDepositedIterator{contract: _Portal.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x08cb48d6c543fb34918502f44a07a8e07d7c4d097161fb4fad82e7ee57a9a915.
//
// Solidity: event Deposited(uint256 nonce, address owner, (uint8,address,uint256,uint256)[] assets)
func (_Portal *PortalFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *PortalDeposited) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalDeposited)
				if err := _Portal.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x08cb48d6c543fb34918502f44a07a8e07d7c4d097161fb4fad82e7ee57a9a915.
//
// Solidity: event Deposited(uint256 nonce, address owner, (uint8,address,uint256,uint256)[] assets)
func (_Portal *PortalFilterer) ParseDeposited(log types.Log) (*PortalDeposited, error) {
	event := new(PortalDeposited)
	if err := _Portal.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Portal contract.
type PortalInitializedIterator struct {
	Event *PortalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalInitialized represents a Initialized event raised by the Portal contract.
type PortalInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Portal *PortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*PortalInitializedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PortalInitializedIterator{contract: _Portal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Portal *PortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PortalInitialized) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalInitialized)
				if err := _Portal.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Portal *PortalFilterer) ParseInitialized(log types.Log) (*PortalInitialized, error) {
	event := new(PortalInitialized)
	if err := _Portal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalMinWithdrawalPeriodUpdatedIterator is returned from FilterMinWithdrawalPeriodUpdated and is used to iterate over the raw logs and unpacked data for MinWithdrawalPeriodUpdated events raised by the Portal contract.
type PortalMinWithdrawalPeriodUpdatedIterator struct {
	Event *PortalMinWithdrawalPeriodUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalMinWithdrawalPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalMinWithdrawalPeriodUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalMinWithdrawalPeriodUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalMinWithdrawalPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalMinWithdrawalPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalMinWithdrawalPeriodUpdated represents a MinWithdrawalPeriodUpdated event raised by the Portal contract.
type PortalMinWithdrawalPeriodUpdated struct {
	TimeWindow *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinWithdrawalPeriodUpdated is a free log retrieval operation binding the contract event 0x47ad7ac5d981acb4912a06e9574c034bb39f552a40d0bcb8a661a0d9cf75b0b9.
//
// Solidity: event MinWithdrawalPeriodUpdated(uint256 timeWindow)
func (_Portal *PortalFilterer) FilterMinWithdrawalPeriodUpdated(opts *bind.FilterOpts) (*PortalMinWithdrawalPeriodUpdatedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "MinWithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &PortalMinWithdrawalPeriodUpdatedIterator{contract: _Portal.contract, event: "MinWithdrawalPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchMinWithdrawalPeriodUpdated is a free log subscription operation binding the contract event 0x47ad7ac5d981acb4912a06e9574c034bb39f552a40d0bcb8a661a0d9cf75b0b9.
//
// Solidity: event MinWithdrawalPeriodUpdated(uint256 timeWindow)
func (_Portal *PortalFilterer) WatchMinWithdrawalPeriodUpdated(opts *bind.WatchOpts, sink chan<- *PortalMinWithdrawalPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "MinWithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalMinWithdrawalPeriodUpdated)
				if err := _Portal.contract.UnpackLog(event, "MinWithdrawalPeriodUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinWithdrawalPeriodUpdated is a log parse operation binding the contract event 0x47ad7ac5d981acb4912a06e9574c034bb39f552a40d0bcb8a661a0d9cf75b0b9.
//
// Solidity: event MinWithdrawalPeriodUpdated(uint256 timeWindow)
func (_Portal *PortalFilterer) ParseMinWithdrawalPeriodUpdated(log types.Log) (*PortalMinWithdrawalPeriodUpdated, error) {
	event := new(PortalMinWithdrawalPeriodUpdated)
	if err := _Portal.contract.UnpackLog(event, "MinWithdrawalPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Portal contract.
type PortalPausedIterator struct {
	Event *PortalPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalPaused represents a Paused event raised by the Portal contract.
type PortalPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Portal *PortalFilterer) FilterPaused(opts *bind.FilterOpts) (*PortalPausedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PortalPausedIterator{contract: _Portal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Portal *PortalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PortalPaused) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalPaused)
				if err := _Portal.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Portal *PortalFilterer) ParsePaused(log types.Log) (*PortalPaused, error) {
	event := new(PortalPaused)
	if err := _Portal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalPaymentUpdatedIterator is returned from FilterPaymentUpdated and is used to iterate over the raw logs and unpacked data for PaymentUpdated events raised by the Portal contract.
type PortalPaymentUpdatedIterator struct {
	Event *PortalPaymentUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalPaymentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalPaymentUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalPaymentUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalPaymentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalPaymentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalPaymentUpdated represents a PaymentUpdated event raised by the Portal contract.
type PortalPaymentUpdated struct {
	Usdc                common.Address
	WithdrawalFeeInUsdc *big.Int
	PaymentToken        common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPaymentUpdated is a free log retrieval operation binding the contract event 0x001aacbf4a93dd712e4629bf70833600ba3c672613e69604b95b4cc97fed4d78.
//
// Solidity: event PaymentUpdated(address usdc, uint256 withdrawalFeeInUsdc, address paymentToken)
func (_Portal *PortalFilterer) FilterPaymentUpdated(opts *bind.FilterOpts) (*PortalPaymentUpdatedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "PaymentUpdated")
	if err != nil {
		return nil, err
	}
	return &PortalPaymentUpdatedIterator{contract: _Portal.contract, event: "PaymentUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentUpdated is a free log subscription operation binding the contract event 0x001aacbf4a93dd712e4629bf70833600ba3c672613e69604b95b4cc97fed4d78.
//
// Solidity: event PaymentUpdated(address usdc, uint256 withdrawalFeeInUsdc, address paymentToken)
func (_Portal *PortalFilterer) WatchPaymentUpdated(opts *bind.WatchOpts, sink chan<- *PortalPaymentUpdated) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "PaymentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalPaymentUpdated)
				if err := _Portal.contract.UnpackLog(event, "PaymentUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentUpdated is a log parse operation binding the contract event 0x001aacbf4a93dd712e4629bf70833600ba3c672613e69604b95b4cc97fed4d78.
//
// Solidity: event PaymentUpdated(address usdc, uint256 withdrawalFeeInUsdc, address paymentToken)
func (_Portal *PortalFilterer) ParsePaymentUpdated(log types.Log) (*PortalPaymentUpdated, error) {
	event := new(PortalPaymentUpdated)
	if err := _Portal.contract.UnpackLog(event, "PaymentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Portal contract.
type PortalRoleAdminChangedIterator struct {
	Event *PortalRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRoleAdminChanged represents a RoleAdminChanged event raised by the Portal contract.
type PortalRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Portal *PortalFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PortalRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Portal.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PortalRoleAdminChangedIterator{contract: _Portal.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Portal *PortalFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PortalRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Portal.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRoleAdminChanged)
				if err := _Portal.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Portal *PortalFilterer) ParseRoleAdminChanged(log types.Log) (*PortalRoleAdminChanged, error) {
	event := new(PortalRoleAdminChanged)
	if err := _Portal.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Portal contract.
type PortalRoleGrantedIterator struct {
	Event *PortalRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRoleGranted represents a RoleGranted event raised by the Portal contract.
type PortalRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Portal *PortalFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PortalRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Portal.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PortalRoleGrantedIterator{contract: _Portal.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Portal *PortalFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PortalRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Portal.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRoleGranted)
				if err := _Portal.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Portal *PortalFilterer) ParseRoleGranted(log types.Log) (*PortalRoleGranted, error) {
	event := new(PortalRoleGranted)
	if err := _Portal.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Portal contract.
type PortalRoleRevokedIterator struct {
	Event *PortalRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRoleRevoked represents a RoleRevoked event raised by the Portal contract.
type PortalRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Portal *PortalFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PortalRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Portal.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PortalRoleRevokedIterator{contract: _Portal.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Portal *PortalFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PortalRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Portal.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRoleRevoked)
				if err := _Portal.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Portal *PortalFilterer) ParseRoleRevoked(log types.Log) (*PortalRoleRevoked, error) {
	event := new(PortalRoleRevoked)
	if err := _Portal.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalRouterContractUpdatedIterator is returned from FilterRouterContractUpdated and is used to iterate over the raw logs and unpacked data for RouterContractUpdated events raised by the Portal contract.
type PortalRouterContractUpdatedIterator struct {
	Event *PortalRouterContractUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalRouterContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRouterContractUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalRouterContractUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalRouterContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRouterContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRouterContractUpdated represents a RouterContractUpdated event raised by the Portal contract.
type PortalRouterContractUpdated struct {
	RouterContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRouterContractUpdated is a free log retrieval operation binding the contract event 0xe101485f5fbb64a12336091ab3a4684b5409d8006eb3b943154baa120d0c920e.
//
// Solidity: event RouterContractUpdated(address routerContract)
func (_Portal *PortalFilterer) FilterRouterContractUpdated(opts *bind.FilterOpts) (*PortalRouterContractUpdatedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "RouterContractUpdated")
	if err != nil {
		return nil, err
	}
	return &PortalRouterContractUpdatedIterator{contract: _Portal.contract, event: "RouterContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRouterContractUpdated is a free log subscription operation binding the contract event 0xe101485f5fbb64a12336091ab3a4684b5409d8006eb3b943154baa120d0c920e.
//
// Solidity: event RouterContractUpdated(address routerContract)
func (_Portal *PortalFilterer) WatchRouterContractUpdated(opts *bind.WatchOpts, sink chan<- *PortalRouterContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "RouterContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRouterContractUpdated)
				if err := _Portal.contract.UnpackLog(event, "RouterContractUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRouterContractUpdated is a log parse operation binding the contract event 0xe101485f5fbb64a12336091ab3a4684b5409d8006eb3b943154baa120d0c920e.
//
// Solidity: event RouterContractUpdated(address routerContract)
func (_Portal *PortalFilterer) ParseRouterContractUpdated(log types.Log) (*PortalRouterContractUpdated, error) {
	event := new(PortalRouterContractUpdated)
	if err := _Portal.contract.UnpackLog(event, "RouterContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the Portal contract.
type PortalTreasuryUpdatedIterator struct {
	Event *PortalTreasuryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalTreasuryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalTreasuryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalTreasuryUpdated represents a TreasuryUpdated event raised by the Portal contract.
type PortalTreasuryUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address arg0)
func (_Portal *PortalFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*PortalTreasuryUpdatedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &PortalTreasuryUpdatedIterator{contract: _Portal.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address arg0)
func (_Portal *PortalFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *PortalTreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalTreasuryUpdated)
				if err := _Portal.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address arg0)
func (_Portal *PortalFilterer) ParseTreasuryUpdated(log types.Log) (*PortalTreasuryUpdated, error) {
	event := new(PortalTreasuryUpdated)
	if err := _Portal.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Portal contract.
type PortalUnpausedIterator struct {
	Event *PortalUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalUnpaused represents a Unpaused event raised by the Portal contract.
type PortalUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Portal *PortalFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PortalUnpausedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PortalUnpausedIterator{contract: _Portal.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Portal *PortalFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PortalUnpaused) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalUnpaused)
				if err := _Portal.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Portal *PortalFilterer) ParseUnpaused(log types.Log) (*PortalUnpaused, error) {
	event := new(PortalUnpaused)
	if err := _Portal.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the Portal contract.
type PortalWithdrewIterator struct {
	Event *PortalWithdrew // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortalWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalWithdrew)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortalWithdrew)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortalWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalWithdrew represents a Withdrew event raised by the Portal contract.
type PortalWithdrew struct {
	Nonce     *big.Int
	Owner     common.Address
	Assets    []Struct0
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0xb7818c8c091e9a80958582d256533a39c5025517f6269b5d6dd4e12163a2ac90.
//
// Solidity: event Withdrew(uint256 nonce, address owner, (uint8,address,uint256,uint256)[] assets, bytes extraData)
func (_Portal *PortalFilterer) FilterWithdrew(opts *bind.FilterOpts) (*PortalWithdrewIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return &PortalWithdrewIterator{contract: _Portal.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0xb7818c8c091e9a80958582d256533a39c5025517f6269b5d6dd4e12163a2ac90.
//
// Solidity: event Withdrew(uint256 nonce, address owner, (uint8,address,uint256,uint256)[] assets, bytes extraData)
func (_Portal *PortalFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *PortalWithdrew) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalWithdrew)
				if err := _Portal.contract.UnpackLog(event, "Withdrew", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrew is a log parse operation binding the contract event 0xb7818c8c091e9a80958582d256533a39c5025517f6269b5d6dd4e12163a2ac90.
//
// Solidity: event Withdrew(uint256 nonce, address owner, (uint8,address,uint256,uint256)[] assets, bytes extraData)
func (_Portal *PortalFilterer) ParseWithdrew(log types.Log) (*PortalWithdrew, error) {
	event := new(PortalWithdrew)
	if err := _Portal.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
