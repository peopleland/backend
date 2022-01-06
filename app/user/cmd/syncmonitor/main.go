package main

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data"
	"backend/pkg/env"
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"log"
	"time"
)

type Config struct {
	EthClientRawUrl                  string `mapstructure:"PEOPLELAND_ETH_CLIENT_RAW_URL"`
	PeopleLandContractTheGraphApiUrl string `mapstructure:"PEOPLELAND_CONTRACT_THE_GRAPH_API_URL"`
	GetRoundInfoApiUrl               string `mapstructure:"PEOPLELAND_GET_ROUND_INFO_API_URL"`
	SyncOpenerRecordBackgroundApiUrl string `mapstructure:"PEOPLELAND_SYNC_OPENER_RECORD_BACKGROUND_API_URL"`
}

type GameInfo struct {
	Data struct {
		Info struct {
			RoundNumber        int    `json:"round_number,omitempty"`
			BuilderTokenAmount string `json:"builder_token_amount,omitempty"`
			EthAmount          string `json:"eth_amount,omitempty"`
			StartTimestamp     int64  `json:"start_timestamp,omitempty"`
			EndTimestamp       int64  `json:"end_timestamp,omitempty"`
			HasWinner          bool   `json:"has_winner,omitempty"`
			WinnerTokenId      int64  `json:"winner_token_id,omitempty"`
		} `json:"info,omitempty"`
		OpenerRecord struct {
			MintAddress    string `json:"mint_address,omitempty"`
			TokenId        int64  `json:"token_id,omitempty"`
			X              string `json:"x,omitempty"`
			Y              string `json:"y,omitempty"`
			BlockNumber    int64  `json:"block_number,omitempty"`
			BlockTimestamp int64  `json:"block_timestamp,omitempty"`
		} `json:"opener_record,omitempty"`
	} `json:"data,omitempty"`
}

var logger = log.Default()
var config *Config
var peopleLandContractTheGraphRepo biz.PeopleLandContractTheGraphRepo

func emitSync() {
	request := gorequest.New()
	_, _, _ = request.Get(config.SyncOpenerRecordBackgroundApiUrl).End()
}

func runMonitor(gameInfo *GameInfo) (isEmit bool, err error) {
	/**
	  if newest_opener_record_token_id
	  	token list = 请求 thegraph token list from newest_opener_record_token_id
	  else
	  	token list = 请求 thegraph token list from start timestamp
	  if have token list
	  	emit event to netlify background fun
	  	return
	*/
	var list []*biz.PeopleLandTokenInfo
	var listBlockTimestamp int64
	if gameInfo.Data.OpenerRecord.TokenId != 0 {
		list, listBlockTimestamp, err = peopleLandContractTheGraphRepo.GetTokenInfoListByFromTokenId(gameInfo.Data.OpenerRecord.TokenId + 1)
		if err != nil {
			return false, err
		}
	} else {
		list, listBlockTimestamp, err = peopleLandContractTheGraphRepo.GetTokenInfoListByFromTimestamp(gameInfo.Data.OpenerRecord.BlockTimestamp)
		if err != nil {
			return false, err
		}
	}
	if len(list) > 0 {
		logger.Println("have_news.emit_sync")
		emitSync()
		return true, err
	}
	if gameInfo.Data.OpenerRecord.TokenId != 0 {
		if listBlockTimestamp-gameInfo.Data.OpenerRecord.BlockNumber >= biz.WinnerTimeCon {
			logger.Println("have_winner.emit_sync")
			emitSync()
			return true, err
		}
	}

	logger.Println("no_news")
	return false, err

}

func getGameInfo() (*GameInfo, error) {
	request := gorequest.New()
	_, body, errs := request.Get(config.GetRoundInfoApiUrl).End()
	if errs != nil {
		return nil, errs[0]
	}
	var gameInfo GameInfo
	err := json.Unmarshal([]byte(body), &gameInfo)
	if err != nil {
		return nil, err
	}

	return &gameInfo, nil
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			logger.Println(r)
			time.Sleep(15 * time.Second)
		}
	}()

	/**
	  	通过 get game info api 获取
	  	1. newest_opener_record_token_id
	  	2. start timestamp
	  if 没有获胜
	  	15 秒轮训一次
	*/

	//peopleLandContractTheGraphRepo
	gameInfo, err := getGameInfo()
	if err != nil {
		panic(err)
	}
	if gameInfo.Data.Info.RoundNumber == 0 {
		panic(errors.New("no_info"))
	}
	if gameInfo.Data.Info.HasWinner {
		panic(errors.New("has_winner"))
	}
	for {
		logger.Println("run_monitor.start")
		isEmit, err := runMonitor(gameInfo)
		logger.Println("run_monitor.end")
		if err != nil {
			panic(err)
		}
		if isEmit {
			time.Sleep(15 * time.Second)
			return
		} else {
			time.Sleep(5 * time.Second)
		}
	}

}

func initEnv(c *Config) biz.PeopleLandContractTheGraphRepo {
	conff := &conf.Config{
		PeopleLandContractTheGraphApiUrl: c.PeopleLandContractTheGraphApiUrl,
		EthClientRawUrl:                  c.EthClientRawUrl,
	}
	return data.NewPeopleLandContractTheGraphRepo(conff)
}

func getConfig() *Config {
	var c Config
	e := env.NewEnv()
	_ = e.LoadFile("./config")
	_ = e.Read(&c)

	logger.Println("EthClientRawUrl", c.EthClientRawUrl)
	logger.Println("PeopleLandContractTheGraphApiUrl", c.PeopleLandContractTheGraphApiUrl)
	logger.Println("GetRoundInfoApiUrl", c.GetRoundInfoApiUrl)
	logger.Println("SyncOpenerRecordBackgroundApiUrl", c.SyncOpenerRecordBackgroundApiUrl)

	return &c
}

func main() {
	logger.Println("syncmonitor", 1)
	config = getConfig()
	peopleLandContractTheGraphRepo = initEnv(config)
	for {
		logger.Println("process.start")
		process()
		logger.Println("process.end")
	}
}
