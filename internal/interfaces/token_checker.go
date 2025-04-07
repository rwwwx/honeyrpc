package interfaces

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"honeyrpc/internal/entity"
)

type TokenChecker interface {
	ParseMintAddress(mintAddress string) (solana.PublicKey, error)
	CheckToken(ctx context.Context, mintAddress solana.PublicKey) (*entity.TokenSafetyResult, error)
}
