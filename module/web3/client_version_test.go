package web3_test

import (
	"errors"
	"testing"

	"github.com/grassrootseconomics/w3-celo/module/web3"
	"github.com/grassrootseconomics/w3-celo/rpctest"
)

func TestClientVersion(t *testing.T) {
	tests := []rpctest.TestCase[string]{
		{
			Golden:  "client_version",
			Call:    web3.ClientVersion(),
			WantRet: "Geth",
		},
		{
			Golden:  "client_version__err",
			Call:    web3.ClientVersion(),
			WantErr: errors.New("w3: call failed: the method web3_clientVersion does not exist/is not available"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
