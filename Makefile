build:
	goflags=-mod=mod go build -o bin/go-faucet cmd/main.go
run:
	goflags=-mod=mod go run cmd/main.go
exec:
	./bin/go-faucet
