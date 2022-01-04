package data

import (
	"backend/app/user/internal/conf"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var contractConfig = &conf.Config{
	EthClientRawUrl:           "https://mainnet.infura.io/v3/99a79f80961b4db7aab7c9f54375eda7",
	PeopleLandContractAddress: "0xD1d30B80C5EFB9145782634fe0F9cbeD5D24ef3b",
}

func TestPeopleLandContractRepo_BalanceOf(t *testing.T) {
	repo := NewPeopleLandContractRepo(contractConfig)
	count, _ := repo.BalanceOf("0x1111111111111111111111111111111111111111")
	assert.Equal(t, count.Cmp(big.NewInt(0)), 1)
}
