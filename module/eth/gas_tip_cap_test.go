package eth_test

import (
	"math/big"
	"testing"

	"github.com/grassrootseconomics/w3-celo"
	"github.com/grassrootseconomics/w3-celo/module/eth"
	"github.com/grassrootseconomics/w3-celo/rpctest"
)

func TestGasTipCap(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "gas_tip_cap",
			Call:    eth.GasTipCap(),
			WantRet: *w3.I("0xc0fe"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
