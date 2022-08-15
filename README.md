# gnark_cubic

## requirement

go 1.17

## run

```shell
go run main.go
```

## test solidity verifying contract
```shell
cd solidity
go run contract/main.go
docker run --rm -v $(pwd):/root ethereum/solc:0.8.16-alpine -o /root --overwrite --abi --bin /root/contract.sol 
docker run --rm -v $(pwd):/root ethereum/client-go:alltools-v1.9.25 abigen --bin /root/Verifier.bin  --abi /root/Verifier.abi --pkg solidity --type Verifier --out /root/solidity.go
go test -v
```
