
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