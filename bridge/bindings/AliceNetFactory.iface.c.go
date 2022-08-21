// Generated by ifacemaker. DO NOT EDIT.

package bindings

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// IAliceNetFactoryCaller ...
type IAliceNetFactoryCaller interface {
	// Contracts is a free data retrieval call binding the contract method 0x6c0f79b6.
	//
	// Solidity: function contracts() view returns(bytes32[] contracts_)
	Contracts(opts *bind.CallOpts) ([][32]byte, error)
	// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
	//
	// Solidity: function getImplementation() view returns(address)
	GetImplementation(opts *bind.CallOpts) (common.Address, error)
	// GetMetamorphicContractAddress is a free data retrieval call binding the contract method 0x8653a465.
	//
	// Solidity: function getMetamorphicContractAddress(bytes32 _salt, address _factory) pure returns(address)
	GetMetamorphicContractAddress(opts *bind.CallOpts, _salt [32]byte, _factory common.Address) (common.Address, error)
	// GetNumContracts is a free data retrieval call binding the contract method 0xcfe10b30.
	//
	// Solidity: function getNumContracts() view returns(uint256)
	GetNumContracts(opts *bind.CallOpts) (*big.Int, error)
	// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
	//
	// Solidity: function lookup(bytes32 salt_) view returns(address addr)
	Lookup(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error)
	// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
	//
	// Solidity: function owner() view returns(address owner_)
	Owner(opts *bind.CallOpts) (common.Address, error)
}
