package executor

import (
	"github.com/portto/solana-go-sdk/common"
)

type InitNebulaContractInstruction struct {
	Instruction              uint8
	Bft                      uint8
	NebulaDataType           uint8
	GravityContractProgramID common.PublicKey
	Consuls                  []byte
}

type UpdateOraclesNebulaContractInstruction struct {
	Instruction              uint8
	Bft                      uint8
	Oracles                  []byte
	NewRound                 uint64
}

