package service

import (
	"context"
	"fmt"
	solanago "github.com/gagliardetto/solana-go"
	"honeyrpc/internal/entity"
	"honeyrpc/internal/interfaces"
	"strings"
)

type TokenCheckerService struct {
	solanaRpcClient interfaces.RpcClient
}

func NewTokenCheckerService(solanaClient interfaces.RpcClient) *TokenCheckerService {
	return &TokenCheckerService{solanaRpcClient: solanaClient}
}

func (service *TokenCheckerService) CheckToken(ctx context.Context, mintAddress solanago.PublicKey) (*entity.TokenSafetyResult, error) {
	mintAccountInfo, err := service.solanaRpcClient.GetMintAccountInfo(ctx, mintAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get mint account info: %w", err)
	}

	return createTokenSafetyResult(
		isStandardProgramUsed(mintAccountInfo),
		hasMintAuthority(mintAccountInfo),
		hasFreezeAuthority(mintAccountInfo),
	), nil
}

func (_ *TokenCheckerService) ParseMintAddress(mintAddress string) (solanago.PublicKey, error) {
	return solanago.PublicKeyFromBase58(mintAddress)
}

var acceptedPrograms = map[string]string{
	"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA": "SPL Token v1",
	"TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb": "SPL Token v2 (Token2022)",
}

const (
	unknownProgramIdErr = "Token uses unknown program ID: "
	freezeAuthorityErr  = "Freeze authority still active"
	mintAuthorityErr    = "Mint authority still active"
	reasonSeparator     = "; "
	tokenIsSafe         = "Token looks safe and using: "
)

type tokenSafetyResultOption func(*entity.TokenSafetyResult)

func createTokenSafetyResult(opts ...tokenSafetyResultOption) *entity.TokenSafetyResult {
	res := entity.TokenSafetyResult{}

	for _, opt := range opts {
		opt(&res)
	}

	if len(res.Reason) == 0 {
		res.Reason = tokenIsSafe + res.ProgramName
	}

	res.Reason = strings.TrimSpace(res.Reason)

	return &res
}

func foldReasons(newReason, previousReasons string) string {
	return newReason + reasonSeparator + previousReasons
}

func isStandardProgramUsed(info *entity.MintAccountInfo) tokenSafetyResultOption {
	return func(result *entity.TokenSafetyResult) {
		tokenProgramAddress := info.OwnerProgram.String()
		tokenProgramName := acceptedPrograms[tokenProgramAddress]
		result.IsNonStandardProgram = tokenProgramName == ""
		if result.IsNonStandardProgram {
			result.Reason = foldReasons(unknownProgramIdErr+tokenProgramAddress, result.Reason)
		} else {
			result.ProgramName = tokenProgramName
		}
	}
}

func hasFreezeAuthority(info *entity.MintAccountInfo) tokenSafetyResultOption {
	return func(result *entity.TokenSafetyResult) {
		result.HasFreezeAuthority = info.FreezeAuthority.IsPresent()
		if result.HasFreezeAuthority {
			result.Reason = foldReasons(freezeAuthorityErr, result.Reason)
		}
	}
}

func hasMintAuthority(info *entity.MintAccountInfo) tokenSafetyResultOption {
	return func(result *entity.TokenSafetyResult) {
		result.HasMintAuthority = info.MintAuthority.IsPresent()
		if result.HasMintAuthority {
			result.Reason = foldReasons(mintAuthorityErr, result.Reason)
		}
	}
}
