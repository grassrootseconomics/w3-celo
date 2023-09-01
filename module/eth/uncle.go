package eth

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/common/hexutil"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// UncleByBlockHashAndIndex requests the uncle of the block with the given hash
// at the given index.
func UncleByBlockHashAndIndex(hash common.Hash, index uint) w3types.CallerFactory[types.Header] {
	return module.NewFactory[types.Header](
		"eth_getUncleByBlockHashAndIndex",
		[]any{hash, hexutil.Uint(index)},
	)
}

// UncleByBlockNumberAndIndex requests the uncle of the block with the given
// number at the given index.
func UncleByBlockNumberAndIndex(number *big.Int, index uint) w3types.CallerFactory[types.Header] {
	return module.NewFactory[types.Header](
		"eth_getUncleByBlockNumberAndIndex",
		[]any{module.BlockNumberArg(number), hexutil.Uint(index)},
	)
}

// UncleCountByBlockHash requests the number of uncles of the block with the
// given hash.
func UncleCountByBlockHash(hash common.Hash) w3types.CallerFactory[uint] {
	return module.NewFactory(
		"eth_getUncleCountByBlockHash",
		[]any{hash},
		module.WithRetWrapper(module.HexUintRetWrapper),
	)
}

// UncleCountByBlockNumber requests the number of uncles of the block with the
// given number.
func UncleCountByBlockNumber(number *big.Int) w3types.CallerFactory[uint] {
	return module.NewFactory(
		"eth_getUncleCountByBlockNumber",
		[]any{module.BlockNumberArg(number)},
		module.WithRetWrapper(module.HexUintRetWrapper),
	)
}
