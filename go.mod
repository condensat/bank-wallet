module github.com/condensat/bank-wallet

go 1.14

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/condensat/bank-core v0.0.2-0.20200504100000-a9feaf562fb0
	github.com/google/uuid v1.1.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/shengdoushi/base58 v1.0.0
	github.com/sirupsen/logrus v1.7.0
	github.com/ybbus/jsonrpc v2.1.2+incompatible
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/btcsuite/btcd => github.com/condensat/btcd v0.20.1-beta.0.20200424100000-5dc523e373e2
