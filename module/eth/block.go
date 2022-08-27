package eth

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/core"
	"github.com/lmittmann/w3/internal/module"
)

// BlockByHash requests the block with the given hash with full transactions.
func BlockByHash(hash common.Hash) core.CallerFactory[types.Block] {
	return module.NewFactory(
		"eth_getBlockByHash",
		[]any{hash, true},
		module.WithRetWrapper(blockRetWrapper),
	)
}

// BlockByNumber requests the block with the given number with full
// transactions. If number is nil, the latest block is requested.
func BlockByNumber(number *big.Int) core.CallerFactory[types.Block] {
	return module.NewFactory(
		"eth_getBlockByNumber",
		[]any{module.BlockNumberArg(number), true},
		module.WithRetWrapper(blockRetWrapper),
	)
}

// HeaderByHash requests the header with the given hash.
func HeaderByHash(hash common.Hash) core.CallerFactory[types.Header] {
	return module.NewFactory[types.Header](
		"eth_getBlockByHash",
		[]any{hash, false},
	)
}

// HeaderByNumber requests the header with the given number. If number is nil,
// the latest header is requested.
func HeaderByNumber(number *big.Int) core.CallerFactory[types.Header] {
	return module.NewFactory[types.Header](
		"eth_getBlockByNumber",
		[]any{module.BlockNumberArg(number), false},
	)
}

var blockRetWrapper = func(ret *types.Block) any { return (*rpcBlock)(ret) }

type rpcBlock types.Block

func (b *rpcBlock) UnmarshalJSON(data []byte) error {
	type rpcBlockTxs struct {
		Transactions []*types.Transaction `json:"transactions"`
	}

	var header types.Header
	if err := json.Unmarshal(data, &header); err != nil {
		return err
	}

	var blockTxs rpcBlockTxs
	if err := json.Unmarshal(data, &blockTxs); err != nil {
		return err
	}

	block := types.NewBlockWithHeader(&header).WithBody(blockTxs.Transactions, nil)
	*b = (rpcBlock)(*block)
	return nil
}
