# Solana Token Safety Checker

A simple gRPC-based service to check the safety of Solana tokens.

---

## Setup Project

Install `protoc` and `protobuf`.

If you're on **macOS**, run:

```bash
brew install protobuf
```

Or refer to the official installation guide:  
https://protobuf.dev/installation/

Then generate the gRPC code:

```bash
make proto
```

---

## Run Server Locally

Using Makefile:

```bash
make run
```

Or run directly:

```bash
go run ./cmd/server/main.go
```

---

## Manual Local Testing

### Request Example

```bash
grpcurl -emit-defaults -d '{
  "tokenMint": "FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"
}' -plaintext localhost:50051 tokenchecker.TokenChecker/CheckTokenSafety
```

### Response Example

```json
{
  "isNonStandardProgram": false,
  "hasMintAuthority": true,
  "hasFreezeAuthority": false,
  "reason": "Mint authority still active; "
}
```

---

## Using the Deployed App

**URL:** `melodic-harmonia-rwwwx-6f8204d9.koyeb.app`

### Request Example

```bash
grpcurl -emit-defaults -d '{
  "tokenMint": "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN"
}' melodic-harmonia-rwwwx-6f8204d9.koyeb.app:443 tokenchecker.TokenChecker/CheckTokenSafety
```

### Response Example

```json
{
  "isNonStandardProgram": false,
  "hasMintAuthority": false,
  "hasFreezeAuthority": false,
  "reason": "Token looks safe and using: SPL Token v1"
}
```

> You can also use **Postman**.  
> The server supports reflection, so proto files are not required.

---

## 🔍 How Token Safety Check Works

The verification process is straightforward:

1. Calls an RPC node with a `getAccountInfo` request.
2. Parses the response to extract:
    - Owner
    - Mint authority
    - Freeze authority
3. Verifies that the owner is listed in the `Accepted SPL Programs`.
4. Constructs the response using the **Functional Option** pattern.

---