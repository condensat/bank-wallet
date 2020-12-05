# Bank Wallet

## Setup

### Download and run required docker images

```
docker run --restart=always -p 4222:4222 --name nats-test -d nats:2.1-alpine
docker run --restart=always -p 6379:6379 --name redis-test -d redis:5-alpine
docker run --restart=always -p 3306:3306 --name mariadb-test -e MYSQL_RANDOM_ROOT_PASSWORD=yes -e MYSQL_USER=condensat -e MYSQL_PASSWORD=condensat -e MYSQL_DATABASE=condensat -v $(pwd)/tests/database/permissions.sql:/docker-entrypoint-initdb.d/permissions.sql:ro -d mariadb:10.3
```

### Clone the SSM repository and run the docker compose

```
git clone https://code.condensat.tech/global/rd/crypto-ssm.git
cd crypto-ssm/tor
make start
```

There might be a permission issue with the `ssm-service` dir that will prevent reading it once copied inside the container. 
To solve this, you can `sudo chown -R 100:101 ssm-service` in your host. 

### create a new wallet in the SSM

Get your onion address from the `hostname` file in `ssm-service`.
Get some entropy with `openssl rand -hex 64`.

```
export CHAIN="elements-regtest"
export SSM_ENDPOINT="http://your_address_here.onion/api/v1
export MASTER_ENTROPY="some random 64B"

curl -s --socks5-hostname 0.0.0.0:9050 curl -i -X POST -H "Content-Type: application/json" -d '{
        "jsonrpc": "2.0",
        "method": "new_master",
        "params": ["'"$CHAIN"'", "'"$MASTER_ENTROPY"'", true],
        "id": "42"
    }' $SSM_ENDPOINT

# expected output
HTTP/1.0 200 OK
Content-Type: application/json
Content-Length: 87
Server: Werkzeug/1.0.1 Python/3.8.2
Date: Sat, 05 Dec 2020 11:52:41 GMT

{"id":"42","jsonrpc":"2.0","result":{"chain":"$CHAIN","fingerprint":"some_random_string"}}
```

If it didn't work:
* `docker logs registry.condensat.space/crypto-ssm-tor:$(VERSION)` to see if something's wrong on the tor side
* `docker logs registry.condensat.space/crypto-ssm:$(VERSION)` for SSM.

### Create a config file

Here's an example for Elements-regtest that you can save at the root of bank dir for example:
```
{
  "chains": [
    {
      "chain": "liquid-regtest",
      "hostname": "elements_regtest",
      "port": $PORT,
      "user": $USER,
      "pass": $PASSWORD
    }
  ],
  "ssm": {
    "devices": [
      {
        "name": "crypto-ssm",
        "endpoint": "http://your_onion.onion/api/v1"
      }
    ],
    "chains": [
      {
        "device":            "crypto-ssm",
        "chain":             "elements-regtest",
        "fingerprint":       $FINGERPRINT,
        "derivation_prefix": "0/0"
      }
    ]
  },
    "tor_proxy": "socks5://127.0.0.1:9050"
}
```

The hostname must be recorded in the `/etc/hosts` file.
$FINGERPRINT is the output of the curl command we made above.

### Install and run elements in regtest

## Run the server

### Run bank-wallet

`go run wallet/cmd/bank-wallet/main.go --cryptoMode=crypto-ssm -chains=path/to/conf --log=debug`

## Run the client

### Get a new deposit address

`go run wallet/cmd/wallet-cli/main.go -cmd=getDepositAddress`

### Send some L-BTC to the address

`elements-cli sendtoaddress $ADDRESS $AMOUNT`

### Issue a new asset

`go run wallet/cmd/wallet-cli/main.go -cmd=issueAsset -issuanceMode=with-token-and-contract -assetAmount=1000 -tokenAmount=0.01 -contractHash="random_32B"`

### Reissue an existing asset

`go run wallet/cmd/wallet-cli/main.go -cmd=reissueAsset -assetAmount=1000 -assetID="some random 64B string"`

### List asset issuances

`go run wallet/cmd/wallet-cli/main.go -cmd=listIssuances -assetID="some random 64B string"`

### Burn asset

`go run wallet/cmd/wallet-cli/main.go -cmd=burnAsset -assetID="some random 64B string" -assetAmount=110`

## Known Issues

* Using express VPN can interact with the containers and prevent nats to work properly. If nats timeout, check that there's no vpn connection, turn it off and maybe reboot if this doesn't seem to solve the issue.
* Using ufw can interact with the containers too, check that it's inactive with `systemctl status ufw` and turn it down with `sudo systemctl stop ufw`. Disable it with `sudo systemctl disable ufw` and reboot to make sure everything's allright.
* If the SSM doesn't seem to work (can't generate new addresses), make sure that the fingerprint in the config file is the same than the one returned by `new-master`, and also that the chain is the same.