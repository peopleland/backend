package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/peopleland"
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
