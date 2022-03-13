build:
	@echo "building the crypto-generator binary"
	go build -o bin/crypto-generator
	@echo "building the crypto-reader binary"
	go build -o bin/crypto-reader

test:
	@echo "running all tests"
	go test -v ./...

bench:
	@echo "running all benchmarks"
	go test -bench . ./...
