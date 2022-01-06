package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/peopleland"
	"context"
	"github.com/shopspring/decimal"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PeopleLandContractRepo struct {
	client          *ethclient.Client
	contractAddress common.Address
}

func NewPeopleLandContractRepo(config *conf.Config) biz.PeopleLandContractRepo {
	client, err := ethclient.Dial(config.EthClientRawUrl)
	if err != nil {
		return nil
	}

	return &PeopleLandContractRepo{
		client:          client,
		contractAddress: common.HexToAddress(config.PeopleLandContractAddress),
	}
}

func (repo *PeopleLandContractRepo) BalanceOf(address string) (*big.Int, error) {
	instance, err := peopleland.NewPeopleland(repo.contractAddress, repo.client)
	if err != nil {
		return nil, err
	}

	count, err := instance.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (repo *PeopleLandContractRepo) GetEthBalanceAt(blockNumber int64) (string, error) {
	number := big.NewInt(blockNumber)
	if blockNumber <= 0 {
		number = nil
	}
	wei, err := repo.client.BalanceAt(context.Background(), repo.contractAddress, number)
	if err != nil {
		return "", err
	}

	weiDecimal, err := decimal.NewFromString(wei.String())
	if err != nil {
		return "", err
	}
	weiUnitDecimal, err := decimal.NewFromString("1000000000000000000")
	if err != nil {
		return "", err
	}

	return weiDecimal.Div(weiUnitDecimal).String(), nil
}
