package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"context"
	"github.com/hasura/go-graphql-client"
)

type PeopleLandContractTheGraphRepo struct {
	graphqlClient *graphql.Client
}

type PeopleLandContractTheGraphQueryByFromTokenId struct {
	TokenInfos []struct {
		TokenId graphql.Int
		X       graphql.Int
		Y       graphql.Int
		Minted  struct {
			Id graphql.String
		}
		Owner struct {
			Id graphql.String
		}
		GivedAtTimestamp   graphql.Int
		GivedAtBlockNumber graphql.Int
	} `graphql:"tokenInfos(first: 100, where: {tokenId_gte: $tokenId_gte}, orderBy: tokenId)"`
}

type PeopleLandContractTheGraphQueryByFromTimestamp struct {
	TokenInfos []struct {
		TokenId graphql.Int
		X       graphql.String
		Y       graphql.String
		Minted  struct {
			Id graphql.String
		}
		Owner struct {
			Id graphql.String
		}
		GivedAtTimestamp   graphql.Int
		GivedAtBlockNumber graphql.Int
	} `graphql:"tokenInfos(first: 100, where: {givedAtTimestamp_gte: $givedAtTimestamp_gte}, orderBy: tokenId)"`
}

func NewPeopleLandContractTheGraphRepo(config *conf.Config) biz.PeopleLandContractTheGraphRepo {
	client := graphql.NewClient(config.PeopleLandContractTheGraphApiUrl, nil)
	return &PeopleLandContractTheGraphRepo{
		graphqlClient: client,
	}
}

func (repo *PeopleLandContractTheGraphRepo) GetTokenInfoListByFromTokenId(fromTokenId int64) ([]*biz.PeopleLandTokenInfo, error) {
	var query PeopleLandContractTheGraphQueryByFromTokenId
	variables := map[string]interface{}{
		"tokenId_gte": fromTokenId,
	}
	err := repo.graphqlClient.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	list := make([]*biz.PeopleLandTokenInfo, 0)
	for _, item := range query.TokenInfos {
		list = append(list, &biz.PeopleLandTokenInfo{
			TokenId:            int64(item.TokenId),
			X:                  string(item.X),
			Y:                  string(item.Y),
			MintedAddress:      string(item.Minted.Id),
			OwnerAddress:       string(item.Owner.Id),
			GivedAtTimestamp:   int64(item.GivedAtTimestamp),
			GivedAtBlockNumber: int64(item.GivedAtBlockNumber),
		})
	}
	return list, nil
}

func (repo *PeopleLandContractTheGraphRepo) GetTokenInfoListByFromTimestamp(fromTimestamp int64) ([]*biz.PeopleLandTokenInfo, error) {
	var query PeopleLandContractTheGraphQueryByFromTimestamp
	variables := map[string]interface{}{
		"givedAtTimestamp_gte": fromTimestamp,
	}
	err := repo.graphqlClient.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	list := make([]*biz.PeopleLandTokenInfo, 0)
	for _, item := range query.TokenInfos {
		list = append(list, &biz.PeopleLandTokenInfo{
			TokenId:            int64(item.TokenId),
			X:                  string(item.X),
			Y:                  string(item.Y),
			MintedAddress:      string(item.Minted.Id),
			OwnerAddress:       string(item.Owner.Id),
			GivedAtTimestamp:   int64(item.GivedAtTimestamp),
			GivedAtBlockNumber: int64(item.GivedAtBlockNumber),
		})
	}
	return list, nil
}
