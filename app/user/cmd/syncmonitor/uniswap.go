package main

import (
	"context"
	"github.com/hasura/go-graphql-client"
)

type UniswapRepo struct {
	graphqlClient *graphql.Client
}

type UniswapCurrentPrice struct {
	Bundles []struct {
		EthPriceUSD graphql.String `graphql:"ethPriceUSD"`
	} `graphql:"bundles(first: 1, subgraphError: allow)"`
}

type UniswapV3DAIBUILDERPool struct {
	Pool struct {
		Token0 struct {
			Id     graphql.String
			Symbol graphql.String
		}
		Token1 struct {
			Id     graphql.String
			Symbol graphql.String
		}
		Token0Price graphql.String `graphql:"token0Price"`
		Token1Price graphql.String `graphql:"token1Price"`
	} `graphql:"pool(id: \"0x2f08654b1482764c7084a2105eb12c3ff50396d2\")"`
}

func NewUniswapRepo() *UniswapRepo {
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3", nil)

	return &UniswapRepo{
		graphqlClient: graphqlClient,
	}
}

func (repo *UniswapRepo) GetEthPrice() (string, error) {
	var query UniswapCurrentPrice
	err := repo.graphqlClient.Query(context.Background(), &query, nil)
	if err != nil {
		return "", err
	}
	return string(query.Bundles[0].EthPriceUSD), nil
}

func (repo *UniswapRepo) GetBuilderPrice() (string, error) {
	var query UniswapV3DAIBUILDERPool
	err := repo.graphqlClient.Query(context.Background(), &query, nil)
	if err != nil {
		return "", err
	}

	return string(query.Pool.Token0Price), nil
}
