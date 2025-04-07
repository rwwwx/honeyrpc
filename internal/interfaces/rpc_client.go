package interfaces

import (
	"context"
	solanago "github.com/gagliardetto/solana-go"
	"honeyrpc/internal/entity"
)

type RpcClient interface {
	GetMintAccountInfo(ctx context.Context, mintAddress solanago.PublicKey) (*entity.MintAccountInfo, error)
}
