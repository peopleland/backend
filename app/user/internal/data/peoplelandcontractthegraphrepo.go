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

type PeopleLandContractTheGraphQuery struct {
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

func NewPeopleLandContractTheGraphRepo(config *conf.Config) biz.PeopleLandContractTheGraphRepo {
	client := graphql.NewClient(config.PeopleLandContractTheGraphApiUrl, nil)
	return &PeopleLandContractTheGraphRepo{
		graphqlClient: client,
	}
}

func (repo *PeopleLandContractTheGraphRepo) GetTokenInfoList(fromTokenId int32) ([]*biz.PeopleLandTokenInfo, error) {
	var query PeopleLandContractTheGraphQuery
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
			TokenId:            int32(item.TokenId),
			X:                  int32(item.X),
			Y:                  int32(item.Y),
			MintedAddress:      string(item.Minted.Id),
			OwnerAddress:       string(item.Owner.Id),
			GivedAtTimestamp:   int32(item.GivedAtTimestamp),
			GivedAtBlockNumber: int32(item.GivedAtBlockNumber),
		})
	}
	return list, nil
}
