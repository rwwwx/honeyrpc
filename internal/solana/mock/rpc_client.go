package mock

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/samber/mo"
	"honeyrpc/internal/entity"
)

type NegativeCaseMockClient struct{}

func (m *NegativeCaseMockClient) GetMintAccountInfo(_ context.Context, _ solana.PublicKey) (*entity.MintAccountInfo, error) {
	return &entity.MintAccountInfo{
		OwnerProgram:    solana.PublicKey{},
		MintAuthority:   mo.Some(solana.PublicKey{}),
		Supply:          111_111111,
		Decimals:        6,
		IsInitialized:   true,
		FreezeAuthority: mo.Some(solana.PublicKey{}),
	}, nil
}
