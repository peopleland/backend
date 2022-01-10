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
	"strconv"
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
var uniswap *UniswapRepo

var lastLiveMessageInfo *LiveMessageHistory

var currentOpenerTokenId int64

const ChannelId = "929277955203547136"

func getEthAndBuilderAmount(ethCount string, builderCount string) (string, error) {
	ethCountFloat, err := strconv.ParseFloat(ethCount, 64)
	if err != nil {
		return "", nil
	}
	builderCountFloat, err := strconv.ParseFloat(builderCount, 64)
	if err != nil {
		return "", nil
	}
	builderPrice, err := uniswap.GetBuilderPrice()
	if err != nil {
		return "", nil
	}
	ethPrice, err := uniswap.GetEthPrice()
	if err != nil {
		return "", nil
	}
	builderPriceFloat, err := strconv.ParseFloat(builderPrice, 64)
	if err != nil {
		return "", nil
	}
	ethPriceFloat, err := strconv.ParseFloat(ethPrice, 64)
	if err != nil {
		return "", nil
	}
	result := ethPriceFloat*ethCountFloat + builderPriceFloat*builderCountFloat
	return fmt.Sprintf("%.2f", result), nil
}

func durationToString(dur time.Duration) string {
	hour := int64(dur.Hours())
	minute := int64(dur.Minutes()) - hour*60
	second := int64(dur.Seconds()) - hour*60*60 - minute*60
	if hour > 0 {
		return fmt.Sprintf("%d hours and %d minutes", hour, minute)
	}
	if minute > 0 {
		return fmt.Sprintf("%d minutes", minute)
	}
	return fmt.Sprintf("%d seconds", second)
}

func durationToCnString(dur time.Duration) string {
	hour := int64(dur.Hours())
	minute := int64(dur.Minutes()) - hour*60
	second := int64(dur.Seconds()) - hour*60*60 - minute*60
	if hour > 0 {
		return fmt.Sprintf("%d 小时 %d 分", hour, minute)
	}
	if minute > 0 {
		return fmt.Sprintf("%d 分钟", minute)
	}
	return fmt.Sprintf("%d 秒", second)
}

func sendOpenerChangeToDiscord(gameInfo *GameInfo) {
	amount, err := getEthAndBuilderAmount(gameInfo.Data.Info.EthAmount, gameInfo.Data.Info.BuilderTokenAmount)
	if err != nil {
		return
	}

	info := gameInfo.Data.OpenerRecord
	timestamp := info.BlockTimestamp
	dur := time.Unix(timestamp, 0).Add(24 * time.Hour).Sub(time.Now())

	if dur <= time.Hour*23 {
		return
	}

	message := fmt.Sprintf(
		"恭喜, %s 成为新任 Opener, MINT 了 (%s, %s), 距离宝箱打开还有 %s，当前宝箱总价值 $%s !!!\n\nCongratulations, %s is the new Opener, MINT (%s, %s), %s more to go until the chest is opened, the total value of the current chest is $%s !!!",
		info.MintAddress, info.X, info.Y, durationToCnString(dur), amount,
		info.MintAddress, info.X, info.Y, durationToString(dur), amount,
	)
	_, err = discordRepo.SendDiscordMessage(ChannelId, &biz.DiscordSendMessageRequest{
		Content: message,
	})
	if err == nil {
		currentOpenerTokenId = info.TokenId
	}
}

func sendLiveMessage(gameInfo *GameInfo, source time.Duration) {
	amount, err := getEthAndBuilderAmount(gameInfo.Data.Info.EthAmount, gameInfo.Data.Info.BuilderTokenAmount)
	if err != nil {
		return
	}

	info := &gameInfo.Data.OpenerRecord
	timestamp := info.BlockTimestamp
	dur := time.Unix(timestamp, 0).Add(24 * time.Hour).Sub(time.Now())

	message := fmt.Sprintf(
		"恭喜, 当前 Opener 是 %s, MINT 了 (%s, %s), 距离宝箱打开还有 %s，当前宝箱总价值 $%s !!!\n\nCongratulations, the current Opener is %s, MINT (%s, %s), %s to open the treasure chest, the total value of the current treasure chest is $%s !!!",
		info.MintAddress, info.X, info.Y, durationToCnString(dur), amount,
		info.MintAddress, info.X, info.Y, durationToString(dur), amount,
	)
	_, err = discordRepo.SendDiscordMessage(ChannelId, &biz.DiscordSendMessageRequest{
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

func trySendOpenerLiveToDiscord(gameInfo *GameInfo) {
	info := &gameInfo.Data.OpenerRecord
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
			sendLiveMessage(gameInfo, time.Minute)
		}
	}
	if dur <= time.Minute*5 {
		if lastLiveMessageInfo == nil || lastLiveMessageInfo.Source > time.Minute*5 {
			sendLiveMessage(gameInfo, time.Minute*5)
		}
	}
	if dur <= time.Minute*10 {
		if lastLiveMessageInfo == nil || lastLiveMessageInfo.Source > time.Minute*10 {
			sendLiveMessage(gameInfo, time.Minute*10)
		}
	}
	if dur <= time.Hour {
		if lastLiveMessageInfo == nil || lastLiveMessageInfo.Source > time.Hour {
			sendLiveMessage(gameInfo, time.Hour)
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
		list, listBlockTimestamp, err = peopleLandContractTheGraphRepo.GetTokenInfoListByFromTimestamp(gameInfo.Data.Info.StartTimestamp)
		if err != nil {
			return false, err
		}
	}

	if gameInfo.Data.OpenerRecord.TokenId != 0 && currentOpenerTokenId != gameInfo.Data.OpenerRecord.TokenId {
		go func(gameInfo *GameInfo) {
			sendOpenerChangeToDiscord(gameInfo)
		}(gameInfo)
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

		trySendOpenerLiveToDiscord(gameInfo)
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
	uniswap = NewUniswapRepo()
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
	logger.Println("syncmonitor", 9)
	getConfig()
	initEnv()
	for {
		logger.Println("process.start")
		process()
		logger.Println("process.end")
	}
}
