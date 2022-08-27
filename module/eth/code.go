package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3/core"
	"github.com/lmittmann/w3/internal/module"
)

// Code requests the code of the given common.Address addr at the given
// blockNumber. If blockNumber is nil, the code at the latest known block is
// requested.
func Code(addr common.Address, blockNumber *big.Int) core.CallerFactory[[]byte] {
	return module.NewFactory(
		"eth_getCode",
		[]any{addr, module.BlockNumberArg(blockNumber)},
		module.WithRetWrapper(module.HexBytesRetWrapper),
	)
}
