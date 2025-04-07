package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"honeyrpc/internal/interfaces"
	"honeyrpc/pkg/token_checker"
)

var (
	invalidMintAddress = status.Errorf(codes.InvalidArgument, "Mint address is invalid")
	failedToFindToken  = status.Errorf(codes.NotFound, "Filed to find token")
)

type GRPCServer struct {
	token_checker.UnimplementedTokenCheckerServer
	tokenCheckerService interfaces.TokenChecker
}

func NewGRPCServer(tokenCheckerService interfaces.TokenChecker) *GRPCServer {
	return &GRPCServer{tokenCheckerService: tokenCheckerService}
}

func (server *GRPCServer) CheckTokenSafety(ctx context.Context, req *token_checker.TokenSafetyRequest) (*token_checker.TokenSafetyResponse, error) {
	mintAddress, err := server.tokenCheckerService.ParseMintAddress(req.GetTokenMint())
	if err != nil {
		return nil, invalidMintAddress
	}

	tokenSafetyResult, err := server.tokenCheckerService.CheckToken(ctx, mintAddress)
	if err != nil {
		return nil, failedToFindToken
	}

	return &token_checker.TokenSafetyResponse{
		IsNonStandardProgram: tokenSafetyResult.IsNonStandardProgram,
		HasMintAuthority:     tokenSafetyResult.HasMintAuthority,
		HasFreezeAuthority:   tokenSafetyResult.HasFreezeAuthority,
		Reason:               tokenSafetyResult.Reason,
	}, nil
}
