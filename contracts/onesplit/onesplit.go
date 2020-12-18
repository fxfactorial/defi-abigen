// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onesplit

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OnesplitABI is the input ABI used to generate the binding from.
const OnesplitABI = "[{\"inputs\":[{\"internalType\":\"contractIOneSplitView\",\"name\":\"_oneSplitView\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destTokenEthPriceTimesGasPrice\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturnWithGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimateGasAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oneSplitView\",\"outputs\":[{\"internalType\":\"contractIOneSplitView\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Onesplit is an auto generated Go binding around an Ethereum contract.
type Onesplit struct {
	OnesplitCaller     // Read-only binding to the contract
	OnesplitTransactor // Write-only binding to the contract
	OnesplitFilterer   // Log filterer for contract events
}

// OnesplitCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnesplitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnesplitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnesplitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnesplitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnesplitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnesplitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnesplitSession struct {
	Contract     *Onesplit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnesplitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnesplitCallerSession struct {
	Contract *OnesplitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OnesplitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnesplitTransactorSession struct {
	Contract     *OnesplitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OnesplitRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnesplitRaw struct {
	Contract *Onesplit // Generic contract binding to access the raw methods on
}

// OnesplitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnesplitCallerRaw struct {
	Contract *OnesplitCaller // Generic read-only contract binding to access the raw methods on
}

// OnesplitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnesplitTransactorRaw struct {
	Contract *OnesplitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnesplit creates a new instance of Onesplit, bound to a specific deployed contract.
func NewOnesplit(address common.Address, backend bind.ContractBackend) (*Onesplit, error) {
	contract, err := bindOnesplit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Onesplit{OnesplitCaller: OnesplitCaller{contract: contract}, OnesplitTransactor: OnesplitTransactor{contract: contract}, OnesplitFilterer: OnesplitFilterer{contract: contract}}, nil
}

// NewOnesplitCaller creates a new read-only instance of Onesplit, bound to a specific deployed contract.
func NewOnesplitCaller(address common.Address, caller bind.ContractCaller) (*OnesplitCaller, error) {
	contract, err := bindOnesplit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnesplitCaller{contract: contract}, nil
}

// NewOnesplitTransactor creates a new write-only instance of Onesplit, bound to a specific deployed contract.
func NewOnesplitTransactor(address common.Address, transactor bind.ContractTransactor) (*OnesplitTransactor, error) {
	contract, err := bindOnesplit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnesplitTransactor{contract: contract}, nil
}

// NewOnesplitFilterer creates a new log filterer instance of Onesplit, bound to a specific deployed contract.
func NewOnesplitFilterer(address common.Address, filterer bind.ContractFilterer) (*OnesplitFilterer, error) {
	contract, err := bindOnesplit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnesplitFilterer{contract: contract}, nil
}

// bindOnesplit binds a generic wrapper to an already deployed contract.
func bindOnesplit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnesplitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Onesplit *OnesplitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Onesplit.Contract.OnesplitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Onesplit *OnesplitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onesplit.Contract.OnesplitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Onesplit *OnesplitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Onesplit.Contract.OnesplitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Onesplit *OnesplitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Onesplit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Onesplit *OnesplitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onesplit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Onesplit *OnesplitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Onesplit.Contract.contract.Transact(opts, method, params...)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Onesplit *OnesplitCaller) GetExpectedReturn(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	var out []interface{}
	err := _Onesplit.contract.Call(opts, &out, "getExpectedReturn", fromToken, destToken, amount, parts, flags)

	outstruct := new(struct {
		ReturnAmount *big.Int
		Distribution []*big.Int
	})

	outstruct.ReturnAmount = out[0].(*big.Int)
	outstruct.Distribution = out[1].([]*big.Int)

	return *outstruct, err

}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Onesplit *OnesplitSession) GetExpectedReturn(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _Onesplit.Contract.GetExpectedReturn(&_Onesplit.CallOpts, fromToken, destToken, amount, parts, flags)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Onesplit *OnesplitCallerSession) GetExpectedReturn(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _Onesplit.Contract.GetExpectedReturn(&_Onesplit.CallOpts, fromToken, destToken, amount, parts, flags)
}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_Onesplit *OnesplitCaller) GetExpectedReturnWithGas(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	var out []interface{}
	err := _Onesplit.contract.Call(opts, &out, "getExpectedReturnWithGas", fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)

	outstruct := new(struct {
		ReturnAmount      *big.Int
		EstimateGasAmount *big.Int
		Distribution      []*big.Int
	})

	outstruct.ReturnAmount = out[0].(*big.Int)
	outstruct.EstimateGasAmount = out[1].(*big.Int)
	outstruct.Distribution = out[2].([]*big.Int)

	return *outstruct, err

}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_Onesplit *OnesplitSession) GetExpectedReturnWithGas(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _Onesplit.Contract.GetExpectedReturnWithGas(&_Onesplit.CallOpts, fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)
}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_Onesplit *OnesplitCallerSession) GetExpectedReturnWithGas(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _Onesplit.Contract.GetExpectedReturnWithGas(&_Onesplit.CallOpts, fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)
}

// OneSplitView is a free data retrieval call binding the contract method 0xfbe4ed95.
//
// Solidity: function oneSplitView() view returns(address)
func (_Onesplit *OnesplitCaller) OneSplitView(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Onesplit.contract.Call(opts, &out, "oneSplitView")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OneSplitView is a free data retrieval call binding the contract method 0xfbe4ed95.
//
// Solidity: function oneSplitView() view returns(address)
func (_Onesplit *OnesplitSession) OneSplitView() (common.Address, error) {
	return _Onesplit.Contract.OneSplitView(&_Onesplit.CallOpts)
}

// OneSplitView is a free data retrieval call binding the contract method 0xfbe4ed95.
//
// Solidity: function oneSplitView() view returns(address)
func (_Onesplit *OnesplitCallerSession) OneSplitView() (common.Address, error) {
	return _Onesplit.Contract.OneSplitView(&_Onesplit.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 flags) payable returns(uint256 returnAmount)
func (_Onesplit *OnesplitTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, flags *big.Int) (*types.Transaction, error) {
	return _Onesplit.contract.Transact(opts, "swap", fromToken, destToken, amount, minReturn, distribution, flags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 flags) payable returns(uint256 returnAmount)
func (_Onesplit *OnesplitSession) Swap(fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, flags *big.Int) (*types.Transaction, error) {
	return _Onesplit.Contract.Swap(&_Onesplit.TransactOpts, fromToken, destToken, amount, minReturn, distribution, flags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 flags) payable returns(uint256 returnAmount)
func (_Onesplit *OnesplitTransactorSession) Swap(fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, flags *big.Int) (*types.Transaction, error) {
	return _Onesplit.Contract.Swap(&_Onesplit.TransactOpts, fromToken, destToken, amount, minReturn, distribution, flags)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Onesplit *OnesplitTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Onesplit.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Onesplit *OnesplitSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Onesplit.Contract.Fallback(&_Onesplit.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Onesplit *OnesplitTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Onesplit.Contract.Fallback(&_Onesplit.TransactOpts, calldata)
}
