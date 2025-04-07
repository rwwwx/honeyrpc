## Setup project
Install `protoc` and `protobuf`.
If you are running macOS use:
```shell
brew install protobuf
```
or refer to this page https://protobuf.dev/installation/

Then run:
```shell
make proto
```

## Run server locally
Using make file:
```shell
make run
```
or: 
```shell
go run ./cmd/server/main.go
```

## Local manual testing
Request example:
```shell
grpcurl -emit-defaults -d '{
  "tokenMint": "FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"
}' -plaintext localhost:50051 tokenchecker.TokenChecker/CheckTokenSafety
```
Response example:
```json
{
  "isNonStandardProgram": false,
  "hasMintAuthority": true,
  "hasFreezeAuthority": false,
  "reason": "Mint authority still active; "
}
```

## Explanation of how program IDs are checked
The token verification process is very simple:   
1. Call RPC node with `GetAccountInfo` request
2. Parse response for owner, freeze and mint authority
3. Verify that the owner is described in `Accepted SPL Programs` list
4. Create server response via `Functional Option` pattern 
