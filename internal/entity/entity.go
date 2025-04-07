package entity

import (
	"github.com/gagliardetto/solana-go"
	"github.com/samber/mo"
)

type TokenSafetyResult struct {
	ProgramName          string
	IsNonStandardProgram bool
	HasMintAuthority     bool
	HasFreezeAuthority   bool
	Reason               string
}

type MintAccountInfo struct {
	OwnerProgram    solana.PublicKey
	MintAuthority   mo.Option[solana.PublicKey]
	Supply          uint64
	Decimals        uint8
	IsInitialized   bool
	FreezeAuthority mo.Option[solana.PublicKey]
}
