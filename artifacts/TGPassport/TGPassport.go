// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Passport

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

// TGPassportPassport is an auto generated low-level Go binding around an user-defined struct.
type TGPassportPassport struct {
	UserAddress      common.Address
	TgId             int64
	Valid            bool
	ValidatorAddress common.Address
	UserName         string
}

// PassportMetaData contains all meta data concerning the Passport contract.
var PassportMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"applyerTg\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet_address\",\"type\":\"address\"}],\"name\":\"passportApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int64\",\"name\":\"applyerTg\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet_address\",\"type\":\"address\"}],\"name\":\"passportAppliedIndexed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"applyerTg\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"passportApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"applyerTg\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"passportDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"from\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"int64\",\"name\":\"to\",\"type\":\"int64\"}],\"name\":\"trustedIndexed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"from\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"int64\",\"name\":\"to\",\"type\":\"int64\"}],\"name\":\"untrustedIndexed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"applyerTg\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"user_name_\",\"type\":\"string\"}],\"name\":\"ApplyForPassport\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"passportToApprove\",\"type\":\"address\"}],\"name\":\"ApprovePassport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"passportToDecline\",\"type\":\"address\"}],\"name\":\"DeclinePassport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"passportToDecline\",\"type\":\"address\"}],\"name\":\"DeletePassport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_wallet\",\"type\":\"address\"}],\"name\":\"GetPassportByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"tgId\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"}],\"internalType\":\"structTGPassport.Passport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"user_name_\",\"type\":\"string\"}],\"name\":\"GetPassportByNickName\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"tgId\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"}],\"internalType\":\"structTGPassport.Passport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"tgId_\",\"type\":\"int64\"}],\"name\":\"GetPassportByTgId\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"tgId\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"}],\"internalType\":\"structTGPassport.Passport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPassportFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"tgId_\",\"type\":\"int64\"}],\"name\":\"GetPassportWalletByID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_wallet\",\"type\":\"address\"}],\"name\":\"GetTgIdByAddress\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"tgid\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"from\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"to\",\"type\":\"int64\"}],\"name\":\"GetTrust\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"user_name_\",\"type\":\"string\"}],\"name\":\"GetWalletByNickName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"enemy_address\",\"type\":\"address\"}],\"name\":\"INOTTrustAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"my_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"enemy_name\",\"type\":\"string\"}],\"name\":\"INOTTrustUsername\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"from\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"to\",\"type\":\"int64\"}],\"name\":\"INotTrustID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"friend_address\",\"type\":\"address\"}],\"name\":\"ITrustAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"from\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"to\",\"type\":\"int64\"}],\"name\":\"ITrustID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"my_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"friend_name\",\"type\":\"string\"}],\"name\":\"ITrustUsername\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"passportFee_\",\"type\":\"uint256\"}],\"name\":\"SetPassportFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"new_user_name_\",\"type\":\"string\"}],\"name\":\"UpdateUserName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opinion_changed\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"passports\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"tgId\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"name\":\"tgIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"name\":\"trust_global\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"username_wallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PassportABI is the input ABI used to generate the binding from.
// Deprecated: Use PassportMetaData.ABI instead.
var PassportABI = PassportMetaData.ABI

// Passport is an auto generated Go binding around an Ethereum contract.
type Passport struct {
	PassportCaller     // Read-only binding to the contract
	PassportTransactor // Write-only binding to the contract
	PassportFilterer   // Log filterer for contract events
}

// PassportCaller is an auto generated read-only Go binding around an Ethereum contract.
type PassportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PassportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PassportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PassportSession struct {
	Contract     *Passport         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PassportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PassportCallerSession struct {
	Contract *PassportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PassportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PassportTransactorSession struct {
	Contract     *PassportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PassportRaw is an auto generated low-level Go binding around an Ethereum contract.
type PassportRaw struct {
	Contract *Passport // Generic contract binding to access the raw methods on
}

// PassportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PassportCallerRaw struct {
	Contract *PassportCaller // Generic read-only contract binding to access the raw methods on
}

// PassportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PassportTransactorRaw struct {
	Contract *PassportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPassport creates a new instance of Passport, bound to a specific deployed contract.
func NewPassport(address common.Address, backend bind.ContractBackend) (*Passport, error) {
	contract, err := bindPassport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Passport{PassportCaller: PassportCaller{contract: contract}, PassportTransactor: PassportTransactor{contract: contract}, PassportFilterer: PassportFilterer{contract: contract}}, nil
}

// NewPassportCaller creates a new read-only instance of Passport, bound to a specific deployed contract.
func NewPassportCaller(address common.Address, caller bind.ContractCaller) (*PassportCaller, error) {
	contract, err := bindPassport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PassportCaller{contract: contract}, nil
}

// NewPassportTransactor creates a new write-only instance of Passport, bound to a specific deployed contract.
func NewPassportTransactor(address common.Address, transactor bind.ContractTransactor) (*PassportTransactor, error) {
	contract, err := bindPassport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PassportTransactor{contract: contract}, nil
}

// NewPassportFilterer creates a new log filterer instance of Passport, bound to a specific deployed contract.
func NewPassportFilterer(address common.Address, filterer bind.ContractFilterer) (*PassportFilterer, error) {
	contract, err := bindPassport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PassportFilterer{contract: contract}, nil
}

// bindPassport binds a generic wrapper to an already deployed contract.
func bindPassport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Passport *PassportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Passport.Contract.PassportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Passport *PassportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Passport.Contract.PassportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Passport *PassportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Passport.Contract.PassportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Passport *PassportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Passport.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Passport *PassportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Passport.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Passport *PassportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Passport.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() view returns(address)
func (_Passport *PassportCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() view returns(address)
func (_Passport *PassportSession) GetOwner() (common.Address, error) {
	return _Passport.Contract.GetOwner(&_Passport.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() view returns(address)
func (_Passport *PassportCallerSession) GetOwner() (common.Address, error) {
	return _Passport.Contract.GetOwner(&_Passport.CallOpts)
}

// GetPassportByAddress is a free data retrieval call binding the contract method 0x1423ea07.
//
// Solidity: function GetPassportByAddress(address user_wallet) view returns((address,int64,bool,address,string))
func (_Passport *PassportCaller) GetPassportByAddress(opts *bind.CallOpts, user_wallet common.Address) (TGPassportPassport, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetPassportByAddress", user_wallet)

	if err != nil {
		return *new(TGPassportPassport), err
	}

	out0 := *abi.ConvertType(out[0], new(TGPassportPassport)).(*TGPassportPassport)

	return out0, err

}

// GetPassportByAddress is a free data retrieval call binding the contract method 0x1423ea07.
//
// Solidity: function GetPassportByAddress(address user_wallet) view returns((address,int64,bool,address,string))
func (_Passport *PassportSession) GetPassportByAddress(user_wallet common.Address) (TGPassportPassport, error) {
	return _Passport.Contract.GetPassportByAddress(&_Passport.CallOpts, user_wallet)
}

// GetPassportByAddress is a free data retrieval call binding the contract method 0x1423ea07.
//
// Solidity: function GetPassportByAddress(address user_wallet) view returns((address,int64,bool,address,string))
func (_Passport *PassportCallerSession) GetPassportByAddress(user_wallet common.Address) (TGPassportPassport, error) {
	return _Passport.Contract.GetPassportByAddress(&_Passport.CallOpts, user_wallet)
}

// GetPassportByNickName is a free data retrieval call binding the contract method 0xb25047c2.
//
// Solidity: function GetPassportByNickName(string user_name_) view returns((address,int64,bool,address,string))
func (_Passport *PassportCaller) GetPassportByNickName(opts *bind.CallOpts, user_name_ string) (TGPassportPassport, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetPassportByNickName", user_name_)

	if err != nil {
		return *new(TGPassportPassport), err
	}

	out0 := *abi.ConvertType(out[0], new(TGPassportPassport)).(*TGPassportPassport)

	return out0, err

}

// GetPassportByNickName is a free data retrieval call binding the contract method 0xb25047c2.
//
// Solidity: function GetPassportByNickName(string user_name_) view returns((address,int64,bool,address,string))
func (_Passport *PassportSession) GetPassportByNickName(user_name_ string) (TGPassportPassport, error) {
	return _Passport.Contract.GetPassportByNickName(&_Passport.CallOpts, user_name_)
}

// GetPassportByNickName is a free data retrieval call binding the contract method 0xb25047c2.
//
// Solidity: function GetPassportByNickName(string user_name_) view returns((address,int64,bool,address,string))
func (_Passport *PassportCallerSession) GetPassportByNickName(user_name_ string) (TGPassportPassport, error) {
	return _Passport.Contract.GetPassportByNickName(&_Passport.CallOpts, user_name_)
}

// GetPassportByTgId is a free data retrieval call binding the contract method 0x65a5637e.
//
// Solidity: function GetPassportByTgId(int64 tgId_) view returns((address,int64,bool,address,string))
func (_Passport *PassportCaller) GetPassportByTgId(opts *bind.CallOpts, tgId_ int64) (TGPassportPassport, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetPassportByTgId", tgId_)

	if err != nil {
		return *new(TGPassportPassport), err
	}

	out0 := *abi.ConvertType(out[0], new(TGPassportPassport)).(*TGPassportPassport)

	return out0, err

}

// GetPassportByTgId is a free data retrieval call binding the contract method 0x65a5637e.
//
// Solidity: function GetPassportByTgId(int64 tgId_) view returns((address,int64,bool,address,string))
func (_Passport *PassportSession) GetPassportByTgId(tgId_ int64) (TGPassportPassport, error) {
	return _Passport.Contract.GetPassportByTgId(&_Passport.CallOpts, tgId_)
}

// GetPassportByTgId is a free data retrieval call binding the contract method 0x65a5637e.
//
// Solidity: function GetPassportByTgId(int64 tgId_) view returns((address,int64,bool,address,string))
func (_Passport *PassportCallerSession) GetPassportByTgId(tgId_ int64) (TGPassportPassport, error) {
	return _Passport.Contract.GetPassportByTgId(&_Passport.CallOpts, tgId_)
}

// GetPassportFee is a free data retrieval call binding the contract method 0x48dc1561.
//
// Solidity: function GetPassportFee() view returns(uint256)
func (_Passport *PassportCaller) GetPassportFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetPassportFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPassportFee is a free data retrieval call binding the contract method 0x48dc1561.
//
// Solidity: function GetPassportFee() view returns(uint256)
func (_Passport *PassportSession) GetPassportFee() (*big.Int, error) {
	return _Passport.Contract.GetPassportFee(&_Passport.CallOpts)
}

// GetPassportFee is a free data retrieval call binding the contract method 0x48dc1561.
//
// Solidity: function GetPassportFee() view returns(uint256)
func (_Passport *PassportCallerSession) GetPassportFee() (*big.Int, error) {
	return _Passport.Contract.GetPassportFee(&_Passport.CallOpts)
}

// GetPassportWalletByID is a free data retrieval call binding the contract method 0x0c2fc6e3.
//
// Solidity: function GetPassportWalletByID(int64 tgId_) view returns(address)
func (_Passport *PassportCaller) GetPassportWalletByID(opts *bind.CallOpts, tgId_ int64) (common.Address, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetPassportWalletByID", tgId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPassportWalletByID is a free data retrieval call binding the contract method 0x0c2fc6e3.
//
// Solidity: function GetPassportWalletByID(int64 tgId_) view returns(address)
func (_Passport *PassportSession) GetPassportWalletByID(tgId_ int64) (common.Address, error) {
	return _Passport.Contract.GetPassportWalletByID(&_Passport.CallOpts, tgId_)
}

// GetPassportWalletByID is a free data retrieval call binding the contract method 0x0c2fc6e3.
//
// Solidity: function GetPassportWalletByID(int64 tgId_) view returns(address)
func (_Passport *PassportCallerSession) GetPassportWalletByID(tgId_ int64) (common.Address, error) {
	return _Passport.Contract.GetPassportWalletByID(&_Passport.CallOpts, tgId_)
}

// GetTgIdByAddress is a free data retrieval call binding the contract method 0x61cabe37.
//
// Solidity: function GetTgIdByAddress(address user_wallet) view returns(int64 tgid)
func (_Passport *PassportCaller) GetTgIdByAddress(opts *bind.CallOpts, user_wallet common.Address) (int64, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetTgIdByAddress", user_wallet)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetTgIdByAddress is a free data retrieval call binding the contract method 0x61cabe37.
//
// Solidity: function GetTgIdByAddress(address user_wallet) view returns(int64 tgid)
func (_Passport *PassportSession) GetTgIdByAddress(user_wallet common.Address) (int64, error) {
	return _Passport.Contract.GetTgIdByAddress(&_Passport.CallOpts, user_wallet)
}

// GetTgIdByAddress is a free data retrieval call binding the contract method 0x61cabe37.
//
// Solidity: function GetTgIdByAddress(address user_wallet) view returns(int64 tgid)
func (_Passport *PassportCallerSession) GetTgIdByAddress(user_wallet common.Address) (int64, error) {
	return _Passport.Contract.GetTgIdByAddress(&_Passport.CallOpts, user_wallet)
}

// GetTrust is a free data retrieval call binding the contract method 0x4ac054bb.
//
// Solidity: function GetTrust(int64 from, int64 to) view returns(bool)
func (_Passport *PassportCaller) GetTrust(opts *bind.CallOpts, from int64, to int64) (bool, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetTrust", from, to)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTrust is a free data retrieval call binding the contract method 0x4ac054bb.
//
// Solidity: function GetTrust(int64 from, int64 to) view returns(bool)
func (_Passport *PassportSession) GetTrust(from int64, to int64) (bool, error) {
	return _Passport.Contract.GetTrust(&_Passport.CallOpts, from, to)
}

// GetTrust is a free data retrieval call binding the contract method 0x4ac054bb.
//
// Solidity: function GetTrust(int64 from, int64 to) view returns(bool)
func (_Passport *PassportCallerSession) GetTrust(from int64, to int64) (bool, error) {
	return _Passport.Contract.GetTrust(&_Passport.CallOpts, from, to)
}

// GetWalletByNickName is a free data retrieval call binding the contract method 0xd2114e18.
//
// Solidity: function GetWalletByNickName(string user_name_) view returns(address)
func (_Passport *PassportCaller) GetWalletByNickName(opts *bind.CallOpts, user_name_ string) (common.Address, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "GetWalletByNickName", user_name_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWalletByNickName is a free data retrieval call binding the contract method 0xd2114e18.
//
// Solidity: function GetWalletByNickName(string user_name_) view returns(address)
func (_Passport *PassportSession) GetWalletByNickName(user_name_ string) (common.Address, error) {
	return _Passport.Contract.GetWalletByNickName(&_Passport.CallOpts, user_name_)
}

// GetWalletByNickName is a free data retrieval call binding the contract method 0xd2114e18.
//
// Solidity: function GetWalletByNickName(string user_name_) view returns(address)
func (_Passport *PassportCallerSession) GetWalletByNickName(user_name_ string) (common.Address, error) {
	return _Passport.Contract.GetWalletByNickName(&_Passport.CallOpts, user_name_)
}

// OpinionChanged is a free data retrieval call binding the contract method 0x6a8ec05e.
//
// Solidity: function opinion_changed(int64 , uint256 ) view returns(int64)
func (_Passport *PassportCaller) OpinionChanged(opts *bind.CallOpts, arg0 int64, arg1 *big.Int) (int64, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "opinion_changed", arg0, arg1)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// OpinionChanged is a free data retrieval call binding the contract method 0x6a8ec05e.
//
// Solidity: function opinion_changed(int64 , uint256 ) view returns(int64)
func (_Passport *PassportSession) OpinionChanged(arg0 int64, arg1 *big.Int) (int64, error) {
	return _Passport.Contract.OpinionChanged(&_Passport.CallOpts, arg0, arg1)
}

// OpinionChanged is a free data retrieval call binding the contract method 0x6a8ec05e.
//
// Solidity: function opinion_changed(int64 , uint256 ) view returns(int64)
func (_Passport *PassportCallerSession) OpinionChanged(arg0 int64, arg1 *big.Int) (int64, error) {
	return _Passport.Contract.OpinionChanged(&_Passport.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Passport *PassportCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Passport *PassportSession) Owner() (common.Address, error) {
	return _Passport.Contract.Owner(&_Passport.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Passport *PassportCallerSession) Owner() (common.Address, error) {
	return _Passport.Contract.Owner(&_Passport.CallOpts)
}

// Passports is a free data retrieval call binding the contract method 0xe37c132b.
//
// Solidity: function passports(address ) view returns(address userAddress, int64 tgId, bool valid, address validatorAddress, string userName)
func (_Passport *PassportCaller) Passports(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserAddress      common.Address
	TgId             int64
	Valid            bool
	ValidatorAddress common.Address
	UserName         string
}, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "passports", arg0)

	outstruct := new(struct {
		UserAddress      common.Address
		TgId             int64
		Valid            bool
		ValidatorAddress common.Address
		UserName         string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TgId = *abi.ConvertType(out[1], new(int64)).(*int64)
	outstruct.Valid = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.UserName = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Passports is a free data retrieval call binding the contract method 0xe37c132b.
//
// Solidity: function passports(address ) view returns(address userAddress, int64 tgId, bool valid, address validatorAddress, string userName)
func (_Passport *PassportSession) Passports(arg0 common.Address) (struct {
	UserAddress      common.Address
	TgId             int64
	Valid            bool
	ValidatorAddress common.Address
	UserName         string
}, error) {
	return _Passport.Contract.Passports(&_Passport.CallOpts, arg0)
}

// Passports is a free data retrieval call binding the contract method 0xe37c132b.
//
// Solidity: function passports(address ) view returns(address userAddress, int64 tgId, bool valid, address validatorAddress, string userName)
func (_Passport *PassportCallerSession) Passports(arg0 common.Address) (struct {
	UserAddress      common.Address
	TgId             int64
	Valid            bool
	ValidatorAddress common.Address
	UserName         string
}, error) {
	return _Passport.Contract.Passports(&_Passport.CallOpts, arg0)
}

// TgIdToAddress is a free data retrieval call binding the contract method 0x9413c642.
//
// Solidity: function tgIdToAddress(int64 ) view returns(address)
func (_Passport *PassportCaller) TgIdToAddress(opts *bind.CallOpts, arg0 int64) (common.Address, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "tgIdToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TgIdToAddress is a free data retrieval call binding the contract method 0x9413c642.
//
// Solidity: function tgIdToAddress(int64 ) view returns(address)
func (_Passport *PassportSession) TgIdToAddress(arg0 int64) (common.Address, error) {
	return _Passport.Contract.TgIdToAddress(&_Passport.CallOpts, arg0)
}

// TgIdToAddress is a free data retrieval call binding the contract method 0x9413c642.
//
// Solidity: function tgIdToAddress(int64 ) view returns(address)
func (_Passport *PassportCallerSession) TgIdToAddress(arg0 int64) (common.Address, error) {
	return _Passport.Contract.TgIdToAddress(&_Passport.CallOpts, arg0)
}

// TrustGlobal is a free data retrieval call binding the contract method 0x239f60fa.
//
// Solidity: function trust_global(int64 , int64 ) view returns(bool)
func (_Passport *PassportCaller) TrustGlobal(opts *bind.CallOpts, arg0 int64, arg1 int64) (bool, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "trust_global", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrustGlobal is a free data retrieval call binding the contract method 0x239f60fa.
//
// Solidity: function trust_global(int64 , int64 ) view returns(bool)
func (_Passport *PassportSession) TrustGlobal(arg0 int64, arg1 int64) (bool, error) {
	return _Passport.Contract.TrustGlobal(&_Passport.CallOpts, arg0, arg1)
}

// TrustGlobal is a free data retrieval call binding the contract method 0x239f60fa.
//
// Solidity: function trust_global(int64 , int64 ) view returns(bool)
func (_Passport *PassportCallerSession) TrustGlobal(arg0 int64, arg1 int64) (bool, error) {
	return _Passport.Contract.TrustGlobal(&_Passport.CallOpts, arg0, arg1)
}

// UsernameWallets is a free data retrieval call binding the contract method 0x59cca93b.
//
// Solidity: function username_wallets(string ) view returns(address)
func (_Passport *PassportCaller) UsernameWallets(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Passport.contract.Call(opts, &out, "username_wallets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsernameWallets is a free data retrieval call binding the contract method 0x59cca93b.
//
// Solidity: function username_wallets(string ) view returns(address)
func (_Passport *PassportSession) UsernameWallets(arg0 string) (common.Address, error) {
	return _Passport.Contract.UsernameWallets(&_Passport.CallOpts, arg0)
}

// UsernameWallets is a free data retrieval call binding the contract method 0x59cca93b.
//
// Solidity: function username_wallets(string ) view returns(address)
func (_Passport *PassportCallerSession) UsernameWallets(arg0 string) (common.Address, error) {
	return _Passport.Contract.UsernameWallets(&_Passport.CallOpts, arg0)
}

// ApplyForPassport is a paid mutator transaction binding the contract method 0xd5998aeb.
//
// Solidity: function ApplyForPassport(int64 applyerTg, string user_name_) payable returns()
func (_Passport *PassportTransactor) ApplyForPassport(opts *bind.TransactOpts, applyerTg int64, user_name_ string) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "ApplyForPassport", applyerTg, user_name_)
}

// ApplyForPassport is a paid mutator transaction binding the contract method 0xd5998aeb.
//
// Solidity: function ApplyForPassport(int64 applyerTg, string user_name_) payable returns()
func (_Passport *PassportSession) ApplyForPassport(applyerTg int64, user_name_ string) (*types.Transaction, error) {
	return _Passport.Contract.ApplyForPassport(&_Passport.TransactOpts, applyerTg, user_name_)
}

// ApplyForPassport is a paid mutator transaction binding the contract method 0xd5998aeb.
//
// Solidity: function ApplyForPassport(int64 applyerTg, string user_name_) payable returns()
func (_Passport *PassportTransactorSession) ApplyForPassport(applyerTg int64, user_name_ string) (*types.Transaction, error) {
	return _Passport.Contract.ApplyForPassport(&_Passport.TransactOpts, applyerTg, user_name_)
}

// ApprovePassport is a paid mutator transaction binding the contract method 0x1755e9e6.
//
// Solidity: function ApprovePassport(address passportToApprove) returns()
func (_Passport *PassportTransactor) ApprovePassport(opts *bind.TransactOpts, passportToApprove common.Address) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "ApprovePassport", passportToApprove)
}

// ApprovePassport is a paid mutator transaction binding the contract method 0x1755e9e6.
//
// Solidity: function ApprovePassport(address passportToApprove) returns()
func (_Passport *PassportSession) ApprovePassport(passportToApprove common.Address) (*types.Transaction, error) {
	return _Passport.Contract.ApprovePassport(&_Passport.TransactOpts, passportToApprove)
}

// ApprovePassport is a paid mutator transaction binding the contract method 0x1755e9e6.
//
// Solidity: function ApprovePassport(address passportToApprove) returns()
func (_Passport *PassportTransactorSession) ApprovePassport(passportToApprove common.Address) (*types.Transaction, error) {
	return _Passport.Contract.ApprovePassport(&_Passport.TransactOpts, passportToApprove)
}

// DeclinePassport is a paid mutator transaction binding the contract method 0x9421da61.
//
// Solidity: function DeclinePassport(address passportToDecline) returns()
func (_Passport *PassportTransactor) DeclinePassport(opts *bind.TransactOpts, passportToDecline common.Address) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "DeclinePassport", passportToDecline)
}

// DeclinePassport is a paid mutator transaction binding the contract method 0x9421da61.
//
// Solidity: function DeclinePassport(address passportToDecline) returns()
func (_Passport *PassportSession) DeclinePassport(passportToDecline common.Address) (*types.Transaction, error) {
	return _Passport.Contract.DeclinePassport(&_Passport.TransactOpts, passportToDecline)
}

// DeclinePassport is a paid mutator transaction binding the contract method 0x9421da61.
//
// Solidity: function DeclinePassport(address passportToDecline) returns()
func (_Passport *PassportTransactorSession) DeclinePassport(passportToDecline common.Address) (*types.Transaction, error) {
	return _Passport.Contract.DeclinePassport(&_Passport.TransactOpts, passportToDecline)
}

// DeletePassport is a paid mutator transaction binding the contract method 0x65104afc.
//
// Solidity: function DeletePassport(address passportToDecline) returns()
func (_Passport *PassportTransactor) DeletePassport(opts *bind.TransactOpts, passportToDecline common.Address) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "DeletePassport", passportToDecline)
}

// DeletePassport is a paid mutator transaction binding the contract method 0x65104afc.
//
// Solidity: function DeletePassport(address passportToDecline) returns()
func (_Passport *PassportSession) DeletePassport(passportToDecline common.Address) (*types.Transaction, error) {
	return _Passport.Contract.DeletePassport(&_Passport.TransactOpts, passportToDecline)
}

// DeletePassport is a paid mutator transaction binding the contract method 0x65104afc.
//
// Solidity: function DeletePassport(address passportToDecline) returns()
func (_Passport *PassportTransactorSession) DeletePassport(passportToDecline common.Address) (*types.Transaction, error) {
	return _Passport.Contract.DeletePassport(&_Passport.TransactOpts, passportToDecline)
}

// INOTTrustAddress is a paid mutator transaction binding the contract method 0x5ddf30d0.
//
// Solidity: function INOTTrustAddress(address enemy_address) returns()
func (_Passport *PassportTransactor) INOTTrustAddress(opts *bind.TransactOpts, enemy_address common.Address) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "INOTTrustAddress", enemy_address)
}

// INOTTrustAddress is a paid mutator transaction binding the contract method 0x5ddf30d0.
//
// Solidity: function INOTTrustAddress(address enemy_address) returns()
func (_Passport *PassportSession) INOTTrustAddress(enemy_address common.Address) (*types.Transaction, error) {
	return _Passport.Contract.INOTTrustAddress(&_Passport.TransactOpts, enemy_address)
}

// INOTTrustAddress is a paid mutator transaction binding the contract method 0x5ddf30d0.
//
// Solidity: function INOTTrustAddress(address enemy_address) returns()
func (_Passport *PassportTransactorSession) INOTTrustAddress(enemy_address common.Address) (*types.Transaction, error) {
	return _Passport.Contract.INOTTrustAddress(&_Passport.TransactOpts, enemy_address)
}

// INOTTrustUsername is a paid mutator transaction binding the contract method 0x87234689.
//
// Solidity: function INOTTrustUsername(string my_username, string enemy_name) returns()
func (_Passport *PassportTransactor) INOTTrustUsername(opts *bind.TransactOpts, my_username string, enemy_name string) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "INOTTrustUsername", my_username, enemy_name)
}

// INOTTrustUsername is a paid mutator transaction binding the contract method 0x87234689.
//
// Solidity: function INOTTrustUsername(string my_username, string enemy_name) returns()
func (_Passport *PassportSession) INOTTrustUsername(my_username string, enemy_name string) (*types.Transaction, error) {
	return _Passport.Contract.INOTTrustUsername(&_Passport.TransactOpts, my_username, enemy_name)
}

// INOTTrustUsername is a paid mutator transaction binding the contract method 0x87234689.
//
// Solidity: function INOTTrustUsername(string my_username, string enemy_name) returns()
func (_Passport *PassportTransactorSession) INOTTrustUsername(my_username string, enemy_name string) (*types.Transaction, error) {
	return _Passport.Contract.INOTTrustUsername(&_Passport.TransactOpts, my_username, enemy_name)
}

// INotTrustID is a paid mutator transaction binding the contract method 0x5c962dfc.
//
// Solidity: function INotTrustID(int64 from, int64 to) returns()
func (_Passport *PassportTransactor) INotTrustID(opts *bind.TransactOpts, from int64, to int64) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "INotTrustID", from, to)
}

// INotTrustID is a paid mutator transaction binding the contract method 0x5c962dfc.
//
// Solidity: function INotTrustID(int64 from, int64 to) returns()
func (_Passport *PassportSession) INotTrustID(from int64, to int64) (*types.Transaction, error) {
	return _Passport.Contract.INotTrustID(&_Passport.TransactOpts, from, to)
}

// INotTrustID is a paid mutator transaction binding the contract method 0x5c962dfc.
//
// Solidity: function INotTrustID(int64 from, int64 to) returns()
func (_Passport *PassportTransactorSession) INotTrustID(from int64, to int64) (*types.Transaction, error) {
	return _Passport.Contract.INotTrustID(&_Passport.TransactOpts, from, to)
}

// ITrustAddress is a paid mutator transaction binding the contract method 0xe5122bab.
//
// Solidity: function ITrustAddress(address friend_address) returns()
func (_Passport *PassportTransactor) ITrustAddress(opts *bind.TransactOpts, friend_address common.Address) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "ITrustAddress", friend_address)
}

// ITrustAddress is a paid mutator transaction binding the contract method 0xe5122bab.
//
// Solidity: function ITrustAddress(address friend_address) returns()
func (_Passport *PassportSession) ITrustAddress(friend_address common.Address) (*types.Transaction, error) {
	return _Passport.Contract.ITrustAddress(&_Passport.TransactOpts, friend_address)
}

// ITrustAddress is a paid mutator transaction binding the contract method 0xe5122bab.
//
// Solidity: function ITrustAddress(address friend_address) returns()
func (_Passport *PassportTransactorSession) ITrustAddress(friend_address common.Address) (*types.Transaction, error) {
	return _Passport.Contract.ITrustAddress(&_Passport.TransactOpts, friend_address)
}

// ITrustID is a paid mutator transaction binding the contract method 0xdca28884.
//
// Solidity: function ITrustID(int64 from, int64 to) returns()
func (_Passport *PassportTransactor) ITrustID(opts *bind.TransactOpts, from int64, to int64) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "ITrustID", from, to)
}

// ITrustID is a paid mutator transaction binding the contract method 0xdca28884.
//
// Solidity: function ITrustID(int64 from, int64 to) returns()
func (_Passport *PassportSession) ITrustID(from int64, to int64) (*types.Transaction, error) {
	return _Passport.Contract.ITrustID(&_Passport.TransactOpts, from, to)
}

// ITrustID is a paid mutator transaction binding the contract method 0xdca28884.
//
// Solidity: function ITrustID(int64 from, int64 to) returns()
func (_Passport *PassportTransactorSession) ITrustID(from int64, to int64) (*types.Transaction, error) {
	return _Passport.Contract.ITrustID(&_Passport.TransactOpts, from, to)
}

// ITrustUsername is a paid mutator transaction binding the contract method 0x68789cea.
//
// Solidity: function ITrustUsername(string my_username, string friend_name) returns()
func (_Passport *PassportTransactor) ITrustUsername(opts *bind.TransactOpts, my_username string, friend_name string) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "ITrustUsername", my_username, friend_name)
}

// ITrustUsername is a paid mutator transaction binding the contract method 0x68789cea.
//
// Solidity: function ITrustUsername(string my_username, string friend_name) returns()
func (_Passport *PassportSession) ITrustUsername(my_username string, friend_name string) (*types.Transaction, error) {
	return _Passport.Contract.ITrustUsername(&_Passport.TransactOpts, my_username, friend_name)
}

// ITrustUsername is a paid mutator transaction binding the contract method 0x68789cea.
//
// Solidity: function ITrustUsername(string my_username, string friend_name) returns()
func (_Passport *PassportTransactorSession) ITrustUsername(my_username string, friend_name string) (*types.Transaction, error) {
	return _Passport.Contract.ITrustUsername(&_Passport.TransactOpts, my_username, friend_name)
}

// SetPassportFee is a paid mutator transaction binding the contract method 0xd2125bd2.
//
// Solidity: function SetPassportFee(uint256 passportFee_) returns()
func (_Passport *PassportTransactor) SetPassportFee(opts *bind.TransactOpts, passportFee_ *big.Int) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "SetPassportFee", passportFee_)
}

// SetPassportFee is a paid mutator transaction binding the contract method 0xd2125bd2.
//
// Solidity: function SetPassportFee(uint256 passportFee_) returns()
func (_Passport *PassportSession) SetPassportFee(passportFee_ *big.Int) (*types.Transaction, error) {
	return _Passport.Contract.SetPassportFee(&_Passport.TransactOpts, passportFee_)
}

// SetPassportFee is a paid mutator transaction binding the contract method 0xd2125bd2.
//
// Solidity: function SetPassportFee(uint256 passportFee_) returns()
func (_Passport *PassportTransactorSession) SetPassportFee(passportFee_ *big.Int) (*types.Transaction, error) {
	return _Passport.Contract.SetPassportFee(&_Passport.TransactOpts, passportFee_)
}

// UpdateUserName is a paid mutator transaction binding the contract method 0x21e6531d.
//
// Solidity: function UpdateUserName(string new_user_name_) returns()
func (_Passport *PassportTransactor) UpdateUserName(opts *bind.TransactOpts, new_user_name_ string) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "UpdateUserName", new_user_name_)
}

// UpdateUserName is a paid mutator transaction binding the contract method 0x21e6531d.
//
// Solidity: function UpdateUserName(string new_user_name_) returns()
func (_Passport *PassportSession) UpdateUserName(new_user_name_ string) (*types.Transaction, error) {
	return _Passport.Contract.UpdateUserName(&_Passport.TransactOpts, new_user_name_)
}

// UpdateUserName is a paid mutator transaction binding the contract method 0x21e6531d.
//
// Solidity: function UpdateUserName(string new_user_name_) returns()
func (_Passport *PassportTransactorSession) UpdateUserName(new_user_name_ string) (*types.Transaction, error) {
	return _Passport.Contract.UpdateUserName(&_Passport.TransactOpts, new_user_name_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Passport *PassportTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Passport *PassportSession) RenounceOwnership() (*types.Transaction, error) {
	return _Passport.Contract.RenounceOwnership(&_Passport.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Passport *PassportTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Passport.Contract.RenounceOwnership(&_Passport.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Passport *PassportTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Passport.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Passport *PassportSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Passport.Contract.TransferOwnership(&_Passport.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Passport *PassportTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Passport.Contract.TransferOwnership(&_Passport.TransactOpts, newOwner)
}

// PassportOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Passport contract.
type PassportOwnershipTransferredIterator struct {
	Event *PassportOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PassportOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportOwnershipTransferred)
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
		it.Event = new(PassportOwnershipTransferred)
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
func (it *PassportOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportOwnershipTransferred represents a OwnershipTransferred event raised by the Passport contract.
type PassportOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Passport *PassportFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Passport.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportOwnershipTransferredIterator{contract: _Passport.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Passport *PassportFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PassportOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Passport.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportOwnershipTransferred)
				if err := _Passport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Passport *PassportFilterer) ParseOwnershipTransferred(log types.Log) (*PassportOwnershipTransferred, error) {
	event := new(PassportOwnershipTransferred)
	if err := _Passport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportPassportAppliedIterator is returned from FilterPassportApplied and is used to iterate over the raw logs and unpacked data for PassportApplied events raised by the Passport contract.
type PassportPassportAppliedIterator struct {
	Event *PassportPassportApplied // Event containing the contract specifics and raw log

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
func (it *PassportPassportAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportPassportApplied)
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
		it.Event = new(PassportPassportApplied)
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
func (it *PassportPassportAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportPassportAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportPassportApplied represents a PassportApplied event raised by the Passport contract.
type PassportPassportApplied struct {
	ApplyerTg     int64
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPassportApplied is a free log retrieval operation binding the contract event 0xcb042d04619e497175f63ce8ec9b9127cb0bf344facc37db4fe8492c2ecfbb9e.
//
// Solidity: event passportApplied(int64 applyerTg, address wallet_address)
func (_Passport *PassportFilterer) FilterPassportApplied(opts *bind.FilterOpts) (*PassportPassportAppliedIterator, error) {

	logs, sub, err := _Passport.contract.FilterLogs(opts, "passportApplied")
	if err != nil {
		return nil, err
	}
	return &PassportPassportAppliedIterator{contract: _Passport.contract, event: "passportApplied", logs: logs, sub: sub}, nil
}

// WatchPassportApplied is a free log subscription operation binding the contract event 0xcb042d04619e497175f63ce8ec9b9127cb0bf344facc37db4fe8492c2ecfbb9e.
//
// Solidity: event passportApplied(int64 applyerTg, address wallet_address)
func (_Passport *PassportFilterer) WatchPassportApplied(opts *bind.WatchOpts, sink chan<- *PassportPassportApplied) (event.Subscription, error) {

	logs, sub, err := _Passport.contract.WatchLogs(opts, "passportApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportPassportApplied)
				if err := _Passport.contract.UnpackLog(event, "passportApplied", log); err != nil {
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

// ParsePassportApplied is a log parse operation binding the contract event 0xcb042d04619e497175f63ce8ec9b9127cb0bf344facc37db4fe8492c2ecfbb9e.
//
// Solidity: event passportApplied(int64 applyerTg, address wallet_address)
func (_Passport *PassportFilterer) ParsePassportApplied(log types.Log) (*PassportPassportApplied, error) {
	event := new(PassportPassportApplied)
	if err := _Passport.contract.UnpackLog(event, "passportApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportPassportAppliedIndexedIterator is returned from FilterPassportAppliedIndexed and is used to iterate over the raw logs and unpacked data for PassportAppliedIndexed events raised by the Passport contract.
type PassportPassportAppliedIndexedIterator struct {
	Event *PassportPassportAppliedIndexed // Event containing the contract specifics and raw log

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
func (it *PassportPassportAppliedIndexedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportPassportAppliedIndexed)
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
		it.Event = new(PassportPassportAppliedIndexed)
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
func (it *PassportPassportAppliedIndexedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportPassportAppliedIndexedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportPassportAppliedIndexed represents a PassportAppliedIndexed event raised by the Passport contract.
type PassportPassportAppliedIndexed struct {
	ApplyerTg     int64
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPassportAppliedIndexed is a free log retrieval operation binding the contract event 0x2d329f0d038ded94203a58e5be2f2f41b14cbbc49cc1f89443696606f104658c.
//
// Solidity: event passportAppliedIndexed(int64 indexed applyerTg, address wallet_address)
func (_Passport *PassportFilterer) FilterPassportAppliedIndexed(opts *bind.FilterOpts, applyerTg []int64) (*PassportPassportAppliedIndexedIterator, error) {

	var applyerTgRule []interface{}
	for _, applyerTgItem := range applyerTg {
		applyerTgRule = append(applyerTgRule, applyerTgItem)
	}

	logs, sub, err := _Passport.contract.FilterLogs(opts, "passportAppliedIndexed", applyerTgRule)
	if err != nil {
		return nil, err
	}
	return &PassportPassportAppliedIndexedIterator{contract: _Passport.contract, event: "passportAppliedIndexed", logs: logs, sub: sub}, nil
}

// WatchPassportAppliedIndexed is a free log subscription operation binding the contract event 0x2d329f0d038ded94203a58e5be2f2f41b14cbbc49cc1f89443696606f104658c.
//
// Solidity: event passportAppliedIndexed(int64 indexed applyerTg, address wallet_address)
func (_Passport *PassportFilterer) WatchPassportAppliedIndexed(opts *bind.WatchOpts, sink chan<- *PassportPassportAppliedIndexed, applyerTg []int64) (event.Subscription, error) {

	var applyerTgRule []interface{}
	for _, applyerTgItem := range applyerTg {
		applyerTgRule = append(applyerTgRule, applyerTgItem)
	}

	logs, sub, err := _Passport.contract.WatchLogs(opts, "passportAppliedIndexed", applyerTgRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportPassportAppliedIndexed)
				if err := _Passport.contract.UnpackLog(event, "passportAppliedIndexed", log); err != nil {
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

// ParsePassportAppliedIndexed is a log parse operation binding the contract event 0x2d329f0d038ded94203a58e5be2f2f41b14cbbc49cc1f89443696606f104658c.
//
// Solidity: event passportAppliedIndexed(int64 indexed applyerTg, address wallet_address)
func (_Passport *PassportFilterer) ParsePassportAppliedIndexed(log types.Log) (*PassportPassportAppliedIndexed, error) {
	event := new(PassportPassportAppliedIndexed)
	if err := _Passport.contract.UnpackLog(event, "passportAppliedIndexed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportPassportApprovedIterator is returned from FilterPassportApproved and is used to iterate over the raw logs and unpacked data for PassportApproved events raised by the Passport contract.
type PassportPassportApprovedIterator struct {
	Event *PassportPassportApproved // Event containing the contract specifics and raw log

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
func (it *PassportPassportApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportPassportApproved)
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
		it.Event = new(PassportPassportApproved)
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
func (it *PassportPassportApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportPassportApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportPassportApproved represents a PassportApproved event raised by the Passport contract.
type PassportPassportApproved struct {
	ApplyerTg     int64
	WalletAddress common.Address
	Issuer        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPassportApproved is a free log retrieval operation binding the contract event 0xddda54aebd53f1c2d3f35d7107ddd15f71b313a804dda2179cded44927110cd0.
//
// Solidity: event passportApproved(int64 applyerTg, address wallet_address, address issuer)
func (_Passport *PassportFilterer) FilterPassportApproved(opts *bind.FilterOpts) (*PassportPassportApprovedIterator, error) {

	logs, sub, err := _Passport.contract.FilterLogs(opts, "passportApproved")
	if err != nil {
		return nil, err
	}
	return &PassportPassportApprovedIterator{contract: _Passport.contract, event: "passportApproved", logs: logs, sub: sub}, nil
}

// WatchPassportApproved is a free log subscription operation binding the contract event 0xddda54aebd53f1c2d3f35d7107ddd15f71b313a804dda2179cded44927110cd0.
//
// Solidity: event passportApproved(int64 applyerTg, address wallet_address, address issuer)
func (_Passport *PassportFilterer) WatchPassportApproved(opts *bind.WatchOpts, sink chan<- *PassportPassportApproved) (event.Subscription, error) {

	logs, sub, err := _Passport.contract.WatchLogs(opts, "passportApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportPassportApproved)
				if err := _Passport.contract.UnpackLog(event, "passportApproved", log); err != nil {
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

// ParsePassportApproved is a log parse operation binding the contract event 0xddda54aebd53f1c2d3f35d7107ddd15f71b313a804dda2179cded44927110cd0.
//
// Solidity: event passportApproved(int64 applyerTg, address wallet_address, address issuer)
func (_Passport *PassportFilterer) ParsePassportApproved(log types.Log) (*PassportPassportApproved, error) {
	event := new(PassportPassportApproved)
	if err := _Passport.contract.UnpackLog(event, "passportApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportPassportDeniedIterator is returned from FilterPassportDenied and is used to iterate over the raw logs and unpacked data for PassportDenied events raised by the Passport contract.
type PassportPassportDeniedIterator struct {
	Event *PassportPassportDenied // Event containing the contract specifics and raw log

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
func (it *PassportPassportDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportPassportDenied)
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
		it.Event = new(PassportPassportDenied)
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
func (it *PassportPassportDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportPassportDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportPassportDenied represents a PassportDenied event raised by the Passport contract.
type PassportPassportDenied struct {
	ApplyerTg int64
	Wallet    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPassportDenied is a free log retrieval operation binding the contract event 0x0511bec3ed97c7099362d7e7b0ed4fe8afe6f448a0dcf73f258be48dceb302db.
//
// Solidity: event passportDenied(int64 applyerTg, address wallet)
func (_Passport *PassportFilterer) FilterPassportDenied(opts *bind.FilterOpts) (*PassportPassportDeniedIterator, error) {

	logs, sub, err := _Passport.contract.FilterLogs(opts, "passportDenied")
	if err != nil {
		return nil, err
	}
	return &PassportPassportDeniedIterator{contract: _Passport.contract, event: "passportDenied", logs: logs, sub: sub}, nil
}

// WatchPassportDenied is a free log subscription operation binding the contract event 0x0511bec3ed97c7099362d7e7b0ed4fe8afe6f448a0dcf73f258be48dceb302db.
//
// Solidity: event passportDenied(int64 applyerTg, address wallet)
func (_Passport *PassportFilterer) WatchPassportDenied(opts *bind.WatchOpts, sink chan<- *PassportPassportDenied) (event.Subscription, error) {

	logs, sub, err := _Passport.contract.WatchLogs(opts, "passportDenied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportPassportDenied)
				if err := _Passport.contract.UnpackLog(event, "passportDenied", log); err != nil {
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

// ParsePassportDenied is a log parse operation binding the contract event 0x0511bec3ed97c7099362d7e7b0ed4fe8afe6f448a0dcf73f258be48dceb302db.
//
// Solidity: event passportDenied(int64 applyerTg, address wallet)
func (_Passport *PassportFilterer) ParsePassportDenied(log types.Log) (*PassportPassportDenied, error) {
	event := new(PassportPassportDenied)
	if err := _Passport.contract.UnpackLog(event, "passportDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportTrustedIndexedIterator is returned from FilterTrustedIndexed and is used to iterate over the raw logs and unpacked data for TrustedIndexed events raised by the Passport contract.
type PassportTrustedIndexedIterator struct {
	Event *PassportTrustedIndexed // Event containing the contract specifics and raw log

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
func (it *PassportTrustedIndexedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportTrustedIndexed)
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
		it.Event = new(PassportTrustedIndexed)
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
func (it *PassportTrustedIndexedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportTrustedIndexedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportTrustedIndexed represents a TrustedIndexed event raised by the Passport contract.
type PassportTrustedIndexed struct {
	From int64
	To   int64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedIndexed is a free log retrieval operation binding the contract event 0xc6910544ad28778f5982555eae73efab16f145f0e382250ac13db569707ca1cf.
//
// Solidity: event trustedIndexed(int64 from, int64 indexed to)
func (_Passport *PassportFilterer) FilterTrustedIndexed(opts *bind.FilterOpts, to []int64) (*PassportTrustedIndexedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Passport.contract.FilterLogs(opts, "trustedIndexed", toRule)
	if err != nil {
		return nil, err
	}
	return &PassportTrustedIndexedIterator{contract: _Passport.contract, event: "trustedIndexed", logs: logs, sub: sub}, nil
}

// WatchTrustedIndexed is a free log subscription operation binding the contract event 0xc6910544ad28778f5982555eae73efab16f145f0e382250ac13db569707ca1cf.
//
// Solidity: event trustedIndexed(int64 from, int64 indexed to)
func (_Passport *PassportFilterer) WatchTrustedIndexed(opts *bind.WatchOpts, sink chan<- *PassportTrustedIndexed, to []int64) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Passport.contract.WatchLogs(opts, "trustedIndexed", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportTrustedIndexed)
				if err := _Passport.contract.UnpackLog(event, "trustedIndexed", log); err != nil {
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

// ParseTrustedIndexed is a log parse operation binding the contract event 0xc6910544ad28778f5982555eae73efab16f145f0e382250ac13db569707ca1cf.
//
// Solidity: event trustedIndexed(int64 from, int64 indexed to)
func (_Passport *PassportFilterer) ParseTrustedIndexed(log types.Log) (*PassportTrustedIndexed, error) {
	event := new(PassportTrustedIndexed)
	if err := _Passport.contract.UnpackLog(event, "trustedIndexed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportUntrustedIndexedIterator is returned from FilterUntrustedIndexed and is used to iterate over the raw logs and unpacked data for UntrustedIndexed events raised by the Passport contract.
type PassportUntrustedIndexedIterator struct {
	Event *PassportUntrustedIndexed // Event containing the contract specifics and raw log

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
func (it *PassportUntrustedIndexedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportUntrustedIndexed)
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
		it.Event = new(PassportUntrustedIndexed)
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
func (it *PassportUntrustedIndexedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportUntrustedIndexedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportUntrustedIndexed represents a UntrustedIndexed event raised by the Passport contract.
type PassportUntrustedIndexed struct {
	From int64
	To   int64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUntrustedIndexed is a free log retrieval operation binding the contract event 0x72ea88320a14be91b677bb10d77500a955301d37422909c3ad479fcf35e10093.
//
// Solidity: event untrustedIndexed(int64 from, int64 indexed to)
func (_Passport *PassportFilterer) FilterUntrustedIndexed(opts *bind.FilterOpts, to []int64) (*PassportUntrustedIndexedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Passport.contract.FilterLogs(opts, "untrustedIndexed", toRule)
	if err != nil {
		return nil, err
	}
	return &PassportUntrustedIndexedIterator{contract: _Passport.contract, event: "untrustedIndexed", logs: logs, sub: sub}, nil
}

// WatchUntrustedIndexed is a free log subscription operation binding the contract event 0x72ea88320a14be91b677bb10d77500a955301d37422909c3ad479fcf35e10093.
//
// Solidity: event untrustedIndexed(int64 from, int64 indexed to)
func (_Passport *PassportFilterer) WatchUntrustedIndexed(opts *bind.WatchOpts, sink chan<- *PassportUntrustedIndexed, to []int64) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Passport.contract.WatchLogs(opts, "untrustedIndexed", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportUntrustedIndexed)
				if err := _Passport.contract.UnpackLog(event, "untrustedIndexed", log); err != nil {
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

// ParseUntrustedIndexed is a log parse operation binding the contract event 0x72ea88320a14be91b677bb10d77500a955301d37422909c3ad479fcf35e10093.
//
// Solidity: event untrustedIndexed(int64 from, int64 indexed to)
func (_Passport *PassportFilterer) ParseUntrustedIndexed(log types.Log) (*PassportUntrustedIndexed, error) {
	event := new(PassportUntrustedIndexed)
	if err := _Passport.contract.UnpackLog(event, "untrustedIndexed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
