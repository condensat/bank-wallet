# Bitcoin / Elements RPC wallet

## Setup

### Start bitcoin and elements regtest nodes

### Make sure that you have some fund available

### Configure environment variables

1. Bitcoin
`export BITCOIN_TESTNET_HOSTNAME="your_hostname"`  (Default `bitcoin_testnet`)
`export BITCOIN_TESTNET_PORT=port_configured_in_bitcoin_conf`  (Default `18332`)
`export BITCOIN_TESTNET_USER="rpc_user_in_conf"` (Default `bank-wallet`)
`export BITCOIN_TESTNET_PASSWORD="rpc_password_in_conf"` (Default `password1`)

2. Elements 
`export ELEMENTS_TESTNET_HOSTNAME="your_hostname"`  (Default `elements_regtest`)
`export ELEMENTS_TESTNET_PORT=port_configured_in_elements_conf`  (Default `28432`)
`export ELEMENTS_TESTNET_USER="rpc_user_in_conf"` (Default `bank-wallet`)
`export ELEMENTS_TESTNET_PASSWORD="rpc_password_in_conf"` (Default `password1`)

## Test a simple Bitcoin transaction

`go run wallet/bitcoin/cmd/cli/main.go -command=RawTransactionBitcoin`

## Test a simple Elements transaction

`go run wallet/bitcoin/cmd/cli/main.go -command=RawTransactionElements`

## Test asset issuance

`go run wallet/bitcoin/cmd/cli/main.go -command=AssetIssuance -dest=$DESTADDR -change=$CHANGEADDR -asset=$ASSETADDR -token=$TOKENADDR`

## Test asset reissuance

`go run wallet/bitcoin/cmd/cli/main.go -command=Reissuance -change=$CHANGEADDR -asset=$ASSETADDR -token=$TOKENADDR -reissue=$ASSET`

## Test asset burn asset

`go run wallet/bitcoin/cmd/cli/main.go -command=BurnAsset -burnAsset=$ASSET -burnAmount=$AMOUNT -dest=$DESTADDR -change=$CHANGEADDR`
