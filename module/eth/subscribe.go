package eth

import (
	"github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// NewHeads subscribes to notifications of updates to the blockchain head.
func NewHeads(ch chan<- *types.Header) w3types.RPCSubscriber {
	return &ethSubscription[*types.Header]{ch, []any{"newHeads"}, nil}
}

// PendingTransactions subscribes to notifications about new pending transactions in the transaction pool.
func PendingTransactions(ch chan<- *types.Transaction) w3types.RPCSubscriber {
	return &ethSubscription[*types.Transaction]{ch, []any{"newPendingTransactions", true}, nil}
}

// NewLogs subscribes to notifications about logs that match the given filter query.
func NewLogs(ch chan<- *types.Log, q celo.FilterQuery) w3types.RPCSubscriber {
	arg, err := toFilterArg(q)
	return &ethSubscription[*types.Log]{ch, []any{"logs", arg}, err}
}

type ethSubscription[T any] struct {
	ch     chan<- T
	params []any
	err    error
}

func (s *ethSubscription[T]) CreateRequest() (string, any, []any, error) {
	return "eth", s.ch, s.params, s.err
}
