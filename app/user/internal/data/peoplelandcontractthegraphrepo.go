package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hasura/go-graphql-client"
	"math/big"
	"strconv"
	"strings"
)

type PeopleLandContractTheGraphRepo struct {
	graphqlClient *graphql.Client
	ethClient     *ethclient.Client
}

type PeopleLandContractTheGraphTokenInfo struct {
	TokenId graphql.String
	X       graphql.String
	Y       graphql.String
	Minted  struct {
		Id graphql.String
	}
	Owner struct {
		Id graphql.String
	}
	GivedAtTimestamp   graphql.String
	GivedAtBlockNumber graphql.String
}

type PeopleLandContractTheGraphMeta struct {
	Block struct {
		Number graphql.Int
	}
}

type PeopleLandContractTheGraphQueryByFromTokenId struct {
	TokenInfos []PeopleLandContractTheGraphTokenInfo `graphql:"tokenInfos(first: 40, where: {tokenId_gte: $tokenId_gte}, orderBy: tokenId)"`
	Meta       PeopleLandContractTheGraphMeta        `graphql:"_meta"`
}

type PeopleLandContractTheGraphQueryByFromTimestamp struct {
	TokenInfos []PeopleLandContractTheGraphTokenInfo `graphql:"tokenInfos(first: 40, where: {givedAtTimestamp_gte: $givedAtTimestamp_gte}, orderBy: tokenId)"`
	Meta       PeopleLandContractTheGraphMeta        `graphql:"_meta"`
}

func NewPeopleLandContractTheGraphRepo(config *conf.Config) biz.PeopleLandContractTheGraphRepo {
	graphqlClient := graphql.NewClient(config.PeopleLandContractTheGraphApiUrl, nil)

	ethClient, err := ethclient.Dial(config.EthClientRawUrl)
	if err != nil {
		return nil
	}

	return &PeopleLandContractTheGraphRepo{
		graphqlClient: graphqlClient,
		ethClient:     ethClient,
	}
}

func (repo *PeopleLandContractTheGraphRepo) GetTokenInfoListByFromTokenId(fromTokenId int64) ([]*biz.PeopleLandTokenInfo, int64, error) {
	var query PeopleLandContractTheGraphQueryByFromTokenId
	variables := map[string]interface{}{
		"tokenId_gte": strconv.FormatInt(fromTokenId, 10),
	}
	err := repo.graphqlClient.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, 0, err
	}
	list := make([]*biz.PeopleLandTokenInfo, 0)
	for _, item := range query.TokenInfos {
		info, err := repo.convertType(item)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, info)
	}
	timestamp, err := repo.getBlockTimestamp(int32(query.Meta.Block.Number))
	if err != nil {
		return nil, 0, err
	}
	return list, timestamp, nil
}

func (repo *PeopleLandContractTheGraphRepo) GetTokenInfoListByFromTimestamp(fromTimestamp int64) ([]*biz.PeopleLandTokenInfo, int64, error) {
	var query PeopleLandContractTheGraphQueryByFromTimestamp
	variables := map[string]interface{}{
		"givedAtTimestamp_gte": strconv.FormatInt(fromTimestamp, 10),
	}
	err := repo.graphqlClient.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, 0, err
	}
	list := make([]*biz.PeopleLandTokenInfo, 0)
	for _, item := range query.TokenInfos {
		info, err := repo.convertType(item)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, info)
	}

	timestamp, err := repo.getBlockTimestamp(int32(query.Meta.Block.Number))
	if err != nil {
		return nil, 0, err
	}
	return list, timestamp, nil
}

func (repo *PeopleLandContractTheGraphRepo) convertType(info PeopleLandContractTheGraphTokenInfo) (*biz.PeopleLandTokenInfo, error) {
	tokenId, err := strconv.ParseInt(string(info.TokenId), 10, 64)
	if err != nil {
		return nil, err
	}
	givedAtTimestamp, err := strconv.ParseInt(string(info.GivedAtTimestamp), 10, 64)
	if err != nil {
		return nil, err
	}
	givedAtBlockNumber, err := strconv.ParseInt(string(info.GivedAtBlockNumber), 10, 64)
	if err != nil {
		return nil, err
	}
	return &biz.PeopleLandTokenInfo{
		TokenId:            tokenId,
		X:                  string(info.X),
		Y:                  string(info.Y),
		MintedAddress:      strings.ToLower(string(info.Minted.Id)),
		OwnerAddress:       strings.ToLower(string(info.Owner.Id)),
		GivedAtTimestamp:   givedAtTimestamp,
		GivedAtBlockNumber: givedAtBlockNumber,
	}, nil
}

func (repo *PeopleLandContractTheGraphRepo) getBlockTimestamp(blockNumber int32) (int64, error) {
	block, err := repo.ethClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		return 0, err
	}
	return int64(block.Time()), nil
}
