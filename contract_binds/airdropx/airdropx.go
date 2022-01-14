// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_airdropx

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

// AirDropXABI is the input ABI used to generate the binding from.
const AirDropXABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"toList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountList\",\"type\":\"uint256[]\"}],\"name\":\"HelpTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AirDropXBin is the compiled bytecode used for deploying new contracts.
var AirDropXBin = "0x608060405234801561001057600080fd5b5060006100216100c460201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100cc565b600033905090565b6108f3806100db6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063534881c314610051578063715018a6146101dd5780638da5cb5b146101e7578063f2fde38b1461021b575b600080fd5b6101db6004803603608081101561006757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156100c457600080fd5b8201836020820111156100d657600080fd5b803590602001918460208302840111640100000000831117156100f857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561015857600080fd5b82018360208201111561016a57600080fd5b8035906020019184602083028401116401000000008311171561018c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929050505061025f565b005b6101e5610507565b005b6101ef610674565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61025d6004803603602081101561023157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061069d565b005b61026761088f565b73ffffffffffffffffffffffffffffffffffffffff16610285610674565b73ffffffffffffffffffffffffffffffffffffffff161461030e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8051825114610385576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f706172616d20657272000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b60005b8251811015610500578473ffffffffffffffffffffffffffffffffffffffff166323b872dd858584815181106103ba57fe5b60200260200101518585815181106103ce57fe5b60200260200101516040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561044657600080fd5b505af115801561045a573d6000803e3d6000fd5b505050506040513d602081101561047057600080fd5b81019080805190602001909291905050506104f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f7472616e73206661696c6564000000000000000000000000000000000000000081525060200191505060405180910390fd5b8080600101915050610388565b5050505050565b61050f61088f565b73ffffffffffffffffffffffffffffffffffffffff1661052d610674565b73ffffffffffffffffffffffffffffffffffffffff16146105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106a561088f565b73ffffffffffffffffffffffffffffffffffffffff166106c3610674565b73ffffffffffffffffffffffffffffffffffffffff161461074c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156107d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806108986026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60003390509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a2646970667358221220ea00d155213ff243857a504c02b83331a2839e3ff57cd5afa4d90a4a8017158564736f6c634300060c0033"

// DeployAirDropX deploys a new Ethereum contract, binding an instance of AirDropX to it.
func DeployAirDropX(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AirDropX, error) {
	parsed, err := abi.JSON(strings.NewReader(AirDropXABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AirDropXBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AirDropX{AirDropXCaller: AirDropXCaller{contract: contract}, AirDropXTransactor: AirDropXTransactor{contract: contract}, AirDropXFilterer: AirDropXFilterer{contract: contract}}, nil
}

// AirDropX is an auto generated Go binding around an Ethereum contract.
type AirDropX struct {
	AirDropXCaller     // Read-only binding to the contract
	AirDropXTransactor // Write-only binding to the contract
	AirDropXFilterer   // Log filterer for contract events
}

// AirDropXCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirDropXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirDropXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirDropXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirDropXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirDropXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirDropXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirDropXSession struct {
	Contract     *AirDropX         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirDropXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirDropXCallerSession struct {
	Contract *AirDropXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AirDropXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirDropXTransactorSession struct {
	Contract     *AirDropXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AirDropXRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirDropXRaw struct {
	Contract *AirDropX // Generic contract binding to access the raw methods on
}

// AirDropXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirDropXCallerRaw struct {
	Contract *AirDropXCaller // Generic read-only contract binding to access the raw methods on
}

// AirDropXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirDropXTransactorRaw struct {
	Contract *AirDropXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirDropX creates a new instance of AirDropX, bound to a specific deployed contract.
func NewAirDropX(address common.Address, backend bind.ContractBackend) (*AirDropX, error) {
	contract, err := bindAirDropX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AirDropX{AirDropXCaller: AirDropXCaller{contract: contract}, AirDropXTransactor: AirDropXTransactor{contract: contract}, AirDropXFilterer: AirDropXFilterer{contract: contract}}, nil
}

// NewAirDropXCaller creates a new read-only instance of AirDropX, bound to a specific deployed contract.
func NewAirDropXCaller(address common.Address, caller bind.ContractCaller) (*AirDropXCaller, error) {
	contract, err := bindAirDropX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirDropXCaller{contract: contract}, nil
}

// NewAirDropXTransactor creates a new write-only instance of AirDropX, bound to a specific deployed contract.
func NewAirDropXTransactor(address common.Address, transactor bind.ContractTransactor) (*AirDropXTransactor, error) {
	contract, err := bindAirDropX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirDropXTransactor{contract: contract}, nil
}

// NewAirDropXFilterer creates a new log filterer instance of AirDropX, bound to a specific deployed contract.
func NewAirDropXFilterer(address common.Address, filterer bind.ContractFilterer) (*AirDropXFilterer, error) {
	contract, err := bindAirDropX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirDropXFilterer{contract: contract}, nil
}

// bindAirDropX binds a generic wrapper to an already deployed contract.
func bindAirDropX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirDropXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AirDropX *AirDropXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AirDropX.Contract.AirDropXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AirDropX *AirDropXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirDropX.Contract.AirDropXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AirDropX *AirDropXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AirDropX.Contract.AirDropXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AirDropX *AirDropXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AirDropX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AirDropX *AirDropXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirDropX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AirDropX *AirDropXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AirDropX.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AirDropX *AirDropXCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AirDropX.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AirDropX *AirDropXSession) Owner() (common.Address, error) {
	return _AirDropX.Contract.Owner(&_AirDropX.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AirDropX *AirDropXCallerSession) Owner() (common.Address, error) {
	return _AirDropX.Contract.Owner(&_AirDropX.CallOpts)
}

// HelpTransfer is a paid mutator transaction binding the contract method 0x534881c3.
//
// Solidity: function HelpTransfer(address tokenAddress, address from, address[] toList, uint256[] amountList) returns()
func (_AirDropX *AirDropXTransactor) HelpTransfer(opts *bind.TransactOpts, tokenAddress common.Address, from common.Address, toList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _AirDropX.contract.Transact(opts, "HelpTransfer", tokenAddress, from, toList, amountList)
}

// HelpTransfer is a paid mutator transaction binding the contract method 0x534881c3.
//
// Solidity: function HelpTransfer(address tokenAddress, address from, address[] toList, uint256[] amountList) returns()
func (_AirDropX *AirDropXSession) HelpTransfer(tokenAddress common.Address, from common.Address, toList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _AirDropX.Contract.HelpTransfer(&_AirDropX.TransactOpts, tokenAddress, from, toList, amountList)
}

// HelpTransfer is a paid mutator transaction binding the contract method 0x534881c3.
//
// Solidity: function HelpTransfer(address tokenAddress, address from, address[] toList, uint256[] amountList) returns()
func (_AirDropX *AirDropXTransactorSession) HelpTransfer(tokenAddress common.Address, from common.Address, toList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _AirDropX.Contract.HelpTransfer(&_AirDropX.TransactOpts, tokenAddress, from, toList, amountList)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AirDropX *AirDropXTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirDropX.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AirDropX *AirDropXSession) RenounceOwnership() (*types.Transaction, error) {
	return _AirDropX.Contract.RenounceOwnership(&_AirDropX.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AirDropX *AirDropXTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AirDropX.Contract.RenounceOwnership(&_AirDropX.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AirDropX *AirDropXTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AirDropX.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AirDropX *AirDropXSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AirDropX.Contract.TransferOwnership(&_AirDropX.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AirDropX *AirDropXTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AirDropX.Contract.TransferOwnership(&_AirDropX.TransactOpts, newOwner)
}

// AirDropXOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AirDropX contract.
type AirDropXOwnershipTransferredIterator struct {
	Event *AirDropXOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AirDropXOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirDropXOwnershipTransferred)
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
		it.Event = new(AirDropXOwnershipTransferred)
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
func (it *AirDropXOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirDropXOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirDropXOwnershipTransferred represents a OwnershipTransferred event raised by the AirDropX contract.
type AirDropXOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AirDropX *AirDropXFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AirDropXOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AirDropX.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AirDropXOwnershipTransferredIterator{contract: _AirDropX.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AirDropX *AirDropXFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AirDropXOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AirDropX.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirDropXOwnershipTransferred)
				if err := _AirDropX.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AirDropX *AirDropXFilterer) ParseOwnershipTransferred(log types.Log) (*AirDropXOwnershipTransferred, error) {
	event := new(AirDropXOwnershipTransferred)
	if err := _AirDropX.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
