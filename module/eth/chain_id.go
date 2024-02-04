package eth

import (
	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// ChainID requests the chains ID.
func ChainID() w3types.RPCCallerFactory[uint64] {
	return module.NewFactory(
		"eth_chainId",
		nil,
		module.WithRetWrapper(module.HexUint64RetWrapper),
	)
}
