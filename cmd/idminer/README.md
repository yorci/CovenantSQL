idminer is a helper command of CovenantSQL

## Install 

Make sure that `$GOPATH/bin` is in your `$PATH`

```bash
$ go get github.com/CovenantSQL/CovenantSQL/cmd/idminer
```

Show usage of `idminer`:

```bash
$ idminer -help
```

## Usage
### Generate Key Pair

```
$ idminer -tool keygen
Enter master key(press Enter for default: ""): 
⏎
Private key file: private.key
Public key's hex: 03bb49a9d4cc6fbfddb520924dd0603c74c31f15fb31641e34bc03a8fe3e0d3f5b
```

The private.key is your encrypted private key file, and the pubkey hex is your public key's hex.

### Generate Wallet Address from Existing Key

```
$ idminer -tool addrgen -private private.key
Enter master key(default: ""):
⏎
wallet address: 4kcD5jFFqLMmiG31QTFWbYkKHFhVgpSTejeZkQHjztaMVcnk2CE
$ idminer -tool addrgen -public 03bb49a9d4cc6fbfddb520924dd0603c74c31f15fb31641e34bc03a8fe3e0d3f5b
wallet address: 4kcD5jFFqLMmiG31QTFWbYkKHFhVgpSTejeZkQHjztaMVcnk2CE
```

You can generate your *wallet* address for test net according to your private key or public key.
