package solana

import (
	"context"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/samber/mo"
	"honeyrpc/internal/entity"
)

type Client struct {
	rpcClient *rpc.Client
}

func NewClient(rpcUrl string) *Client {
	return &Client{
		rpcClient: rpc.New(rpcUrl),
	}
}

func (client *Client) GetMintAccountInfo(ctx context.Context, mintAddress solana.PublicKey) (*entity.MintAccountInfo, error) {
	resp, err := client.rpcClient.GetAccountInfo(ctx, mintAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get account info: %w", err)
	}

	tokenAccountOwner := resp.Value.Owner

	var mintAccount token.Mint
	if err := bin.NewBinDecoder(resp.Value.Data.GetBinary()).Decode(&mintAccount); err != nil {
		return nil, fmt.Errorf("failed to decode mint account: %w", err)
	}

	var mintAuthority mo.Option[solana.PublicKey]

	if mintAccount.MintAuthority == nil {
		mintAuthority = mo.None[solana.PublicKey]()
	} else {
		mintAuthority = mo.Some(*mintAccount.MintAuthority)
	}

	var freezeAuthority mo.Option[solana.PublicKey]

	if mintAccount.FreezeAuthority == nil {
		freezeAuthority = mo.None[solana.PublicKey]()
	} else {
		freezeAuthority = mo.Some(*mintAccount.FreezeAuthority)
	}

	return &entity.MintAccountInfo{
		OwnerProgram:    tokenAccountOwner,
		MintAuthority:   mintAuthority,
		Supply:          mintAccount.Supply,
		Decimals:        mintAccount.Decimals,
		IsInitialized:   mintAccount.IsInitialized,
		FreezeAuthority: freezeAuthority,
	}, nil
}
