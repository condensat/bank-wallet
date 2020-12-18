# Lightning cli

## Usage

### Common flags

All flags are mandatory

```bash
  --endpoint string
    	Lightning node onion address
  --macaroon string
    	macaroon in hex
  --torproxy string
    	Tor Proxy for hidden services access
```

Common flags can be loaded from `.env` file

```bash
LN_CLI_TORPROXY=socks5://127.0.0.1:9050

LN_CLI_ENDPOINT=http://rtl.onion

LN_CLI_MACAROON=<hex_macaroon>
```

### GetInfo flags

```bash
Usage of getInfo:
  No flags.
```

### KeySend flags

```bash
Usage of keySend:
  --amount int
    	KeySend amount (sat)
  --label string
    	KeySend label
  --pubKey string
    	KeySend pubKey
```

### Pay flags

```bash
Usage of pay:
  -invoice string
    	Pay invoice
```

### DecodePay flags

```bash
Usage of decodepay:
  -invoice string
    	Invoice to decode (Bolt	11)
```
