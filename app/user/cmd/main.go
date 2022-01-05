package main

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data"
	"backend/app/user/internal/server"
	"backend/app/user/internal/service"
	"backend/pkg/env"
	"backend/pkg/http"
	"log"
)

func initApp(conf *conf.Config, logger *log.Logger) (*http.Server, error) {
	d, err := data.NewData(conf, logger)
	if err != nil {
		return nil, err
	}
	userRepo := data.NewUserRepo(d, logger)
	twitterRepo := data.NewTwitterRepo(conf)
	discordRepo := data.NewDiscordRepo(conf)
	peopleLandContractRepo := data.NewPeopleLandContractRepo(conf)
	mintRecordRepo := data.NewMintRecordRepo(d, logger)
	openerRecordRepo := data.NewOpenerRecordRepo(d, logger)
	openerGameRoundInfoRepo := data.NewOpenerGameRoundInfoRepo(d, logger)
	peopleLandContractTheGraphRepo := data.NewPeopleLandContractTheGraphRepo(conf)

	userUseCase := biz.NewUserUseCase(userRepo, twitterRepo, discordRepo, peopleLandContractRepo, conf, logger)
	mintRecordCase := biz.NewOpenerGameCase(userRepo, mintRecordRepo, openerRecordRepo, openerGameRoundInfoRepo, peopleLandContractTheGraphRepo, peopleLandContractRepo, conf, logger)

	userService := service.NewUserService(userUseCase, mintRecordCase, conf, logger)

	lambdaServer := server.NewLambdaServer(conf, userService, logger)
	return lambdaServer, nil
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
	app, err := initApp(&config, logger)
	if err != nil {
		panic(err)
	}

	if config.Env == "local_dev" {
		app.HttpStart(":8081")
	} else {
		app.LambdaStart()
	}
}
