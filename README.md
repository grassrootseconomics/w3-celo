# w3-celo

[![Go Reference](https://pkg.go.dev/badge/github.com/grassrootseconomics/w3-celo.svg)](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo)
[![Go Report Card](https://goreportcard.com/badge/github.com/grassrootseconomics/w3-celo)](https://goreportcard.com/report/github.com/grassrootseconomics/w3-celo)
[![Coverage Status](https://coveralls.io/repos/github/grassrootseconomics/w3-celo/badge.svg?branch=master)](https://coveralls.io/github/grassrootseconomics/w3-celo?branch=master)
[![Latest Release](https://img.shields.io/github/v/release/grassrootseconomics/w3-celo)](https://github.com/grassrootseconomics/w3-celo/releases)

<img src="https://user-images.githubusercontent.com/3458786/153202258-24bf253e-5ab0-4efd-a0ed-43dc1bf093c9.png" align="right" alt="W3 Gopher" width="158" height="224">

Package `w3-celo` implements a blazing fast and modular Celo JSON RPC client with
first-class ABI support.

`w3-celo`

* **Batch request** support significantly reduces the duration of requests to
  both remote and local endpoints.
* **ABI** bindings are specified for individual functions using Solidity syntax.
  No need for `abigen` and ABI JSON files.
* **Modular** API allows to create custom RPC method integrations that can be
  used alongside the methods implemented by the package.

`w3-celo` is closely linked to [celo-blockchain](https://github.com/celo-org/celo-blockchain)
and uses a variety of its types, such as [`common.Address`](https://pkg.go.dev/github.com/celo-org/celo-blockchain/common#Address)
or [`types.Transaction`](https://pkg.go.dev/github.com/celo-org/celo-blockchain/core/types#Transaction).

Batch requests with `w3-celo` are up to **85x faster** than sequential requests with
`celo-blockchain/ethclient`.

<details>
<summary>Benchmarks</summary>
<pre>
name               ethclient time/op  w3 time/op  delta
Call_BalanceNonce  78.3ms ± 2%        39.0ms ± 1%  -50.15%  (p=0.000 n=23+22)
Call_Balance100     3.90s ± 5%         0.05s ± 2%  -98.84%  (p=0.000 n=20+24)
Call_BalanceOf100   3.99s ± 3%         0.05s ± 2%  -98.73%  (p=0.000 n=22+23)
Call_Block100       6.89s ± 7%         1.94s ±11%  -71.77%  (p=0.000 n=24+23)
</pre>
</details>

## About

> **Note**
> Check out the [original w3](https://github.com/lmittmann/w3)!

`w3-celo` is a fork of the original library replacing all incompatible Ethereum types with Celo types.

Versioning closely follows the original.

## Install

```
go get github.com/grassrootseconomics/w3-celo
```


## Getting Started

> **Note**
> Check out the [examples](https://github.com/grassrootseconomics/w3-celo/tree/main/examples)!

Connect to an RPC endpoint via HTTP, WebSocket, or IPC using [`Dial`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Dial)
or [`MustDial`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#MustDial).

```go
// Connect (or panic on error)
client := w3.MustDial("https://rpc.ankr.com/celo")
defer client.Close()
```


## Batch Requests

Batch request support in the [`Client`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Client)
allows to send multiple RPC requests in a single HTTP request. The speed gains
to remote endpoints are huge. Fetching 100 blocks in a single batch request
with `w3` is ~80x faster compared to sequential requests with `ethclient`.

Example: Request the nonce and balance of an address in a single request

```go
var (
	addr = w3.A("0x000000000000000000000000000000000000c0Fe")

	nonce   uint64
	balance big.Int
)
err := client.Call(
	eth.Nonce(addr, nil).Returns(&nonce),
	eth.Balance(addr, nil).Returns(&balance),
)
```


## ABI Bindings

ABI bindings in `w3` are specified for individual functions using Solidity
syntax and are usable for any contract that supports that function.

Example: ABI binding for the ERC20-function `balanceOf`

```go
funcBalanceOf := w3.MustNewFunc("balanceOf(address)", "uint256")
```

A [`Func`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Func) can be used to

* encode arguments to the contracts input data ([`Func.EncodeArgs`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Func.EncodeArgs)),
* decode arguments from the contracts input data ([`Func.DecodeArgs`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Func.DecodeArgs)), and
* decode returns form the contracts output data ([`Func.DecodeReturns`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Func.DecodeReturns)).

### Reading Contracts

[`Func`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Func)'s can be used with
[`eth.CallFunc`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo/module/eth#CallFunc)
in the client to read contract data.

```go
var (
	weth9 = w3.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	dai   = w3.A("0x6B175474E89094C44Da98b954EedeAC495271d0F")

	weth9Balance big.Int
	daiBalance   big.Int
)

err := client.Call(
	eth.CallFunc(funcBalanceOf, weth9, addr).Returns(&weth9Balance),
	eth.CallFunc(funcBalanceOf, dai, addr).Returns(&daiBalance),
)
```

### Writing Contracts

Sending a transaction to a contract requires three steps.

1. Encode the transaction input data using [`Func.EncodeArgs`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo#Func.EncodeArgs).

```go
var funcTransfer = w3.MustNewFunc("transfer(address,uint256)", "bool")

input, err := funcTransfer.EncodeArgs(w3.A("0x…"), w3.I("1 ether"))
```

2. Create a signed transaction to the contract using [celo-blockchain/types](https://github.com/celo-org/celo-blockchain).

```go
signer := types.LatestSigner(params.MainnetChainConfig)
tx := types.MustSignNewTx(privKey, signer, &types.CeloDynamicFeeTx{
	To:        w3.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
	Nonce:     0,
	Data:      input,
	Gas:       75000,
	GasFeeCap: w3.I("100 gwei"),
	GasTipCap: w3.I("1 gwei"),
})
```

3. Send the signed transaction.

```go
var txHash common.Hash
err := client.Call(
	eth.SendTx(tx).Returns(&txHash),
)
```


## Custom RPC Methods

Custom RPC methods can be called with the `w3` client by creating a
[`core.Caller`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo/core#Caller)
implementation.
The `w3/module/eth` package can be used as implementation reference.


## Utils

Static addresses, hashes, hex byte slices or `big.Int`'s can be parsed from
strings with the following utility functions.

```go
var (
	addr  = w3.A("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	hash  = w3.H("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
	bytes = w3.B("0x27c5342c")
	big   = w3.I("12.34 ether")
)
```

Note that these functions panic if the string cannot be parsed. Use
[celo-blockchain/common](https://pkg.go.dev/github.com/celo-org/celo-blockchain/common)
to parse strings that may not be valid instead.


## RPC Methods

List of supported RPC methods.

### [`eth`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo/module/eth)

| Method                                    | Go Code
| :---------------------------------------- | :-------
| `eth_blockNumber`                         | `eth.BlockNumber().Returns(blockNumber *big.Int)`
| `eth_call`                                | `eth.Call(msg *w3types.Message, blockNumber *big.Int, overrides w3types.State).Returns(output *[]byte)`<br>`eth.CallFunc(fn core.Func, contract common.Address, args ...any).Returns(returns ...any)`
| `eth_chainId`                             | `eth.ChainID().Returns(chainID *uint64)`
| `eth_createAccessList`                    | `eth.AccessList(msg *w3types.Message, blockNumber *big.Int).Returns(resp *eth.AccessListResponse)`
| `eth_estimateGas`                         | `eth.EstimateGas(msg *w3types.Message, blockNumber *big.Int).Returns(gas *uint64)`
| `eth_gasPrice`                            | `eth.GasPrice().Returns(gasPrice *big.Int)`
| `eth_getBalance`                          | `eth.Balance(addr common.Address, blockNumber *big.Int).Returns(balance *big.Int)`
| `eth_getBlockByHash`                      | `eth.BlockByHash(hash common.Hash).Returns(block *types.Block)`<br>`eth.HeaderByHash(hash common.Hash).Returns(header *types.Header)`
| `eth_getBlockByNumber`                    | `eth.BlockByNumber(number *big.Int).Returns(block *types.Block)`<br>`eth.HeaderByNumber(number *big.Int).Returns(header *types.Header)`
| `eth_getBlockTransactionCountByHash`      | `eth.BlockTxCountByHash(hash common.Hash).Returns(count *uint)`
| `eth_getBlockTransactionCountByNumber`    | `eth.BlockTxCountByNumber(number *big.Int).Returns(count *uint)`
| `eth_getCode`                             | `eth.Code(addr common.Address, blockNumber *big.Int).Returns(code *[]byte)`
| `eth_getLogs`                             | `eth.Logs(q celo.FilterQuery).Returns(logs *[]types.Log)`
| `eth_getStorageAt`                        | `eth.StorageAt(addr common.Address, slot common.Hash, blockNumber *big.Int).Returns(storage *common.Hash)`
| `eth_getTransactionByHash`                | `eth.Tx(hash common.Hash).Returns(tx *types.Transaction)`
| `eth_getTransactionByBlockHashAndIndex`   | `eth.TxByBlockHashAndIndex(blockHash common.Hash, index uint).Returns(tx *types.Transaction)`
| `eth_getTransactionByBlockNumberAndIndex` | `eth.TxByBlockNumberAndIndex(blockNumber *big.Int, index uint).Returns(tx *types.Transaction)`
| `eth_getTransactionCount`                 | `eth.Nonce(addr common.Address, blockNumber *big.Int).Returns(nonce *uint)`
| `eth_getTransactionReceipt`               | `eth.TxReceipt(txHash common.Hash).Returns(receipt *types.Receipt)`
| `eth_sendRawTransaction`                  | `eth.SendRawTx(rawTx []byte).Returns(hash *common.Hash)`<br>`eth.SendTx(tx *types.Transaction).Returns(hash *common.Hash)`
| `eth_getUncleByBlockHashAndIndex`         | `eth.UncleByBlockHashAndIndex(hash common.Hash, index uint).Returns(uncle *types.Header)`
| `eth_getUncleByBlockNumberAndIndex`       | `eth.UncleByBlockNumberAndIndex(number *big.Int, index uint).Returns(uncle *types.Header)`
| `eth_getUncleCountByBlockHash`            | `eth.UncleCountByBlockHash(hash common.Hash).Returns(count *uint)`
| `eth_getUncleCountByBlockNumber`          | `eth.UncleCountByBlockNumber(number *big.Int).Returns(count *uint)`

### [`debug`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo/module/debug)

| Method                   | Go Code
| :----------------------- | :-------
| `debug_traceCall`        | `debug.TraceCall(msg *w3types.Message, blockNumber *big.Int, config *debug.TraceConfig).Returns(trace *debug.Trace)`<br>`debug.CallTraceCall(msg *w3types.Message, blockNumber *big.Int, overrides w3types.State).Returns(trace *debug.CallTrace)`
| `debug_traceTransaction` | `debug.TraceTx(txHash common.Hash, config *debug.TraceConfig).Returns(trace *debug.Trace)`<br>`debug.CallTraceTx(txHash common.Hash, overrides w3types.State).Returns(trace *debug.CallTrace)`

### [`txpool`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo/module/txpool)

| Method               | Go Code
| :--------------------| :-------
| `txpool_content`     | `txpool.Content().Returns(resp *txpool.ContentResponse)`
| `txpool_contentFrom` | `txpool.ContentFrom(addr common.Address).Returns(resp *txpool.ContentFromResponse)`
| `txpool_status`      | `txpool.Status().Returns(resp *txpool.StatusResponse)`

### [`web3`](https://pkg.go.dev/github.com/grassrootseconomics/w3-celo/module/web3)

| Method               | Go Code
| :------------------- | :-------
| `web3_clientVersion` | `web3.ClientVersion().Returns(clientVersion *string)`

### Packages using this library

| Package                                                                  | Description
| :----------------------------------------------------------------------- | :-----------
| [github.com/grassrootseconomics/celoutils](https://github.com/grassrootseconomics/celoutils) | High level Celo utilities specific to Grassroots Economics.
