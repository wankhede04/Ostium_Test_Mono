// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bet

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

// BetMetaData contains all meta data concerning the Bet contract.
var BetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_USDC\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closePrice\",\"type\":\"uint256\"}],\"name\":\"BetClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"}],\"name\":\"BetJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"betAmount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"closingTime\",\"type\":\"uint32\"}],\"name\":\"BetOpened\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"btcOpenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"betAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"closingTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"}],\"name\":\"closeBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"contractEscrow\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"}],\"name\":\"getBetInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"btcOpenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"betAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"closingTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"}],\"name\":\"joinBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_betAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_expirationTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_closingTime\",\"type\":\"uint32\"}],\"name\":\"openBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BetABI is the input ABI used to generate the binding from.
// Deprecated: Use BetMetaData.ABI instead.
var BetABI = BetMetaData.ABI

// Bet is an auto generated Go binding around an Ethereum contract.
type Bet struct {
	BetCaller     // Read-only binding to the contract
	BetTransactor // Write-only binding to the contract
	BetFilterer   // Log filterer for contract events
}

// BetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BetSession struct {
	Contract     *Bet              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BetCallerSession struct {
	Contract *BetCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BetTransactorSession struct {
	Contract     *BetTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BetRaw struct {
	Contract *Bet // Generic contract binding to access the raw methods on
}

// BetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BetCallerRaw struct {
	Contract *BetCaller // Generic read-only contract binding to access the raw methods on
}

// BetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BetTransactorRaw struct {
	Contract *BetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBet creates a new instance of Bet, bound to a specific deployed contract.
func NewBet(address common.Address, backend bind.ContractBackend) (*Bet, error) {
	contract, err := bindBet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bet{BetCaller: BetCaller{contract: contract}, BetTransactor: BetTransactor{contract: contract}, BetFilterer: BetFilterer{contract: contract}}, nil
}

// NewBetCaller creates a new read-only instance of Bet, bound to a specific deployed contract.
func NewBetCaller(address common.Address, caller bind.ContractCaller) (*BetCaller, error) {
	contract, err := bindBet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BetCaller{contract: contract}, nil
}

// NewBetTransactor creates a new write-only instance of Bet, bound to a specific deployed contract.
func NewBetTransactor(address common.Address, transactor bind.ContractTransactor) (*BetTransactor, error) {
	contract, err := bindBet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BetTransactor{contract: contract}, nil
}

// NewBetFilterer creates a new log filterer instance of Bet, bound to a specific deployed contract.
func NewBetFilterer(address common.Address, filterer bind.ContractFilterer) (*BetFilterer, error) {
	contract, err := bindBet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BetFilterer{contract: contract}, nil
}

// bindBet binds a generic wrapper to an already deployed contract.
func bindBet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bet *BetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bet.Contract.BetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bet *BetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bet.Contract.BetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bet *BetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bet.Contract.BetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bet *BetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bet *BetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bet *BetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bet.Contract.contract.Transact(opts, method, params...)
}

// Bets is a free data retrieval call binding the contract method 0x22af00fa.
//
// Solidity: function bets(uint256 ) view returns(address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active)
func (_Bet *BetCaller) Bets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UserA          common.Address
	UserB          common.Address
	BtcOpenPrice   *big.Int
	BetAmount      uint64
	ExpirationTime uint32
	ClosingTime    uint32
	IsLong         bool
	Active         bool
}, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "bets", arg0)

	outstruct := new(struct {
		UserA          common.Address
		UserB          common.Address
		BtcOpenPrice   *big.Int
		BetAmount      uint64
		ExpirationTime uint32
		ClosingTime    uint32
		IsLong         bool
		Active         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.UserB = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BtcOpenPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BetAmount = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.ExpirationTime = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ClosingTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.IsLong = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Active = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Bets is a free data retrieval call binding the contract method 0x22af00fa.
//
// Solidity: function bets(uint256 ) view returns(address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active)
func (_Bet *BetSession) Bets(arg0 *big.Int) (struct {
	UserA          common.Address
	UserB          common.Address
	BtcOpenPrice   *big.Int
	BetAmount      uint64
	ExpirationTime uint32
	ClosingTime    uint32
	IsLong         bool
	Active         bool
}, error) {
	return _Bet.Contract.Bets(&_Bet.CallOpts, arg0)
}

// Bets is a free data retrieval call binding the contract method 0x22af00fa.
//
// Solidity: function bets(uint256 ) view returns(address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active)
func (_Bet *BetCallerSession) Bets(arg0 *big.Int) (struct {
	UserA          common.Address
	UserB          common.Address
	BtcOpenPrice   *big.Int
	BetAmount      uint64
	ExpirationTime uint32
	ClosingTime    uint32
	IsLong         bool
	Active         bool
}, error) {
	return _Bet.Contract.Bets(&_Bet.CallOpts, arg0)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Bet *BetCaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Bet *BetSession) Escrow() (common.Address, error) {
	return _Bet.Contract.Escrow(&_Bet.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Bet *BetCallerSession) Escrow() (common.Address, error) {
	return _Bet.Contract.Escrow(&_Bet.CallOpts)
}

// GetBetInfo is a free data retrieval call binding the contract method 0x79141f80.
//
// Solidity: function getBetInfo(uint256 betIndex) view returns(address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active)
func (_Bet *BetCaller) GetBetInfo(opts *bind.CallOpts, betIndex *big.Int) (struct {
	UserA          common.Address
	UserB          common.Address
	BtcOpenPrice   *big.Int
	BetAmount      uint64
	ExpirationTime uint32
	ClosingTime    uint32
	IsLong         bool
	Active         bool
}, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "getBetInfo", betIndex)

	outstruct := new(struct {
		UserA          common.Address
		UserB          common.Address
		BtcOpenPrice   *big.Int
		BetAmount      uint64
		ExpirationTime uint32
		ClosingTime    uint32
		IsLong         bool
		Active         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.UserB = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BtcOpenPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BetAmount = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.ExpirationTime = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ClosingTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.IsLong = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Active = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// GetBetInfo is a free data retrieval call binding the contract method 0x79141f80.
//
// Solidity: function getBetInfo(uint256 betIndex) view returns(address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active)
func (_Bet *BetSession) GetBetInfo(betIndex *big.Int) (struct {
	UserA          common.Address
	UserB          common.Address
	BtcOpenPrice   *big.Int
	BetAmount      uint64
	ExpirationTime uint32
	ClosingTime    uint32
	IsLong         bool
	Active         bool
}, error) {
	return _Bet.Contract.GetBetInfo(&_Bet.CallOpts, betIndex)
}

// GetBetInfo is a free data retrieval call binding the contract method 0x79141f80.
//
// Solidity: function getBetInfo(uint256 betIndex) view returns(address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active)
func (_Bet *BetCallerSession) GetBetInfo(betIndex *big.Int) (struct {
	UserA          common.Address
	UserB          common.Address
	BtcOpenPrice   *big.Int
	BetAmount      uint64
	ExpirationTime uint32
	ClosingTime    uint32
	IsLong         bool
	Active         bool
}, error) {
	return _Bet.Contract.GetBetInfo(&_Bet.CallOpts, betIndex)
}

// TotalBets is a free data retrieval call binding the contract method 0xbefa1e2f.
//
// Solidity: function totalBets() view returns(uint256)
func (_Bet *BetCaller) TotalBets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "totalBets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBets is a free data retrieval call binding the contract method 0xbefa1e2f.
//
// Solidity: function totalBets() view returns(uint256)
func (_Bet *BetSession) TotalBets() (*big.Int, error) {
	return _Bet.Contract.TotalBets(&_Bet.CallOpts)
}

// TotalBets is a free data retrieval call binding the contract method 0xbefa1e2f.
//
// Solidity: function totalBets() view returns(uint256)
func (_Bet *BetCallerSession) TotalBets() (*big.Int, error) {
	return _Bet.Contract.TotalBets(&_Bet.CallOpts)
}

// CloseBet is a paid mutator transaction binding the contract method 0xa24beff7.
//
// Solidity: function closeBet(uint256 betIndex) returns()
func (_Bet *BetTransactor) CloseBet(opts *bind.TransactOpts, betIndex *big.Int) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "closeBet", betIndex)
}

// CloseBet is a paid mutator transaction binding the contract method 0xa24beff7.
//
// Solidity: function closeBet(uint256 betIndex) returns()
func (_Bet *BetSession) CloseBet(betIndex *big.Int) (*types.Transaction, error) {
	return _Bet.Contract.CloseBet(&_Bet.TransactOpts, betIndex)
}

// CloseBet is a paid mutator transaction binding the contract method 0xa24beff7.
//
// Solidity: function closeBet(uint256 betIndex) returns()
func (_Bet *BetTransactorSession) CloseBet(betIndex *big.Int) (*types.Transaction, error) {
	return _Bet.Contract.CloseBet(&_Bet.TransactOpts, betIndex)
}

// JoinBet is a paid mutator transaction binding the contract method 0x39fb3f45.
//
// Solidity: function joinBet(uint256 betIndex) returns()
func (_Bet *BetTransactor) JoinBet(opts *bind.TransactOpts, betIndex *big.Int) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "joinBet", betIndex)
}

// JoinBet is a paid mutator transaction binding the contract method 0x39fb3f45.
//
// Solidity: function joinBet(uint256 betIndex) returns()
func (_Bet *BetSession) JoinBet(betIndex *big.Int) (*types.Transaction, error) {
	return _Bet.Contract.JoinBet(&_Bet.TransactOpts, betIndex)
}

// JoinBet is a paid mutator transaction binding the contract method 0x39fb3f45.
//
// Solidity: function joinBet(uint256 betIndex) returns()
func (_Bet *BetTransactorSession) JoinBet(betIndex *big.Int) (*types.Transaction, error) {
	return _Bet.Contract.JoinBet(&_Bet.TransactOpts, betIndex)
}

// OpenBet is a paid mutator transaction binding the contract method 0x7169bf9a.
//
// Solidity: function openBet(bool _isLong, uint64 _betAmount, uint32 _expirationTime, uint32 _closingTime) returns()
func (_Bet *BetTransactor) OpenBet(opts *bind.TransactOpts, _isLong bool, _betAmount uint64, _expirationTime uint32, _closingTime uint32) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "openBet", _isLong, _betAmount, _expirationTime, _closingTime)
}

// OpenBet is a paid mutator transaction binding the contract method 0x7169bf9a.
//
// Solidity: function openBet(bool _isLong, uint64 _betAmount, uint32 _expirationTime, uint32 _closingTime) returns()
func (_Bet *BetSession) OpenBet(_isLong bool, _betAmount uint64, _expirationTime uint32, _closingTime uint32) (*types.Transaction, error) {
	return _Bet.Contract.OpenBet(&_Bet.TransactOpts, _isLong, _betAmount, _expirationTime, _closingTime)
}

// OpenBet is a paid mutator transaction binding the contract method 0x7169bf9a.
//
// Solidity: function openBet(bool _isLong, uint64 _betAmount, uint32 _expirationTime, uint32 _closingTime) returns()
func (_Bet *BetTransactorSession) OpenBet(_isLong bool, _betAmount uint64, _expirationTime uint32, _closingTime uint32) (*types.Transaction, error) {
	return _Bet.Contract.OpenBet(&_Bet.TransactOpts, _isLong, _betAmount, _expirationTime, _closingTime)
}

// BetBetClosedIterator is returned from FilterBetClosed and is used to iterate over the raw logs and unpacked data for BetClosed events raised by the Bet contract.
type BetBetClosedIterator struct {
	Event *BetBetClosed // Event containing the contract specifics and raw log

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
func (it *BetBetClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetBetClosed)
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
		it.Event = new(BetBetClosed)
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
func (it *BetBetClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetBetClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetBetClosed represents a BetClosed event raised by the Bet contract.
type BetBetClosed struct {
	BetIndex   *big.Int
	Winner     common.Address
	ClosePrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBetClosed is a free log retrieval operation binding the contract event 0x7f87517059366dbdc85d3ed06c99a243c3f704a883ab3d056a97fed73016983b.
//
// Solidity: event BetClosed(uint256 betIndex, address indexed winner, uint256 closePrice)
func (_Bet *BetFilterer) FilterBetClosed(opts *bind.FilterOpts, winner []common.Address) (*BetBetClosedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Bet.contract.FilterLogs(opts, "BetClosed", winnerRule)
	if err != nil {
		return nil, err
	}
	return &BetBetClosedIterator{contract: _Bet.contract, event: "BetClosed", logs: logs, sub: sub}, nil
}

// WatchBetClosed is a free log subscription operation binding the contract event 0x7f87517059366dbdc85d3ed06c99a243c3f704a883ab3d056a97fed73016983b.
//
// Solidity: event BetClosed(uint256 betIndex, address indexed winner, uint256 closePrice)
func (_Bet *BetFilterer) WatchBetClosed(opts *bind.WatchOpts, sink chan<- *BetBetClosed, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Bet.contract.WatchLogs(opts, "BetClosed", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetBetClosed)
				if err := _Bet.contract.UnpackLog(event, "BetClosed", log); err != nil {
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

// ParseBetClosed is a log parse operation binding the contract event 0x7f87517059366dbdc85d3ed06c99a243c3f704a883ab3d056a97fed73016983b.
//
// Solidity: event BetClosed(uint256 betIndex, address indexed winner, uint256 closePrice)
func (_Bet *BetFilterer) ParseBetClosed(log types.Log) (*BetBetClosed, error) {
	event := new(BetBetClosed)
	if err := _Bet.contract.UnpackLog(event, "BetClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetBetJoinedIterator is returned from FilterBetJoined and is used to iterate over the raw logs and unpacked data for BetJoined events raised by the Bet contract.
type BetBetJoinedIterator struct {
	Event *BetBetJoined // Event containing the contract specifics and raw log

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
func (it *BetBetJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetBetJoined)
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
		it.Event = new(BetBetJoined)
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
func (it *BetBetJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetBetJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetBetJoined represents a BetJoined event raised by the Bet contract.
type BetBetJoined struct {
	BetIndex  *big.Int
	UserB     common.Address
	OpenPrice *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBetJoined is a free log retrieval operation binding the contract event 0x178c34f2ab8b20fd561010b42eb45708692aa14a05fdf45525fd2f0968277c5e.
//
// Solidity: event BetJoined(uint256 betIndex, address indexed userB, uint256 openPrice)
func (_Bet *BetFilterer) FilterBetJoined(opts *bind.FilterOpts, userB []common.Address) (*BetBetJoinedIterator, error) {

	var userBRule []interface{}
	for _, userBItem := range userB {
		userBRule = append(userBRule, userBItem)
	}

	logs, sub, err := _Bet.contract.FilterLogs(opts, "BetJoined", userBRule)
	if err != nil {
		return nil, err
	}
	return &BetBetJoinedIterator{contract: _Bet.contract, event: "BetJoined", logs: logs, sub: sub}, nil
}

// WatchBetJoined is a free log subscription operation binding the contract event 0x178c34f2ab8b20fd561010b42eb45708692aa14a05fdf45525fd2f0968277c5e.
//
// Solidity: event BetJoined(uint256 betIndex, address indexed userB, uint256 openPrice)
func (_Bet *BetFilterer) WatchBetJoined(opts *bind.WatchOpts, sink chan<- *BetBetJoined, userB []common.Address) (event.Subscription, error) {

	var userBRule []interface{}
	for _, userBItem := range userB {
		userBRule = append(userBRule, userBItem)
	}

	logs, sub, err := _Bet.contract.WatchLogs(opts, "BetJoined", userBRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetBetJoined)
				if err := _Bet.contract.UnpackLog(event, "BetJoined", log); err != nil {
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

// ParseBetJoined is a log parse operation binding the contract event 0x178c34f2ab8b20fd561010b42eb45708692aa14a05fdf45525fd2f0968277c5e.
//
// Solidity: event BetJoined(uint256 betIndex, address indexed userB, uint256 openPrice)
func (_Bet *BetFilterer) ParseBetJoined(log types.Log) (*BetBetJoined, error) {
	event := new(BetBetJoined)
	if err := _Bet.contract.UnpackLog(event, "BetJoined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetBetOpenedIterator is returned from FilterBetOpened and is used to iterate over the raw logs and unpacked data for BetOpened events raised by the Bet contract.
type BetBetOpenedIterator struct {
	Event *BetBetOpened // Event containing the contract specifics and raw log

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
func (it *BetBetOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetBetOpened)
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
		it.Event = new(BetBetOpened)
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
func (it *BetBetOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetBetOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetBetOpened represents a BetOpened event raised by the Bet contract.
type BetBetOpened struct {
	BetIndex       *big.Int
	UserA          common.Address
	BetAmount      uint64
	IsLong         bool
	ExpirationTime uint32
	ClosingTime    uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBetOpened is a free log retrieval operation binding the contract event 0xc090ba65dabc5b8503d1d000190784700d638f232a9ec7f7c2ff3c76f45abade.
//
// Solidity: event BetOpened(uint256 betIndex, address indexed userA, uint64 betAmount, bool isLong, uint32 expirationTime, uint32 closingTime)
func (_Bet *BetFilterer) FilterBetOpened(opts *bind.FilterOpts, userA []common.Address) (*BetBetOpenedIterator, error) {

	var userARule []interface{}
	for _, userAItem := range userA {
		userARule = append(userARule, userAItem)
	}

	logs, sub, err := _Bet.contract.FilterLogs(opts, "BetOpened", userARule)
	if err != nil {
		return nil, err
	}
	return &BetBetOpenedIterator{contract: _Bet.contract, event: "BetOpened", logs: logs, sub: sub}, nil
}

// WatchBetOpened is a free log subscription operation binding the contract event 0xc090ba65dabc5b8503d1d000190784700d638f232a9ec7f7c2ff3c76f45abade.
//
// Solidity: event BetOpened(uint256 betIndex, address indexed userA, uint64 betAmount, bool isLong, uint32 expirationTime, uint32 closingTime)
func (_Bet *BetFilterer) WatchBetOpened(opts *bind.WatchOpts, sink chan<- *BetBetOpened, userA []common.Address) (event.Subscription, error) {

	var userARule []interface{}
	for _, userAItem := range userA {
		userARule = append(userARule, userAItem)
	}

	logs, sub, err := _Bet.contract.WatchLogs(opts, "BetOpened", userARule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetBetOpened)
				if err := _Bet.contract.UnpackLog(event, "BetOpened", log); err != nil {
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

// ParseBetOpened is a log parse operation binding the contract event 0xc090ba65dabc5b8503d1d000190784700d638f232a9ec7f7c2ff3c76f45abade.
//
// Solidity: event BetOpened(uint256 betIndex, address indexed userA, uint64 betAmount, bool isLong, uint32 expirationTime, uint32 closingTime)
func (_Bet *BetFilterer) ParseBetOpened(log types.Log) (*BetBetOpened, error) {
	event := new(BetBetOpened)
	if err := _Bet.contract.UnpackLog(event, "BetOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
