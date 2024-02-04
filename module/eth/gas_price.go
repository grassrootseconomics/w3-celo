package eth

import (
	"math/big"

	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// GasPrice requests the current gas price in wei.
func GasPrice() w3types.RPCCallerFactory[big.Int] {
	return module.NewFactory(
		"eth_gasPrice",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
