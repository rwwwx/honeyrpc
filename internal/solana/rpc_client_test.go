package solana

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"log"
	"testing"
)

const mainNetUrl = "https://api.mainnet-beta.solana.com"

func TestGetMintAccountInfoForTokenProgram(t *testing.T) {
	client := NewClient(mainNetUrl)

	tokenAddress := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt")
	res, err := client.GetMintAccountInfo(context.Background(), tokenAddress)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(res)
}

func TestGetMintAccountInfoForToken2022Program(t *testing.T) {
	client := NewClient(mainNetUrl)

	tokenAddress := solana.MustPublicKeyFromBase58("SecreSFXKJAfDBr1fnguPomXTxCMsFuXiXZGR57JRhA")
	res, err := client.GetMintAccountInfo(context.Background(), tokenAddress)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(res)
}
