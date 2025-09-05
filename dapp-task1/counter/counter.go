// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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
	_ = abi.ConvertType
)

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"addNb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101918061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063036798f8146100435780637c28f21d14610061578063d2c00e241461007f575b5f5ffd5b61004b61009d565b60405161005891906100e2565b60405180910390f35b6100696100bd565b60405161007691906100e2565b60405180910390f35b6100876100c5565b60405161009491906100e2565b60405180910390f35b5f60015f5f8282546100af9190610128565b925050819055505f54905090565b5f5f54905090565b5f5481565b5f819050919050565b6100dc816100ca565b82525050565b5f6020820190506100f55f8301846100d3565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610132826100ca565b915061013d836100ca565b9250828201905080821115610155576101546100fb565b5b9291505056fea2646970667358221220834e5a4d5d1150a30a33dd78cc41efe0c0b924b4265a855e340c6506d66d514e64736f6c634300081e0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// GetNb is a free data retrieval call binding the contract method 0x7c28f21d.
//
// Solidity: function getNb() view returns(uint256)
func (_Counter *CounterCaller) GetNb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getNb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNb is a free data retrieval call binding the contract method 0x7c28f21d.
//
// Solidity: function getNb() view returns(uint256)
func (_Counter *CounterSession) GetNb() (*big.Int, error) {
	return _Counter.Contract.GetNb(&_Counter.CallOpts)
}

// GetNb is a free data retrieval call binding the contract method 0x7c28f21d.
//
// Solidity: function getNb() view returns(uint256)
func (_Counter *CounterCallerSession) GetNb() (*big.Int, error) {
	return _Counter.Contract.GetNb(&_Counter.CallOpts)
}

// Nb is a free data retrieval call binding the contract method 0xd2c00e24.
//
// Solidity: function nb() view returns(uint256)
func (_Counter *CounterCaller) Nb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "nb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nb is a free data retrieval call binding the contract method 0xd2c00e24.
//
// Solidity: function nb() view returns(uint256)
func (_Counter *CounterSession) Nb() (*big.Int, error) {
	return _Counter.Contract.Nb(&_Counter.CallOpts)
}

// Nb is a free data retrieval call binding the contract method 0xd2c00e24.
//
// Solidity: function nb() view returns(uint256)
func (_Counter *CounterCallerSession) Nb() (*big.Int, error) {
	return _Counter.Contract.Nb(&_Counter.CallOpts)
}

// AddNb is a paid mutator transaction binding the contract method 0x036798f8.
//
// Solidity: function addNb() returns(uint256)
func (_Counter *CounterTransactor) AddNb(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "addNb")
}

// AddNb is a paid mutator transaction binding the contract method 0x036798f8.
//
// Solidity: function addNb() returns(uint256)
func (_Counter *CounterSession) AddNb() (*types.Transaction, error) {
	return _Counter.Contract.AddNb(&_Counter.TransactOpts)
}

// AddNb is a paid mutator transaction binding the contract method 0x036798f8.
//
// Solidity: function addNb() returns(uint256)
func (_Counter *CounterTransactorSession) AddNb() (*types.Transaction, error) {
	return _Counter.Contract.AddNb(&_Counter.TransactOpts)
}
