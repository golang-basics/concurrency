# Crypto Reader

Blazing fast crypto transactions reader capable of working with giant log files (gigabytes) without too much spin.
Give it a try 🚀

### Prerequisites

- Make sure you have installed Go version >= `1.17`

### Build

```shell
# compiles and generates binaries for crypto-reader and crypto-generator inside the ./bin directory
make build
``` 

### Run

```shell
# run the program directory without generating any binary
go run cmd/crypto-generator/main.go -dir <path/to/dir/testdata> -interval <interval_between_transactions> rotation <interval_rotation_between_transaction_file> total <total_lifetime_of_transactions>
go run cmd/crypto-reader/main.go -directory <path/to/log/files> -address <crypto_address_to_look_for> -limit <limit_number_of_transactions_output> -interval <interval_of_transactions_to_look_for>
make build
# generate crypto transaction files using the binary defaults
./bin/crypto-generator
# display maximum 10 crypto transactions from testdata directory that happened in the last 5 minutes
# for the address 0xd1ABA973674601DD10FEF7Abb239E4e975E26a44
./bin/crypto-reader -directory ./testdata -address 0xd1ABA973674601DD10FEF7Abb239E4e975E26a44 -interval 5m -limit 10
```

### Test

```shell
# runs all the tests present in test files
make test
# generate testdata for the benchmark first
go run cmd/crypto-generator/main.go -dir crypto/bench
# runs all the benchmarks present in test files
make bench
```
