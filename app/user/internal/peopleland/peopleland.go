// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package peopleland

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

// IPeopleLandLand is an auto generated low-level Go binding around an user-defined struct.
type IPeopleLandLand struct {
	X             *big.Int
	Y             *big.Int
	Slogan        string
	MintedAddress common.Address
	GivedAddress  common.Address
	IsMinted      bool
	IsGived       bool
}

// ITokenSVGTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type ITokenSVGTokenInfo struct {
	X          *big.Int
	Y          *big.Int
	TokenId    *big.Int
	HasTokenId bool
}

// PeoplelandMetaData contains all meta data concerning the Peopleland contract.
var PeoplelandMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_startUp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenSVG\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"}],\"name\":\"GiveTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintedAddress\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"}],\"name\":\"SetSlogan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGN_MESSAGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCoordinates\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"}],\"name\":\"getCoordinatesStrings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"sx\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sy\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getInviteParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_ip\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_ib\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasTokenId\",\"type\":\"bool\"}],\"internalType\":\"structITokenSVG.TokenInfo\",\"name\":\"_invite\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mintedAddress\",\"type\":\"address\"}],\"name\":\"getMintLands\",\"outputs\":[{\"components\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mintedAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMinted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isGived\",\"type\":\"bool\"}],\"internalType\":\"structIPeopleLand.Land[]\",\"name\":\"_mintLands\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMintedAndInvitedList\",\"outputs\":[{\"components\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasTokenId\",\"type\":\"bool\"}],\"internalType\":\"structITokenSVG.TokenInfo[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"}],\"name\":\"getNeighborsParams\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"tokenIds\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"}],\"name\":\"giveTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_givedAddress\",\"type\":\"address\"}],\"name\":\"givedLand\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGived\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mintedAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMinted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isGived\",\"type\":\"bool\"}],\"internalType\":\"structIPeopleLand.Land\",\"name\":\"_land\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBuilder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPeople\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"_x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_y\",\"type\":\"int128\"}],\"name\":\"land\",\"outputs\":[{\"components\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mintedAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMinted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isGived\",\"type\":\"bool\"}],\"internalType\":\"structIPeopleLand.Land\",\"name\":\"_land\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"x2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y2\",\"type\":\"int128\"}],\"name\":\"mint2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y1\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"givedAddress1\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"x2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y2\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"givedAddress2\",\"type\":\"address\"}],\"name\":\"mint2AndGiveTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"}],\"name\":\"mintAndGiveTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"}],\"name\":\"mintAndGiveToWithSlogan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintLandCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintSelfSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"}],\"name\":\"mintToBuilderByOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"givedAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"}],\"name\":\"mintToBuilderByOwnerWithSlogan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"mintToSelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openMintSelfSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"}],\"name\":\"packedXY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_packedXY\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"}],\"name\":\"setSlogan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_attr\",\"type\":\"address\"}],\"name\":\"setTokenSVGAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenSVGAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PeoplelandABI is the input ABI used to generate the binding from.
// Deprecated: Use PeoplelandMetaData.ABI instead.
var PeoplelandABI = PeoplelandMetaData.ABI

// Peopleland is an auto generated Go binding around an Ethereum contract.
type Peopleland struct {
	PeoplelandCaller     // Read-only binding to the contract
	PeoplelandTransactor // Write-only binding to the contract
	PeoplelandFilterer   // Log filterer for contract events
}

// PeoplelandCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeoplelandCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeoplelandTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeoplelandTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeoplelandFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeoplelandFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeoplelandSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeoplelandSession struct {
	Contract     *Peopleland       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeoplelandCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeoplelandCallerSession struct {
	Contract *PeoplelandCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PeoplelandTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeoplelandTransactorSession struct {
	Contract     *PeoplelandTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PeoplelandRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeoplelandRaw struct {
	Contract *Peopleland // Generic contract binding to access the raw methods on
}

// PeoplelandCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeoplelandCallerRaw struct {
	Contract *PeoplelandCaller // Generic read-only contract binding to access the raw methods on
}

// PeoplelandTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeoplelandTransactorRaw struct {
	Contract *PeoplelandTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeopleland creates a new instance of Peopleland, bound to a specific deployed contract.
func NewPeopleland(address common.Address, backend bind.ContractBackend) (*Peopleland, error) {
	contract, err := bindPeopleland(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Peopleland{PeoplelandCaller: PeoplelandCaller{contract: contract}, PeoplelandTransactor: PeoplelandTransactor{contract: contract}, PeoplelandFilterer: PeoplelandFilterer{contract: contract}}, nil
}

// NewPeoplelandCaller creates a new read-only instance of Peopleland, bound to a specific deployed contract.
func NewPeoplelandCaller(address common.Address, caller bind.ContractCaller) (*PeoplelandCaller, error) {
	contract, err := bindPeopleland(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeoplelandCaller{contract: contract}, nil
}

// NewPeoplelandTransactor creates a new write-only instance of Peopleland, bound to a specific deployed contract.
func NewPeoplelandTransactor(address common.Address, transactor bind.ContractTransactor) (*PeoplelandTransactor, error) {
	contract, err := bindPeopleland(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeoplelandTransactor{contract: contract}, nil
}

// NewPeoplelandFilterer creates a new log filterer instance of Peopleland, bound to a specific deployed contract.
func NewPeoplelandFilterer(address common.Address, filterer bind.ContractFilterer) (*PeoplelandFilterer, error) {
	contract, err := bindPeopleland(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeoplelandFilterer{contract: contract}, nil
}

// bindPeopleland binds a generic wrapper to an already deployed contract.
func bindPeopleland(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeoplelandABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peopleland *PeoplelandRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peopleland.Contract.PeoplelandCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peopleland *PeoplelandRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peopleland.Contract.PeoplelandTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peopleland *PeoplelandRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peopleland.Contract.PeoplelandTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peopleland *PeoplelandCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peopleland.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peopleland *PeoplelandTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peopleland.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peopleland *PeoplelandTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peopleland.Contract.contract.Transact(opts, method, params...)
}

// PRICE is a free data retrieval call binding the contract method 0x8d859f3e.
//
// Solidity: function PRICE() view returns(uint256)
func (_Peopleland *PeoplelandCaller) PRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICE is a free data retrieval call binding the contract method 0x8d859f3e.
//
// Solidity: function PRICE() view returns(uint256)
func (_Peopleland *PeoplelandSession) PRICE() (*big.Int, error) {
	return _Peopleland.Contract.PRICE(&_Peopleland.CallOpts)
}

// PRICE is a free data retrieval call binding the contract method 0x8d859f3e.
//
// Solidity: function PRICE() view returns(uint256)
func (_Peopleland *PeoplelandCallerSession) PRICE() (*big.Int, error) {
	return _Peopleland.Contract.PRICE(&_Peopleland.CallOpts)
}

// SIGNMESSAGEADDRESS is a free data retrieval call binding the contract method 0xde4bc39d.
//
// Solidity: function SIGN_MESSAGE_ADDRESS() view returns(address)
func (_Peopleland *PeoplelandCaller) SIGNMESSAGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "SIGN_MESSAGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SIGNMESSAGEADDRESS is a free data retrieval call binding the contract method 0xde4bc39d.
//
// Solidity: function SIGN_MESSAGE_ADDRESS() view returns(address)
func (_Peopleland *PeoplelandSession) SIGNMESSAGEADDRESS() (common.Address, error) {
	return _Peopleland.Contract.SIGNMESSAGEADDRESS(&_Peopleland.CallOpts)
}

// SIGNMESSAGEADDRESS is a free data retrieval call binding the contract method 0xde4bc39d.
//
// Solidity: function SIGN_MESSAGE_ADDRESS() view returns(address)
func (_Peopleland *PeoplelandCallerSession) SIGNMESSAGEADDRESS() (common.Address, error) {
	return _Peopleland.Contract.SIGNMESSAGEADDRESS(&_Peopleland.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Peopleland *PeoplelandCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Peopleland *PeoplelandSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Peopleland.Contract.BalanceOf(&_Peopleland.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Peopleland *PeoplelandCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Peopleland.Contract.BalanceOf(&_Peopleland.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Peopleland *PeoplelandCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Peopleland *PeoplelandSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Peopleland.Contract.GetApproved(&_Peopleland.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Peopleland *PeoplelandCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Peopleland.Contract.GetApproved(&_Peopleland.CallOpts, tokenId)
}

// GetCoordinates is a free data retrieval call binding the contract method 0x0bf37818.
//
// Solidity: function getCoordinates(uint256 tokenId) view returns(int128 x, int128 y)
func (_Peopleland *PeoplelandCaller) GetCoordinates(opts *bind.CallOpts, tokenId *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getCoordinates", tokenId)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCoordinates is a free data retrieval call binding the contract method 0x0bf37818.
//
// Solidity: function getCoordinates(uint256 tokenId) view returns(int128 x, int128 y)
func (_Peopleland *PeoplelandSession) GetCoordinates(tokenId *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Peopleland.Contract.GetCoordinates(&_Peopleland.CallOpts, tokenId)
}

// GetCoordinates is a free data retrieval call binding the contract method 0x0bf37818.
//
// Solidity: function getCoordinates(uint256 tokenId) view returns(int128 x, int128 y)
func (_Peopleland *PeoplelandCallerSession) GetCoordinates(tokenId *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Peopleland.Contract.GetCoordinates(&_Peopleland.CallOpts, tokenId)
}

// GetCoordinatesStrings is a free data retrieval call binding the contract method 0x8fe22784.
//
// Solidity: function getCoordinatesStrings(int128 x, int128 y) view returns(string sx, string sy)
func (_Peopleland *PeoplelandCaller) GetCoordinatesStrings(opts *bind.CallOpts, x *big.Int, y *big.Int) (struct {
	Sx string
	Sy string
}, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getCoordinatesStrings", x, y)

	outstruct := new(struct {
		Sx string
		Sy string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sx = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Sy = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetCoordinatesStrings is a free data retrieval call binding the contract method 0x8fe22784.
//
// Solidity: function getCoordinatesStrings(int128 x, int128 y) view returns(string sx, string sy)
func (_Peopleland *PeoplelandSession) GetCoordinatesStrings(x *big.Int, y *big.Int) (struct {
	Sx string
	Sy string
}, error) {
	return _Peopleland.Contract.GetCoordinatesStrings(&_Peopleland.CallOpts, x, y)
}

// GetCoordinatesStrings is a free data retrieval call binding the contract method 0x8fe22784.
//
// Solidity: function getCoordinatesStrings(int128 x, int128 y) view returns(string sx, string sy)
func (_Peopleland *PeoplelandCallerSession) GetCoordinatesStrings(x *big.Int, y *big.Int) (struct {
	Sx string
	Sy string
}, error) {
	return _Peopleland.Contract.GetCoordinatesStrings(&_Peopleland.CallOpts, x, y)
}

// GetInviteParams is a free data retrieval call binding the contract method 0x86c52ac0.
//
// Solidity: function getInviteParams(uint256 tokenId) view returns(bool _ip, bool _ib, (int128,int128,uint256,bool) _invite)
func (_Peopleland *PeoplelandCaller) GetInviteParams(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Ip     bool
	Ib     bool
	Invite ITokenSVGTokenInfo
}, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getInviteParams", tokenId)

	outstruct := new(struct {
		Ip     bool
		Ib     bool
		Invite ITokenSVGTokenInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ip = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Ib = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Invite = *abi.ConvertType(out[2], new(ITokenSVGTokenInfo)).(*ITokenSVGTokenInfo)

	return *outstruct, err

}

// GetInviteParams is a free data retrieval call binding the contract method 0x86c52ac0.
//
// Solidity: function getInviteParams(uint256 tokenId) view returns(bool _ip, bool _ib, (int128,int128,uint256,bool) _invite)
func (_Peopleland *PeoplelandSession) GetInviteParams(tokenId *big.Int) (struct {
	Ip     bool
	Ib     bool
	Invite ITokenSVGTokenInfo
}, error) {
	return _Peopleland.Contract.GetInviteParams(&_Peopleland.CallOpts, tokenId)
}

// GetInviteParams is a free data retrieval call binding the contract method 0x86c52ac0.
//
// Solidity: function getInviteParams(uint256 tokenId) view returns(bool _ip, bool _ib, (int128,int128,uint256,bool) _invite)
func (_Peopleland *PeoplelandCallerSession) GetInviteParams(tokenId *big.Int) (struct {
	Ip     bool
	Ib     bool
	Invite ITokenSVGTokenInfo
}, error) {
	return _Peopleland.Contract.GetInviteParams(&_Peopleland.CallOpts, tokenId)
}

// GetMintLands is a free data retrieval call binding the contract method 0x3413f6e6.
//
// Solidity: function getMintLands(address _mintedAddress) view returns((int128,int128,string,address,address,bool,bool)[] _mintLands)
func (_Peopleland *PeoplelandCaller) GetMintLands(opts *bind.CallOpts, _mintedAddress common.Address) ([]IPeopleLandLand, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getMintLands", _mintedAddress)

	if err != nil {
		return *new([]IPeopleLandLand), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPeopleLandLand)).(*[]IPeopleLandLand)

	return out0, err

}

// GetMintLands is a free data retrieval call binding the contract method 0x3413f6e6.
//
// Solidity: function getMintLands(address _mintedAddress) view returns((int128,int128,string,address,address,bool,bool)[] _mintLands)
func (_Peopleland *PeoplelandSession) GetMintLands(_mintedAddress common.Address) ([]IPeopleLandLand, error) {
	return _Peopleland.Contract.GetMintLands(&_Peopleland.CallOpts, _mintedAddress)
}

// GetMintLands is a free data retrieval call binding the contract method 0x3413f6e6.
//
// Solidity: function getMintLands(address _mintedAddress) view returns((int128,int128,string,address,address,bool,bool)[] _mintLands)
func (_Peopleland *PeoplelandCallerSession) GetMintLands(_mintedAddress common.Address) ([]IPeopleLandLand, error) {
	return _Peopleland.Contract.GetMintLands(&_Peopleland.CallOpts, _mintedAddress)
}

// GetMintedAndInvitedList is a free data retrieval call binding the contract method 0xa5e9f6fb.
//
// Solidity: function getMintedAndInvitedList(uint256 tokenId) view returns((int128,int128,uint256,bool)[] _list)
func (_Peopleland *PeoplelandCaller) GetMintedAndInvitedList(opts *bind.CallOpts, tokenId *big.Int) ([]ITokenSVGTokenInfo, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getMintedAndInvitedList", tokenId)

	if err != nil {
		return *new([]ITokenSVGTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenSVGTokenInfo)).(*[]ITokenSVGTokenInfo)

	return out0, err

}

// GetMintedAndInvitedList is a free data retrieval call binding the contract method 0xa5e9f6fb.
//
// Solidity: function getMintedAndInvitedList(uint256 tokenId) view returns((int128,int128,uint256,bool)[] _list)
func (_Peopleland *PeoplelandSession) GetMintedAndInvitedList(tokenId *big.Int) ([]ITokenSVGTokenInfo, error) {
	return _Peopleland.Contract.GetMintedAndInvitedList(&_Peopleland.CallOpts, tokenId)
}

// GetMintedAndInvitedList is a free data retrieval call binding the contract method 0xa5e9f6fb.
//
// Solidity: function getMintedAndInvitedList(uint256 tokenId) view returns((int128,int128,uint256,bool)[] _list)
func (_Peopleland *PeoplelandCallerSession) GetMintedAndInvitedList(tokenId *big.Int) ([]ITokenSVGTokenInfo, error) {
	return _Peopleland.Contract.GetMintedAndInvitedList(&_Peopleland.CallOpts, tokenId)
}

// GetNeighborsParams is a free data retrieval call binding the contract method 0x3dfe96de.
//
// Solidity: function getNeighborsParams(int128 x, int128 y) view returns(string[] tokenIds)
func (_Peopleland *PeoplelandCaller) GetNeighborsParams(opts *bind.CallOpts, x *big.Int, y *big.Int) ([]string, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getNeighborsParams", x, y)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetNeighborsParams is a free data retrieval call binding the contract method 0x3dfe96de.
//
// Solidity: function getNeighborsParams(int128 x, int128 y) view returns(string[] tokenIds)
func (_Peopleland *PeoplelandSession) GetNeighborsParams(x *big.Int, y *big.Int) ([]string, error) {
	return _Peopleland.Contract.GetNeighborsParams(&_Peopleland.CallOpts, x, y)
}

// GetNeighborsParams is a free data retrieval call binding the contract method 0x3dfe96de.
//
// Solidity: function getNeighborsParams(int128 x, int128 y) view returns(string[] tokenIds)
func (_Peopleland *PeoplelandCallerSession) GetNeighborsParams(x *big.Int, y *big.Int) ([]string, error) {
	return _Peopleland.Contract.GetNeighborsParams(&_Peopleland.CallOpts, x, y)
}

// GetTokenId is a free data retrieval call binding the contract method 0x7e5a6be8.
//
// Solidity: function getTokenId(int128 x, int128 y) view returns(uint256 tokenId)
func (_Peopleland *PeoplelandCaller) GetTokenId(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "getTokenId", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenId is a free data retrieval call binding the contract method 0x7e5a6be8.
//
// Solidity: function getTokenId(int128 x, int128 y) view returns(uint256 tokenId)
func (_Peopleland *PeoplelandSession) GetTokenId(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.GetTokenId(&_Peopleland.CallOpts, x, y)
}

// GetTokenId is a free data retrieval call binding the contract method 0x7e5a6be8.
//
// Solidity: function getTokenId(int128 x, int128 y) view returns(uint256 tokenId)
func (_Peopleland *PeoplelandCallerSession) GetTokenId(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.GetTokenId(&_Peopleland.CallOpts, x, y)
}

// GivedLand is a free data retrieval call binding the contract method 0xbe29010d.
//
// Solidity: function givedLand(address _givedAddress) view returns(bool isGived, (int128,int128,string,address,address,bool,bool) _land)
func (_Peopleland *PeoplelandCaller) GivedLand(opts *bind.CallOpts, _givedAddress common.Address) (struct {
	IsGived bool
	Land    IPeopleLandLand
}, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "givedLand", _givedAddress)

	outstruct := new(struct {
		IsGived bool
		Land    IPeopleLandLand
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsGived = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Land = *abi.ConvertType(out[1], new(IPeopleLandLand)).(*IPeopleLandLand)

	return *outstruct, err

}

// GivedLand is a free data retrieval call binding the contract method 0xbe29010d.
//
// Solidity: function givedLand(address _givedAddress) view returns(bool isGived, (int128,int128,string,address,address,bool,bool) _land)
func (_Peopleland *PeoplelandSession) GivedLand(_givedAddress common.Address) (struct {
	IsGived bool
	Land    IPeopleLandLand
}, error) {
	return _Peopleland.Contract.GivedLand(&_Peopleland.CallOpts, _givedAddress)
}

// GivedLand is a free data retrieval call binding the contract method 0xbe29010d.
//
// Solidity: function givedLand(address _givedAddress) view returns(bool isGived, (int128,int128,string,address,address,bool,bool) _land)
func (_Peopleland *PeoplelandCallerSession) GivedLand(_givedAddress common.Address) (struct {
	IsGived bool
	Land    IPeopleLandLand
}, error) {
	return _Peopleland.Contract.GivedLand(&_Peopleland.CallOpts, _givedAddress)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Peopleland *PeoplelandCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Peopleland *PeoplelandSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Peopleland.Contract.IsApprovedForAll(&_Peopleland.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Peopleland *PeoplelandCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Peopleland.Contract.IsApprovedForAll(&_Peopleland.CallOpts, owner, operator)
}

// IsBuilder is a free data retrieval call binding the contract method 0xb6b6b475.
//
// Solidity: function isBuilder(address ) view returns(bool)
func (_Peopleland *PeoplelandCaller) IsBuilder(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "isBuilder", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuilder is a free data retrieval call binding the contract method 0xb6b6b475.
//
// Solidity: function isBuilder(address ) view returns(bool)
func (_Peopleland *PeoplelandSession) IsBuilder(arg0 common.Address) (bool, error) {
	return _Peopleland.Contract.IsBuilder(&_Peopleland.CallOpts, arg0)
}

// IsBuilder is a free data retrieval call binding the contract method 0xb6b6b475.
//
// Solidity: function isBuilder(address ) view returns(bool)
func (_Peopleland *PeoplelandCallerSession) IsBuilder(arg0 common.Address) (bool, error) {
	return _Peopleland.Contract.IsBuilder(&_Peopleland.CallOpts, arg0)
}

// IsPeople is a free data retrieval call binding the contract method 0xc9d2f633.
//
// Solidity: function isPeople(address ) view returns(bool)
func (_Peopleland *PeoplelandCaller) IsPeople(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "isPeople", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeople is a free data retrieval call binding the contract method 0xc9d2f633.
//
// Solidity: function isPeople(address ) view returns(bool)
func (_Peopleland *PeoplelandSession) IsPeople(arg0 common.Address) (bool, error) {
	return _Peopleland.Contract.IsPeople(&_Peopleland.CallOpts, arg0)
}

// IsPeople is a free data retrieval call binding the contract method 0xc9d2f633.
//
// Solidity: function isPeople(address ) view returns(bool)
func (_Peopleland *PeoplelandCallerSession) IsPeople(arg0 common.Address) (bool, error) {
	return _Peopleland.Contract.IsPeople(&_Peopleland.CallOpts, arg0)
}

// Land is a free data retrieval call binding the contract method 0x8b9c06b8.
//
// Solidity: function land(int128 _x, int128 _y) view returns((int128,int128,string,address,address,bool,bool) _land)
func (_Peopleland *PeoplelandCaller) Land(opts *bind.CallOpts, _x *big.Int, _y *big.Int) (IPeopleLandLand, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "land", _x, _y)

	if err != nil {
		return *new(IPeopleLandLand), err
	}

	out0 := *abi.ConvertType(out[0], new(IPeopleLandLand)).(*IPeopleLandLand)

	return out0, err

}

// Land is a free data retrieval call binding the contract method 0x8b9c06b8.
//
// Solidity: function land(int128 _x, int128 _y) view returns((int128,int128,string,address,address,bool,bool) _land)
func (_Peopleland *PeoplelandSession) Land(_x *big.Int, _y *big.Int) (IPeopleLandLand, error) {
	return _Peopleland.Contract.Land(&_Peopleland.CallOpts, _x, _y)
}

// Land is a free data retrieval call binding the contract method 0x8b9c06b8.
//
// Solidity: function land(int128 _x, int128 _y) view returns((int128,int128,string,address,address,bool,bool) _land)
func (_Peopleland *PeoplelandCallerSession) Land(_x *big.Int, _y *big.Int) (IPeopleLandLand, error) {
	return _Peopleland.Contract.Land(&_Peopleland.CallOpts, _x, _y)
}

// MintLandCount is a free data retrieval call binding the contract method 0xd9206aa3.
//
// Solidity: function mintLandCount(address ) view returns(uint8)
func (_Peopleland *PeoplelandCaller) MintLandCount(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "mintLandCount", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MintLandCount is a free data retrieval call binding the contract method 0xd9206aa3.
//
// Solidity: function mintLandCount(address ) view returns(uint8)
func (_Peopleland *PeoplelandSession) MintLandCount(arg0 common.Address) (uint8, error) {
	return _Peopleland.Contract.MintLandCount(&_Peopleland.CallOpts, arg0)
}

// MintLandCount is a free data retrieval call binding the contract method 0xd9206aa3.
//
// Solidity: function mintLandCount(address ) view returns(uint8)
func (_Peopleland *PeoplelandCallerSession) MintLandCount(arg0 common.Address) (uint8, error) {
	return _Peopleland.Contract.MintLandCount(&_Peopleland.CallOpts, arg0)
}

// MintSelfSwitch is a free data retrieval call binding the contract method 0xe503fc0b.
//
// Solidity: function mintSelfSwitch() view returns(bool)
func (_Peopleland *PeoplelandCaller) MintSelfSwitch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "mintSelfSwitch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintSelfSwitch is a free data retrieval call binding the contract method 0xe503fc0b.
//
// Solidity: function mintSelfSwitch() view returns(bool)
func (_Peopleland *PeoplelandSession) MintSelfSwitch() (bool, error) {
	return _Peopleland.Contract.MintSelfSwitch(&_Peopleland.CallOpts)
}

// MintSelfSwitch is a free data retrieval call binding the contract method 0xe503fc0b.
//
// Solidity: function mintSelfSwitch() view returns(bool)
func (_Peopleland *PeoplelandCallerSession) MintSelfSwitch() (bool, error) {
	return _Peopleland.Contract.MintSelfSwitch(&_Peopleland.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peopleland *PeoplelandCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peopleland *PeoplelandSession) Name() (string, error) {
	return _Peopleland.Contract.Name(&_Peopleland.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peopleland *PeoplelandCallerSession) Name() (string, error) {
	return _Peopleland.Contract.Name(&_Peopleland.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Peopleland *PeoplelandCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Peopleland *PeoplelandSession) Owner() (common.Address, error) {
	return _Peopleland.Contract.Owner(&_Peopleland.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Peopleland *PeoplelandCallerSession) Owner() (common.Address, error) {
	return _Peopleland.Contract.Owner(&_Peopleland.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Peopleland *PeoplelandCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Peopleland *PeoplelandSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Peopleland.Contract.OwnerOf(&_Peopleland.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Peopleland *PeoplelandCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Peopleland.Contract.OwnerOf(&_Peopleland.CallOpts, tokenId)
}

// PackedXY is a free data retrieval call binding the contract method 0xa8696223.
//
// Solidity: function packedXY(int128 x, int128 y) pure returns(uint256 _packedXY)
func (_Peopleland *PeoplelandCaller) PackedXY(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "packedXY", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackedXY is a free data retrieval call binding the contract method 0xa8696223.
//
// Solidity: function packedXY(int128 x, int128 y) pure returns(uint256 _packedXY)
func (_Peopleland *PeoplelandSession) PackedXY(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.PackedXY(&_Peopleland.CallOpts, x, y)
}

// PackedXY is a free data retrieval call binding the contract method 0xa8696223.
//
// Solidity: function packedXY(int128 x, int128 y) pure returns(uint256 _packedXY)
func (_Peopleland *PeoplelandCallerSession) PackedXY(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.PackedXY(&_Peopleland.CallOpts, x, y)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Peopleland *PeoplelandCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Peopleland *PeoplelandSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Peopleland.Contract.SupportsInterface(&_Peopleland.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Peopleland *PeoplelandCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Peopleland.Contract.SupportsInterface(&_Peopleland.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peopleland *PeoplelandCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peopleland *PeoplelandSession) Symbol() (string, error) {
	return _Peopleland.Contract.Symbol(&_Peopleland.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peopleland *PeoplelandCallerSession) Symbol() (string, error) {
	return _Peopleland.Contract.Symbol(&_Peopleland.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Peopleland *PeoplelandCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Peopleland *PeoplelandSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.TokenByIndex(&_Peopleland.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Peopleland *PeoplelandCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.TokenByIndex(&_Peopleland.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Peopleland *PeoplelandCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Peopleland *PeoplelandSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.TokenOfOwnerByIndex(&_Peopleland.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Peopleland *PeoplelandCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Peopleland.Contract.TokenOfOwnerByIndex(&_Peopleland.CallOpts, owner, index)
}

// TokenSVGAddress is a free data retrieval call binding the contract method 0x0b4612fb.
//
// Solidity: function tokenSVGAddress() view returns(address)
func (_Peopleland *PeoplelandCaller) TokenSVGAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "tokenSVGAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenSVGAddress is a free data retrieval call binding the contract method 0x0b4612fb.
//
// Solidity: function tokenSVGAddress() view returns(address)
func (_Peopleland *PeoplelandSession) TokenSVGAddress() (common.Address, error) {
	return _Peopleland.Contract.TokenSVGAddress(&_Peopleland.CallOpts)
}

// TokenSVGAddress is a free data retrieval call binding the contract method 0x0b4612fb.
//
// Solidity: function tokenSVGAddress() view returns(address)
func (_Peopleland *PeoplelandCallerSession) TokenSVGAddress() (common.Address, error) {
	return _Peopleland.Contract.TokenSVGAddress(&_Peopleland.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_Peopleland *PeoplelandCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_Peopleland *PeoplelandSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Peopleland.Contract.TokenURI(&_Peopleland.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_Peopleland *PeoplelandCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Peopleland.Contract.TokenURI(&_Peopleland.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peopleland *PeoplelandCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peopleland.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peopleland *PeoplelandSession) TotalSupply() (*big.Int, error) {
	return _Peopleland.Contract.TotalSupply(&_Peopleland.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peopleland *PeoplelandCallerSession) TotalSupply() (*big.Int, error) {
	return _Peopleland.Contract.TotalSupply(&_Peopleland.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.Approve(&_Peopleland.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.Approve(&_Peopleland.TransactOpts, to, tokenId)
}

// GetAllEth is a paid mutator transaction binding the contract method 0x3223c16f.
//
// Solidity: function getAllEth() returns()
func (_Peopleland *PeoplelandTransactor) GetAllEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "getAllEth")
}

// GetAllEth is a paid mutator transaction binding the contract method 0x3223c16f.
//
// Solidity: function getAllEth() returns()
func (_Peopleland *PeoplelandSession) GetAllEth() (*types.Transaction, error) {
	return _Peopleland.Contract.GetAllEth(&_Peopleland.TransactOpts)
}

// GetAllEth is a paid mutator transaction binding the contract method 0x3223c16f.
//
// Solidity: function getAllEth() returns()
func (_Peopleland *PeoplelandTransactorSession) GetAllEth() (*types.Transaction, error) {
	return _Peopleland.Contract.GetAllEth(&_Peopleland.TransactOpts)
}

// GetEth is a paid mutator transaction binding the contract method 0x8e3073a6.
//
// Solidity: function getEth(uint256 value) returns()
func (_Peopleland *PeoplelandTransactor) GetEth(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "getEth", value)
}

// GetEth is a paid mutator transaction binding the contract method 0x8e3073a6.
//
// Solidity: function getEth(uint256 value) returns()
func (_Peopleland *PeoplelandSession) GetEth(value *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.GetEth(&_Peopleland.TransactOpts, value)
}

// GetEth is a paid mutator transaction binding the contract method 0x8e3073a6.
//
// Solidity: function getEth(uint256 value) returns()
func (_Peopleland *PeoplelandTransactorSession) GetEth(value *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.GetEth(&_Peopleland.TransactOpts, value)
}

// GiveTo is a paid mutator transaction binding the contract method 0xb3271d65.
//
// Solidity: function giveTo(int128 x, int128 y, address givedAddress) returns()
func (_Peopleland *PeoplelandTransactor) GiveTo(opts *bind.TransactOpts, x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "giveTo", x, y, givedAddress)
}

// GiveTo is a paid mutator transaction binding the contract method 0xb3271d65.
//
// Solidity: function giveTo(int128 x, int128 y, address givedAddress) returns()
func (_Peopleland *PeoplelandSession) GiveTo(x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.GiveTo(&_Peopleland.TransactOpts, x, y, givedAddress)
}

// GiveTo is a paid mutator transaction binding the contract method 0xb3271d65.
//
// Solidity: function giveTo(int128 x, int128 y, address givedAddress) returns()
func (_Peopleland *PeoplelandTransactorSession) GiveTo(x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.GiveTo(&_Peopleland.TransactOpts, x, y, givedAddress)
}

// Mint is a paid mutator transaction binding the contract method 0xb154d6e9.
//
// Solidity: function mint(int128 x, int128 y) payable returns()
func (_Peopleland *PeoplelandTransactor) Mint(opts *bind.TransactOpts, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mint", x, y)
}

// Mint is a paid mutator transaction binding the contract method 0xb154d6e9.
//
// Solidity: function mint(int128 x, int128 y) payable returns()
func (_Peopleland *PeoplelandSession) Mint(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.Mint(&_Peopleland.TransactOpts, x, y)
}

// Mint is a paid mutator transaction binding the contract method 0xb154d6e9.
//
// Solidity: function mint(int128 x, int128 y) payable returns()
func (_Peopleland *PeoplelandTransactorSession) Mint(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.Mint(&_Peopleland.TransactOpts, x, y)
}

// Mint2 is a paid mutator transaction binding the contract method 0x267d1352.
//
// Solidity: function mint2(int128 x1, int128 y1, int128 x2, int128 y2) payable returns()
func (_Peopleland *PeoplelandTransactor) Mint2(opts *bind.TransactOpts, x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mint2", x1, y1, x2, y2)
}

// Mint2 is a paid mutator transaction binding the contract method 0x267d1352.
//
// Solidity: function mint2(int128 x1, int128 y1, int128 x2, int128 y2) payable returns()
func (_Peopleland *PeoplelandSession) Mint2(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.Mint2(&_Peopleland.TransactOpts, x1, y1, x2, y2)
}

// Mint2 is a paid mutator transaction binding the contract method 0x267d1352.
//
// Solidity: function mint2(int128 x1, int128 y1, int128 x2, int128 y2) payable returns()
func (_Peopleland *PeoplelandTransactorSession) Mint2(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.Mint2(&_Peopleland.TransactOpts, x1, y1, x2, y2)
}

// Mint2AndGiveTo is a paid mutator transaction binding the contract method 0x940bbf04.
//
// Solidity: function mint2AndGiveTo(int128 x1, int128 y1, address givedAddress1, int128 x2, int128 y2, address givedAddress2) payable returns()
func (_Peopleland *PeoplelandTransactor) Mint2AndGiveTo(opts *bind.TransactOpts, x1 *big.Int, y1 *big.Int, givedAddress1 common.Address, x2 *big.Int, y2 *big.Int, givedAddress2 common.Address) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mint2AndGiveTo", x1, y1, givedAddress1, x2, y2, givedAddress2)
}

// Mint2AndGiveTo is a paid mutator transaction binding the contract method 0x940bbf04.
//
// Solidity: function mint2AndGiveTo(int128 x1, int128 y1, address givedAddress1, int128 x2, int128 y2, address givedAddress2) payable returns()
func (_Peopleland *PeoplelandSession) Mint2AndGiveTo(x1 *big.Int, y1 *big.Int, givedAddress1 common.Address, x2 *big.Int, y2 *big.Int, givedAddress2 common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.Mint2AndGiveTo(&_Peopleland.TransactOpts, x1, y1, givedAddress1, x2, y2, givedAddress2)
}

// Mint2AndGiveTo is a paid mutator transaction binding the contract method 0x940bbf04.
//
// Solidity: function mint2AndGiveTo(int128 x1, int128 y1, address givedAddress1, int128 x2, int128 y2, address givedAddress2) payable returns()
func (_Peopleland *PeoplelandTransactorSession) Mint2AndGiveTo(x1 *big.Int, y1 *big.Int, givedAddress1 common.Address, x2 *big.Int, y2 *big.Int, givedAddress2 common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.Mint2AndGiveTo(&_Peopleland.TransactOpts, x1, y1, givedAddress1, x2, y2, givedAddress2)
}

// MintAndGiveTo is a paid mutator transaction binding the contract method 0xee830315.
//
// Solidity: function mintAndGiveTo(int128 x, int128 y, address givedAddress) payable returns()
func (_Peopleland *PeoplelandTransactor) MintAndGiveTo(opts *bind.TransactOpts, x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mintAndGiveTo", x, y, givedAddress)
}

// MintAndGiveTo is a paid mutator transaction binding the contract method 0xee830315.
//
// Solidity: function mintAndGiveTo(int128 x, int128 y, address givedAddress) payable returns()
func (_Peopleland *PeoplelandSession) MintAndGiveTo(x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.MintAndGiveTo(&_Peopleland.TransactOpts, x, y, givedAddress)
}

// MintAndGiveTo is a paid mutator transaction binding the contract method 0xee830315.
//
// Solidity: function mintAndGiveTo(int128 x, int128 y, address givedAddress) payable returns()
func (_Peopleland *PeoplelandTransactorSession) MintAndGiveTo(x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.MintAndGiveTo(&_Peopleland.TransactOpts, x, y, givedAddress)
}

// MintAndGiveToWithSlogan is a paid mutator transaction binding the contract method 0x1dc7d0c5.
//
// Solidity: function mintAndGiveToWithSlogan(int128 x, int128 y, address givedAddress, string slogan) payable returns()
func (_Peopleland *PeoplelandTransactor) MintAndGiveToWithSlogan(opts *bind.TransactOpts, x *big.Int, y *big.Int, givedAddress common.Address, slogan string) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mintAndGiveToWithSlogan", x, y, givedAddress, slogan)
}

// MintAndGiveToWithSlogan is a paid mutator transaction binding the contract method 0x1dc7d0c5.
//
// Solidity: function mintAndGiveToWithSlogan(int128 x, int128 y, address givedAddress, string slogan) payable returns()
func (_Peopleland *PeoplelandSession) MintAndGiveToWithSlogan(x *big.Int, y *big.Int, givedAddress common.Address, slogan string) (*types.Transaction, error) {
	return _Peopleland.Contract.MintAndGiveToWithSlogan(&_Peopleland.TransactOpts, x, y, givedAddress, slogan)
}

// MintAndGiveToWithSlogan is a paid mutator transaction binding the contract method 0x1dc7d0c5.
//
// Solidity: function mintAndGiveToWithSlogan(int128 x, int128 y, address givedAddress, string slogan) payable returns()
func (_Peopleland *PeoplelandTransactorSession) MintAndGiveToWithSlogan(x *big.Int, y *big.Int, givedAddress common.Address, slogan string) (*types.Transaction, error) {
	return _Peopleland.Contract.MintAndGiveToWithSlogan(&_Peopleland.TransactOpts, x, y, givedAddress, slogan)
}

// MintToBuilderByOwner is a paid mutator transaction binding the contract method 0x67e72603.
//
// Solidity: function mintToBuilderByOwner(int128 x, int128 y, address givedAddress) returns()
func (_Peopleland *PeoplelandTransactor) MintToBuilderByOwner(opts *bind.TransactOpts, x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mintToBuilderByOwner", x, y, givedAddress)
}

// MintToBuilderByOwner is a paid mutator transaction binding the contract method 0x67e72603.
//
// Solidity: function mintToBuilderByOwner(int128 x, int128 y, address givedAddress) returns()
func (_Peopleland *PeoplelandSession) MintToBuilderByOwner(x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.MintToBuilderByOwner(&_Peopleland.TransactOpts, x, y, givedAddress)
}

// MintToBuilderByOwner is a paid mutator transaction binding the contract method 0x67e72603.
//
// Solidity: function mintToBuilderByOwner(int128 x, int128 y, address givedAddress) returns()
func (_Peopleland *PeoplelandTransactorSession) MintToBuilderByOwner(x *big.Int, y *big.Int, givedAddress common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.MintToBuilderByOwner(&_Peopleland.TransactOpts, x, y, givedAddress)
}

// MintToBuilderByOwnerWithSlogan is a paid mutator transaction binding the contract method 0x08c20a78.
//
// Solidity: function mintToBuilderByOwnerWithSlogan(int128 x, int128 y, address givedAddress, string slogan) returns()
func (_Peopleland *PeoplelandTransactor) MintToBuilderByOwnerWithSlogan(opts *bind.TransactOpts, x *big.Int, y *big.Int, givedAddress common.Address, slogan string) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mintToBuilderByOwnerWithSlogan", x, y, givedAddress, slogan)
}

// MintToBuilderByOwnerWithSlogan is a paid mutator transaction binding the contract method 0x08c20a78.
//
// Solidity: function mintToBuilderByOwnerWithSlogan(int128 x, int128 y, address givedAddress, string slogan) returns()
func (_Peopleland *PeoplelandSession) MintToBuilderByOwnerWithSlogan(x *big.Int, y *big.Int, givedAddress common.Address, slogan string) (*types.Transaction, error) {
	return _Peopleland.Contract.MintToBuilderByOwnerWithSlogan(&_Peopleland.TransactOpts, x, y, givedAddress, slogan)
}

// MintToBuilderByOwnerWithSlogan is a paid mutator transaction binding the contract method 0x08c20a78.
//
// Solidity: function mintToBuilderByOwnerWithSlogan(int128 x, int128 y, address givedAddress, string slogan) returns()
func (_Peopleland *PeoplelandTransactorSession) MintToBuilderByOwnerWithSlogan(x *big.Int, y *big.Int, givedAddress common.Address, slogan string) (*types.Transaction, error) {
	return _Peopleland.Contract.MintToBuilderByOwnerWithSlogan(&_Peopleland.TransactOpts, x, y, givedAddress, slogan)
}

// MintToSelf is a paid mutator transaction binding the contract method 0x9d573677.
//
// Solidity: function mintToSelf(int128 x, int128 y, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) returns()
func (_Peopleland *PeoplelandTransactor) MintToSelf(opts *bind.TransactOpts, x *big.Int, y *big.Int, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "mintToSelf", x, y, messageHash, v, r, s)
}

// MintToSelf is a paid mutator transaction binding the contract method 0x9d573677.
//
// Solidity: function mintToSelf(int128 x, int128 y, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) returns()
func (_Peopleland *PeoplelandSession) MintToSelf(x *big.Int, y *big.Int, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Peopleland.Contract.MintToSelf(&_Peopleland.TransactOpts, x, y, messageHash, v, r, s)
}

// MintToSelf is a paid mutator transaction binding the contract method 0x9d573677.
//
// Solidity: function mintToSelf(int128 x, int128 y, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) returns()
func (_Peopleland *PeoplelandTransactorSession) MintToSelf(x *big.Int, y *big.Int, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Peopleland.Contract.MintToSelf(&_Peopleland.TransactOpts, x, y, messageHash, v, r, s)
}

// OpenMintSelfSwitch is a paid mutator transaction binding the contract method 0x11096acd.
//
// Solidity: function openMintSelfSwitch() returns()
func (_Peopleland *PeoplelandTransactor) OpenMintSelfSwitch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "openMintSelfSwitch")
}

// OpenMintSelfSwitch is a paid mutator transaction binding the contract method 0x11096acd.
//
// Solidity: function openMintSelfSwitch() returns()
func (_Peopleland *PeoplelandSession) OpenMintSelfSwitch() (*types.Transaction, error) {
	return _Peopleland.Contract.OpenMintSelfSwitch(&_Peopleland.TransactOpts)
}

// OpenMintSelfSwitch is a paid mutator transaction binding the contract method 0x11096acd.
//
// Solidity: function openMintSelfSwitch() returns()
func (_Peopleland *PeoplelandTransactorSession) OpenMintSelfSwitch() (*types.Transaction, error) {
	return _Peopleland.Contract.OpenMintSelfSwitch(&_Peopleland.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Peopleland *PeoplelandTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Peopleland *PeoplelandSession) RenounceOwnership() (*types.Transaction, error) {
	return _Peopleland.Contract.RenounceOwnership(&_Peopleland.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Peopleland *PeoplelandTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Peopleland.Contract.RenounceOwnership(&_Peopleland.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.SafeTransferFrom(&_Peopleland.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.SafeTransferFrom(&_Peopleland.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Peopleland *PeoplelandTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Peopleland *PeoplelandSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Peopleland.Contract.SafeTransferFrom0(&_Peopleland.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Peopleland *PeoplelandTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Peopleland.Contract.SafeTransferFrom0(&_Peopleland.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Peopleland *PeoplelandTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Peopleland *PeoplelandSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Peopleland.Contract.SetApprovalForAll(&_Peopleland.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Peopleland *PeoplelandTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Peopleland.Contract.SetApprovalForAll(&_Peopleland.TransactOpts, operator, approved)
}

// SetSlogan is a paid mutator transaction binding the contract method 0xd019b713.
//
// Solidity: function setSlogan(int128 x, int128 y, string slogan) returns()
func (_Peopleland *PeoplelandTransactor) SetSlogan(opts *bind.TransactOpts, x *big.Int, y *big.Int, slogan string) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "setSlogan", x, y, slogan)
}

// SetSlogan is a paid mutator transaction binding the contract method 0xd019b713.
//
// Solidity: function setSlogan(int128 x, int128 y, string slogan) returns()
func (_Peopleland *PeoplelandSession) SetSlogan(x *big.Int, y *big.Int, slogan string) (*types.Transaction, error) {
	return _Peopleland.Contract.SetSlogan(&_Peopleland.TransactOpts, x, y, slogan)
}

// SetSlogan is a paid mutator transaction binding the contract method 0xd019b713.
//
// Solidity: function setSlogan(int128 x, int128 y, string slogan) returns()
func (_Peopleland *PeoplelandTransactorSession) SetSlogan(x *big.Int, y *big.Int, slogan string) (*types.Transaction, error) {
	return _Peopleland.Contract.SetSlogan(&_Peopleland.TransactOpts, x, y, slogan)
}

// SetTokenSVGAddress is a paid mutator transaction binding the contract method 0x03814d1c.
//
// Solidity: function setTokenSVGAddress(address _attr) returns()
func (_Peopleland *PeoplelandTransactor) SetTokenSVGAddress(opts *bind.TransactOpts, _attr common.Address) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "setTokenSVGAddress", _attr)
}

// SetTokenSVGAddress is a paid mutator transaction binding the contract method 0x03814d1c.
//
// Solidity: function setTokenSVGAddress(address _attr) returns()
func (_Peopleland *PeoplelandSession) SetTokenSVGAddress(_attr common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.SetTokenSVGAddress(&_Peopleland.TransactOpts, _attr)
}

// SetTokenSVGAddress is a paid mutator transaction binding the contract method 0x03814d1c.
//
// Solidity: function setTokenSVGAddress(address _attr) returns()
func (_Peopleland *PeoplelandTransactorSession) SetTokenSVGAddress(_attr common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.SetTokenSVGAddress(&_Peopleland.TransactOpts, _attr)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.TransferFrom(&_Peopleland.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Peopleland *PeoplelandTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Peopleland.Contract.TransferFrom(&_Peopleland.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Peopleland *PeoplelandTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Peopleland.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Peopleland *PeoplelandSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.TransferOwnership(&_Peopleland.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Peopleland *PeoplelandTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Peopleland.Contract.TransferOwnership(&_Peopleland.TransactOpts, newOwner)
}

// PeoplelandApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Peopleland contract.
type PeoplelandApprovalIterator struct {
	Event *PeoplelandApproval // Event containing the contract specifics and raw log

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
func (it *PeoplelandApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeoplelandApproval)
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
		it.Event = new(PeoplelandApproval)
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
func (it *PeoplelandApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeoplelandApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeoplelandApproval represents a Approval event raised by the Peopleland contract.
type PeoplelandApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Peopleland *PeoplelandFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PeoplelandApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Peopleland.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PeoplelandApprovalIterator{contract: _Peopleland.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Peopleland *PeoplelandFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PeoplelandApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Peopleland.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeoplelandApproval)
				if err := _Peopleland.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Peopleland *PeoplelandFilterer) ParseApproval(log types.Log) (*PeoplelandApproval, error) {
	event := new(PeoplelandApproval)
	if err := _Peopleland.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeoplelandApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Peopleland contract.
type PeoplelandApprovalForAllIterator struct {
	Event *PeoplelandApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PeoplelandApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeoplelandApprovalForAll)
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
		it.Event = new(PeoplelandApprovalForAll)
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
func (it *PeoplelandApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeoplelandApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeoplelandApprovalForAll represents a ApprovalForAll event raised by the Peopleland contract.
type PeoplelandApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Peopleland *PeoplelandFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PeoplelandApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Peopleland.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PeoplelandApprovalForAllIterator{contract: _Peopleland.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Peopleland *PeoplelandFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PeoplelandApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Peopleland.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeoplelandApprovalForAll)
				if err := _Peopleland.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Peopleland *PeoplelandFilterer) ParseApprovalForAll(log types.Log) (*PeoplelandApprovalForAll, error) {
	event := new(PeoplelandApprovalForAll)
	if err := _Peopleland.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeoplelandGiveToIterator is returned from FilterGiveTo and is used to iterate over the raw logs and unpacked data for GiveTo events raised by the Peopleland contract.
type PeoplelandGiveToIterator struct {
	Event *PeoplelandGiveTo // Event containing the contract specifics and raw log

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
func (it *PeoplelandGiveToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeoplelandGiveTo)
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
		it.Event = new(PeoplelandGiveTo)
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
func (it *PeoplelandGiveToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeoplelandGiveToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeoplelandGiveTo represents a GiveTo event raised by the Peopleland contract.
type PeoplelandGiveTo struct {
	X            *big.Int
	Y            *big.Int
	GivedAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGiveTo is a free log retrieval operation binding the contract event 0x315581f8fced1d4d8d321a598f896b9f857c52dfefae7188fbfde64900d62d9b.
//
// Solidity: event GiveTo(int128 x, int128 y, address givedAddress)
func (_Peopleland *PeoplelandFilterer) FilterGiveTo(opts *bind.FilterOpts) (*PeoplelandGiveToIterator, error) {

	logs, sub, err := _Peopleland.contract.FilterLogs(opts, "GiveTo")
	if err != nil {
		return nil, err
	}
	return &PeoplelandGiveToIterator{contract: _Peopleland.contract, event: "GiveTo", logs: logs, sub: sub}, nil
}

// WatchGiveTo is a free log subscription operation binding the contract event 0x315581f8fced1d4d8d321a598f896b9f857c52dfefae7188fbfde64900d62d9b.
//
// Solidity: event GiveTo(int128 x, int128 y, address givedAddress)
func (_Peopleland *PeoplelandFilterer) WatchGiveTo(opts *bind.WatchOpts, sink chan<- *PeoplelandGiveTo) (event.Subscription, error) {

	logs, sub, err := _Peopleland.contract.WatchLogs(opts, "GiveTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeoplelandGiveTo)
				if err := _Peopleland.contract.UnpackLog(event, "GiveTo", log); err != nil {
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

// ParseGiveTo is a log parse operation binding the contract event 0x315581f8fced1d4d8d321a598f896b9f857c52dfefae7188fbfde64900d62d9b.
//
// Solidity: event GiveTo(int128 x, int128 y, address givedAddress)
func (_Peopleland *PeoplelandFilterer) ParseGiveTo(log types.Log) (*PeoplelandGiveTo, error) {
	event := new(PeoplelandGiveTo)
	if err := _Peopleland.contract.UnpackLog(event, "GiveTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeoplelandMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Peopleland contract.
type PeoplelandMintIterator struct {
	Event *PeoplelandMint // Event containing the contract specifics and raw log

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
func (it *PeoplelandMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeoplelandMint)
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
		it.Event = new(PeoplelandMint)
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
func (it *PeoplelandMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeoplelandMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeoplelandMint represents a Mint event raised by the Peopleland contract.
type PeoplelandMint struct {
	X             *big.Int
	Y             *big.Int
	MintedAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xdc098a7ae450e373b74167b44944e42b9782c72ad4b58d0e2f6999b93e1c7ead.
//
// Solidity: event Mint(int128 x, int128 y, address mintedAddress)
func (_Peopleland *PeoplelandFilterer) FilterMint(opts *bind.FilterOpts) (*PeoplelandMintIterator, error) {

	logs, sub, err := _Peopleland.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &PeoplelandMintIterator{contract: _Peopleland.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xdc098a7ae450e373b74167b44944e42b9782c72ad4b58d0e2f6999b93e1c7ead.
//
// Solidity: event Mint(int128 x, int128 y, address mintedAddress)
func (_Peopleland *PeoplelandFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PeoplelandMint) (event.Subscription, error) {

	logs, sub, err := _Peopleland.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeoplelandMint)
				if err := _Peopleland.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xdc098a7ae450e373b74167b44944e42b9782c72ad4b58d0e2f6999b93e1c7ead.
//
// Solidity: event Mint(int128 x, int128 y, address mintedAddress)
func (_Peopleland *PeoplelandFilterer) ParseMint(log types.Log) (*PeoplelandMint, error) {
	event := new(PeoplelandMint)
	if err := _Peopleland.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeoplelandOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Peopleland contract.
type PeoplelandOwnershipTransferredIterator struct {
	Event *PeoplelandOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PeoplelandOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeoplelandOwnershipTransferred)
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
		it.Event = new(PeoplelandOwnershipTransferred)
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
func (it *PeoplelandOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeoplelandOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeoplelandOwnershipTransferred represents a OwnershipTransferred event raised by the Peopleland contract.
type PeoplelandOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Peopleland *PeoplelandFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PeoplelandOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Peopleland.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PeoplelandOwnershipTransferredIterator{contract: _Peopleland.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Peopleland *PeoplelandFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PeoplelandOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Peopleland.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeoplelandOwnershipTransferred)
				if err := _Peopleland.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Peopleland *PeoplelandFilterer) ParseOwnershipTransferred(log types.Log) (*PeoplelandOwnershipTransferred, error) {
	event := new(PeoplelandOwnershipTransferred)
	if err := _Peopleland.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeoplelandSetSloganIterator is returned from FilterSetSlogan and is used to iterate over the raw logs and unpacked data for SetSlogan events raised by the Peopleland contract.
type PeoplelandSetSloganIterator struct {
	Event *PeoplelandSetSlogan // Event containing the contract specifics and raw log

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
func (it *PeoplelandSetSloganIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeoplelandSetSlogan)
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
		it.Event = new(PeoplelandSetSlogan)
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
func (it *PeoplelandSetSloganIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeoplelandSetSloganIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeoplelandSetSlogan represents a SetSlogan event raised by the Peopleland contract.
type PeoplelandSetSlogan struct {
	X      *big.Int
	Y      *big.Int
	Slogan string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSlogan is a free log retrieval operation binding the contract event 0x33aa80b38af4e0a1e2607b5743673d59a16ae46aecc240c828d6e71bddaf6762.
//
// Solidity: event SetSlogan(int128 x, int128 y, string slogan)
func (_Peopleland *PeoplelandFilterer) FilterSetSlogan(opts *bind.FilterOpts) (*PeoplelandSetSloganIterator, error) {

	logs, sub, err := _Peopleland.contract.FilterLogs(opts, "SetSlogan")
	if err != nil {
		return nil, err
	}
	return &PeoplelandSetSloganIterator{contract: _Peopleland.contract, event: "SetSlogan", logs: logs, sub: sub}, nil
}

// WatchSetSlogan is a free log subscription operation binding the contract event 0x33aa80b38af4e0a1e2607b5743673d59a16ae46aecc240c828d6e71bddaf6762.
//
// Solidity: event SetSlogan(int128 x, int128 y, string slogan)
func (_Peopleland *PeoplelandFilterer) WatchSetSlogan(opts *bind.WatchOpts, sink chan<- *PeoplelandSetSlogan) (event.Subscription, error) {

	logs, sub, err := _Peopleland.contract.WatchLogs(opts, "SetSlogan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeoplelandSetSlogan)
				if err := _Peopleland.contract.UnpackLog(event, "SetSlogan", log); err != nil {
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

// ParseSetSlogan is a log parse operation binding the contract event 0x33aa80b38af4e0a1e2607b5743673d59a16ae46aecc240c828d6e71bddaf6762.
//
// Solidity: event SetSlogan(int128 x, int128 y, string slogan)
func (_Peopleland *PeoplelandFilterer) ParseSetSlogan(log types.Log) (*PeoplelandSetSlogan, error) {
	event := new(PeoplelandSetSlogan)
	if err := _Peopleland.contract.UnpackLog(event, "SetSlogan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeoplelandTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Peopleland contract.
type PeoplelandTransferIterator struct {
	Event *PeoplelandTransfer // Event containing the contract specifics and raw log

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
func (it *PeoplelandTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeoplelandTransfer)
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
		it.Event = new(PeoplelandTransfer)
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
func (it *PeoplelandTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeoplelandTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeoplelandTransfer represents a Transfer event raised by the Peopleland contract.
type PeoplelandTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Peopleland *PeoplelandFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PeoplelandTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Peopleland.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PeoplelandTransferIterator{contract: _Peopleland.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Peopleland *PeoplelandFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PeoplelandTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Peopleland.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeoplelandTransfer)
				if err := _Peopleland.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Peopleland *PeoplelandFilterer) ParseTransfer(log types.Log) (*PeoplelandTransfer, error) {
	event := new(PeoplelandTransfer)
	if err := _Peopleland.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
