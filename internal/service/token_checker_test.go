package service

import (
	"context"
	"honeyrpc/internal/solana"
	"honeyrpc/internal/solana/mock"
	"testing"
)

const mainNetUrl = "https://api.mainnet-beta.solana.com"

func TestTokenVerification(t *testing.T) {
	solanaClient := solana.NewClient(mainNetUrl)
	tockenService := NewTokenCheckerService(solanaClient)

	mint, err := tockenService.ParseMintAddress("SecreSFXKJAfDBr1fnguPomXTxCMsFuXiXZGR57JRhA")
	if err != nil {
		t.Error(err)
	}

	res, err := tockenService.CheckToken(context.Background(), mint)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestNegativeTokenVerification(t *testing.T) {
	solanaClient := &mock.NegativeCaseMockClient{}
	tockenService := NewTokenCheckerService(solanaClient)

	mint, err := tockenService.ParseMintAddress("SecreSFXKJAfDBr1fnguPomXTxCMsFuXiXZGR57JRhA")
	if err != nil {
		t.Error(err)
	}

	res, err := tockenService.CheckToken(context.Background(), mint)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)

	if !res.HasMintAuthority {
		t.Error("Should have mint authority")
	}

	if !res.HasFreezeAuthority {
		t.Error("Should have freeze authority")
	}

	if !res.IsNonStandardProgram {
		t.Error("Should be non standard program")
	}

}
