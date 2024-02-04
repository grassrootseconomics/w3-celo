package eth

import (
	"fmt"

	"github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/rpc"
	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// Logs requests the logs of the given celo.FilterQuery q.
func Logs(q celo.FilterQuery) w3types.CallerFactory[[]types.Log] {
	return &logsFactory{filterQuery: q}
}

type logsFactory struct {
	// args
	filterQuery celo.FilterQuery

	// returns
	returns *[]types.Log
}

func (f *logsFactory) Returns(logs *[]types.Log) w3types.RPCCaller {
	f.returns = logs
	return f
}

// CreateRequest implements the w3types.RequestCreator interface.
func (f *logsFactory) CreateRequest() (rpc.BatchElem, error) {
	arg, err := toFilterArg(f.filterQuery)
	if err != nil {
		return rpc.BatchElem{}, err
	}

	return rpc.BatchElem{
		Method: "eth_getLogs",
		Args:   []any{arg},
		Result: f.returns,
	}, nil
}

// HandleResponse implements the w3types.ResponseHandler interface.
func (f *logsFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}

func toFilterArg(q celo.FilterQuery) (any, error) {
	arg := map[string]any{
		"topics": q.Topics,
	}
	if len(q.Addresses) > 0 {
		arg["address"] = q.Addresses
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = module.BlockNumberArg(q.FromBlock)
		}
		arg["toBlock"] = module.BlockNumberArg(q.ToBlock)
	}
	return arg, nil
}
