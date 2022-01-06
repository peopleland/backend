package main

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data"
	"backend/pkg/env"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

var gameCase *biz.OpenerGameCase

func initApp(conf *conf.Config, logger *log.Logger) {
	d, err := data.NewData(conf, logger)
	if err != nil {
		panic(err)
	}
	userRepo := data.NewUserRepo(d, logger)

	peopleLandContractRepo := data.NewPeopleLandContractRepo(conf)
	mintRecordRepo := data.NewMintRecordRepo(d, logger)
	openerRecordRepo := data.NewOpenerRecordRepo(d, logger)
	openerGameRoundInfoRepo := data.NewOpenerGameRoundInfoRepo(d, logger)
	peopleLandContractTheGraphRepo := data.NewPeopleLandContractTheGraphRepo(conf)

	gameCase = biz.NewOpenerGameCase(userRepo, mintRecordRepo, openerRecordRepo, openerGameRoundInfoRepo, peopleLandContractTheGraphRepo, peopleLandContractRepo, conf, logger)
}

func handler(ctx context.Context, _ events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	gameCase.SyncOpenerRecord(ctx)
	gameCase.SyncRoundInfoEth(ctx)
	return nil, nil
}

func main() {
	var config conf.Config

	e := env.NewEnv()
	e.LoadEnvWithReplace("PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM", "\\n", "\n")
	e.LoadEnvWithReplace("PEOPLELAND_JWT_RSA_PUBLIC_KEY_PEM", "\\n", "\n")
	e.LoadEnv("PEOPLELAND_FAUNADB_SECRET")

	e.LoadEnv("PEOPLELAND_TWITTER_CONSUMER_KEY")
	e.LoadEnv("PEOPLELAND_TWITTER_CONSUMER_SECRET")
	e.LoadEnv("PEOPLELAND_TWITTER_TOKEN")
	e.LoadEnv("PEOPLELAND_TWITTER_TOKEN_SECRET")

	e.LoadEnv("PEOPLELAND_DISCORD_BOT_CLIENT_ID")
	e.LoadEnv("PEOPLELAND_DISCORD_BOT_CLIENT_SECRET")

	e.LoadEnv("PEOPLELAND_ETH_CLIENT_RAW_URL")
	e.LoadEnv("PEOPLELAND_CONTRACT_ADDRESS")

	_ = e.LoadFile("./app/user/configs")
	_ = e.Read(&config)

	logger := log.Default()
	initApp(&config, logger)
	lambda.Start(handler)
}
