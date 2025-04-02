run:
	go run cmd/server/main.go

test:
	go test ./test/... -v

build:
	go build -o bin/dns-server cmd/server/main.go