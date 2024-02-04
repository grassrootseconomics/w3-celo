package web3

import (
	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// ClientVersion requests the endpoints client version.
func ClientVersion() w3types.RPCCallerFactory[string] {
	return module.NewFactory[string](
		"web3_clientVersion",
		nil,
	)
}
