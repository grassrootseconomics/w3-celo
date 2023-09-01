package eth_test

import (
	"testing"

	"github.com/grassrootseconomics/w3-celo"
	"github.com/grassrootseconomics/w3-celo/module/eth"
	"github.com/grassrootseconomics/w3-celo/rpctest"
)

func TestCode(t *testing.T) {
	tests := []rpctest.TestCase[[]byte]{
		{
			Golden:  "get_code",
			Call:    eth.Code(w3.A("0x000000000000000000000000000000000000c0DE"), nil),
			WantRet: w3.B("0xdeadbeef"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
