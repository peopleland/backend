package main

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data"
	"backend/pkg/env"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"time"
)

type Config struct {
	EthClientRawUrl                  string `mapstructure:"PEOPLELAND_ETH_CLIENT_RAW_URL"`
	PeopleLandContractTheGraphApiUrl string `mapstructure:"PEOPLELAND_CONTRACT_THE_GRAPH_API_URL"`
	GetRoundInfoApiUrl               string `mapstructure:"PEOPLELAND_GET_ROUND_INFO_API_URL"`
	SyncOpenerRecordBackgroundApiUrl string `mapstructure:"PEOPLELAND_SYNC_OPENER_RECORD_BACKGROUND_API_URL"`

	DiscordBotToken string `mapstructure:"PEOPLELAND_DISCORD_BOT_TOKEN"`
}

type OpenerRecord struct {
	MintAddress    string `json:"mint_address,omitempty"`
	TokenId        int64  `json:"token_id,omitempty"`
	X              string `json:"x,omitempty"`
	Y              string `json:"y,omitempty"`
	BlockNumber    int64  `json:"block_number,omitempty"`
	BlockTimestamp int64  `json:"block_timestamp,omitempty"`
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
		OpenerRecord OpenerRecord `json:"opener_record,omitempty"`
	} `json:"data,omitempty"`
}

type LiveMessageHistory struct {
	TokenId     int64
	MintAddress string
	Source      time.Duration
}

var logger = log.Default()
var config *Config
var peopleLandContractTheGraphRepo biz.PeopleLandContractTheGraphRepo
var discordRepo biz.DiscordRepo

var lastLiveMessageInfo *LiveMessageHistory

var currentOpenerTokenId int64

const ChannelId = "929277955203547136"

func durationToString(dur time.Duration) string {
	hour := int64(dur.Hours())
	minute := int64(dur.Minutes()) - hour*60
	second := int64(dur.Seconds()) - hour*60*60 - minute*60
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func sendOpenerChangeToDiscord(info *OpenerRecord) {
	timestamp := info.BlockTimestamp
	dur := time.Unix(timestamp, 0).Add(24 * time.Hour).Sub(time.Now())

	if dur <= time.Hour*23 {
		return
	}

	message := fmt.Sprintf(
		"恭喜, 当前 Opener 是 %s, MINT 了 (%s, %s), 距离宝箱打开还有 %s",
		info.MintAddress, info.X, info.Y, durationToString(dur),
	)
	_, err := discordRepo.SendDiscordMessage(ChannelId, &biz.DiscordSendMessageRequest{
		Content: message,
	})
	if err == nil {
		currentOpenerTokenId = info.TokenId
	}
}

func sendLiveMessage(info *OpenerRecord, source time.Duration) {
	timestamp := info.BlockTimestamp
	dur := time.Unix(timestamp, 0).Add(24 * time.Hour).Sub(time.Now())

	message := fmt.Sprintf(
		"恭喜, 当前 Opener 是 %s, MINT 了 (%s, %s), 距离宝箱打开还有 %s",
		info.MintAddress, info.X, info.Y, durationToString(dur),
	)
	_, err := discordRepo.SendDiscordMessage(ChannelId, &biz.DiscordSendMessageRequest{
		Content: message,
	})
	if err == nil {
		lastLiveMessageInfo = &LiveMessageHistory{
			TokenId:     info.TokenId,
			MintAddress: info.MintAddress,
			Source:      source,
		}
	}
}

func trySendOpenerLiveToDiscord(info *OpenerRecord) {
	timestamp := info.BlockTimestamp
	// 剩余 1 个小时， 10分钟，五分钟，一分钟
	dur := time.Unix(timestamp, 0).Add(24 * time.Hour).Sub(time.Now())

	if dur <= time.Second*30 {
		return
	}

	if lastLiveMessageInfo != nil && lastLiveMessageInfo.TokenId != info.TokenId {
		lastLiveMessageInfo = nil
	}
	if dur <= time.Minute {
		if lastLiveMessageInfo == nil || lastLiveMessageInfo.Source > time.Minute {
			sendLiveMessage(info, time.Minute)
		}
	}
	if dur <= time.Minute*5 {
		if lastLiveMessageInfo == nil || lastLiveMessageInfo.Source > time.Minute*5 {
			sendLiveMessage(info, time.Minute*5)
		}
	}
	if dur <= time.Minute*10 {
		if lastLiveMessageInfo == nil || lastLiveMessageInfo.Source > time.Minute*10 {
			sendLiveMessage(info, time.Minute*10)
		}
	}
	if dur <= time.Hour {
		if lastLiveMessageInfo == nil || lastLiveMessageInfo.Source > time.Hour {
			sendLiveMessage(info, time.Hour)
		}
	}
}

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

	if gameInfo.Data.OpenerRecord.TokenId != 0 && currentOpenerTokenId != gameInfo.Data.OpenerRecord.TokenId {
		go func(info *OpenerRecord) {
			sendOpenerChangeToDiscord(info)
		}(&gameInfo.Data.OpenerRecord)
	}

	if len(list) > 0 {
		logger.Println("have_news.emit_sync")
		emitSync()
		return true, err
	}
	if gameInfo.Data.OpenerRecord.TokenId != 0 {
		if listBlockTimestamp-gameInfo.Data.OpenerRecord.BlockTimestamp >= biz.WinnerTimeCon {
			logger.Println("have_winner.emit_sync")
			emitSync()
			return true, err
		}

		trySendOpenerLiveToDiscord(&gameInfo.Data.OpenerRecord)
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
	logger.Println("get_game_info")
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
	通过 get game info api 获取 game info
	if 没有获胜
		轮训
			如果有新状态，触发同步任务，暂停 15 秒 退出
			如果没有新状态，暂停 5 秒
	*/

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

func initEnv() {
	conff := &conf.Config{
		PeopleLandContractTheGraphApiUrl: config.PeopleLandContractTheGraphApiUrl,
		EthClientRawUrl:                  config.EthClientRawUrl,
		DiscordBotToken:                  config.DiscordBotToken,
	}
	peopleLandContractTheGraphRepo = data.NewPeopleLandContractTheGraphRepo(conff)
	discordRepo = data.NewDiscordRepo(conff)
}

func getConfig() {
	var c Config
	e := env.NewEnv()
	_ = e.LoadFile("./config")
	_ = e.Read(&c)

	logger.Println("EthClientRawUrl", c.EthClientRawUrl)
	logger.Println("PeopleLandContractTheGraphApiUrl", c.PeopleLandContractTheGraphApiUrl)
	logger.Println("GetRoundInfoApiUrl", c.GetRoundInfoApiUrl)
	logger.Println("SyncOpenerRecordBackgroundApiUrl", c.SyncOpenerRecordBackgroundApiUrl)
	logger.Println("PEOPLELAND_DISCORD_BOT_TOKEN", c.DiscordBotToken)

	config = &c
}

func main() {
	logger.Println("syncmonitor", 6)
	getConfig()
	initEnv()
	for {
		logger.Println("process.start")
		process()
		logger.Println("process.end")
	}
}
