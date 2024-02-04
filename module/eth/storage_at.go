package eth

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// StorageAt requests the storage of the given common.Address addr at the
// given common.Hash slot at the given blockNumber. If block number is nil, the
// slot at the latest known block is requested.
func StorageAt(addr common.Address, slot common.Hash, blockNumber *big.Int) w3types.RPCCallerFactory[common.Hash] {
	return module.NewFactory[common.Hash](
		"eth_getStorageAt",
		[]any{addr, slot, module.BlockNumberArg(blockNumber)},
	)
}
