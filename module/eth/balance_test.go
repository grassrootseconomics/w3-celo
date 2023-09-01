package eth_test

import (
	"math/big"
	"testing"

	"github.com/grassrootseconomics/w3-celo"
	"github.com/grassrootseconomics/w3-celo/module/eth"
	"github.com/grassrootseconomics/w3-celo/rpctest"
)

func TestBalance(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "get_balance",
			Call:    eth.Balance(w3.A("0x000000000000000000000000000000000000c0Fe"), nil),
			WantRet: *w3.I("1 ether"),
		},
		{
			Golden:  "get_balance__at_block",
			Call:    eth.Balance(w3.A("0x000000000000000000000000000000000000c0Fe"), big.NewInt(255)),
			WantRet: *w3.I("0.1 ether"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
