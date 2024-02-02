package eth

import (
	"math/big"

	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// GasTipCap requests the currently suggested gas tip cap after EIP-1559 to
// allow a timely execution of a transaction.
func GasTipCap() w3types.RPCCallerFactory[big.Int] {
	return module.NewFactory(
		"eth_maxPriorityFeePerGas",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
