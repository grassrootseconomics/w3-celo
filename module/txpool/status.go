package txpool

import (
	"encoding/json"

	"github.com/celo-org/celo-blockchain/common/hexutil"
	"github.com/grassrootseconomics/w3-celo/internal/module"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// Status requests the number of pending and queued transactions in the transaction pool.
func Status() w3types.CallerFactory[StatusResponse] {
	return module.NewFactory[StatusResponse](
		"txpool_status",
		nil,
	)
}

type StatusResponse struct {
	Pending uint
	Queued  uint
}

func (s *StatusResponse) UnmarshalJSON(data []byte) error {
	type statusResponse struct {
		Pending hexutil.Uint `json:"pending"`
		Queued  hexutil.Uint `json:"queued"`
	}

	var dec statusResponse
	if err := json.Unmarshal(data, &dec); err != nil {
		return err
	}
	s.Pending = uint(dec.Pending)
	s.Queued = uint(dec.Queued)
	return nil
}
