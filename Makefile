run:
	go run ./cmd/server/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/token_checker.proto

test_all:
	go test ./...

format:
	gofmt -s -w .