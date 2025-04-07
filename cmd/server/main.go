package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"honeyrpc/internal/config"
	"honeyrpc/internal/handler"
	"honeyrpc/internal/service"
	"honeyrpc/internal/solana"
	"honeyrpc/pkg/token_checker"
	"log"
	"net"
)

func main() {
	appConfig := config.LoadConfig("./configs/default.json")
	grpcPort := ":" + appConfig.Port

	lis, err := net.Listen(appConfig.NetworkType, grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	solanaClient := solana.NewClient(appConfig.SolanaRpcNodeUrl)
	tokenCheckerService := service.NewTokenCheckerService(solanaClient)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	token_checker.RegisterTokenCheckerServer(grpcServer, handler.NewGRPCServer(tokenCheckerService))

	log.Println("Server listening on ", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
