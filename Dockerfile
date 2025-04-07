# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache \
    git \
    make \
    protoc \
    protobuf-dev

# Install protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /honeyrpc

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate protobuf files
RUN protoc --go_out=. --go-grpc_out=. proto/*.proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /honeyrpc/cmd/server ./cmd/server

# Stage 2: Create minimal runtime image
FROM gcr.io/distroless/static-debian11

# Copy the binary
COPY --from=builder /honeyrpc/cmd/server/server /honeyrpc/
COPY --from=builder /honeyrpc/configs /honeyrpc/configs

WORKDIR /honeyrpc

EXPOSE 50051

ENTRYPOINT ["/honeyrpc/server"]