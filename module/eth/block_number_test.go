package eth_test

import (
	"math/big"
	"testing"

	"github.com/grassrootseconomics/w3-celo"
	"github.com/grassrootseconomics/w3-celo/module/eth"
	"github.com/grassrootseconomics/w3-celo/rpctest"
)

func TestBlockNumber(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "block_number",
			Call:    eth.BlockNumber(),
			WantRet: *w3.I("0xc0fe"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
