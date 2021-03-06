package biz

import (
	"backend/app/user/internal/data/model"
	"context"
	"math/big"
	"time"
)

type DiscordSendMessageRequest struct {
	Content string `json:"content"`
	Tts     bool   `json:"tts"`
	Embeds  []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"embeds"`
}

type DiscordMessageResponse struct {
	Id              string    `json:"id"`
	Type            int       `json:"type"`
	Content         string    `json:"content"`
	ChannelId       string    `json:"channel_id"`
	MentionEveryone bool      `json:"mention_everyone"`
	Timestamp       time.Time `json:"timestamp"`
}

type UserRepo interface {
	CreateUser(ctx context.Context, address string) (*model.User, error)
	GetUser(ctx context.Context, userid string) (*model.User, error)
	GetOneUserByAddress(ctx context.Context, address string) (*model.User, error)
	GetUserListByAddressList(ctx context.Context, addressList []string) ([]*model.User, error)
	FindOrCreateUser(ctx context.Context, address string) (*model.User, error)
	UpdateUserByAddress(ctx context.Context, address string, updateData map[string]interface{}) (*model.User, error)
	UpdateUser(ctx context.Context, userid string, updateData map[string]interface{}) (*model.User, error)
	CreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error)
	GetOrCreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error)
	GenVerifyCode(ctx context.Context, userid string) (string, error)
	GetUserByVerifyCode(ctx context.Context, verifyCode string) (*model.User, error)
	GetUserByTelegramVerifyCode(ctx context.Context, code string) (string, error)
}

type TwitterRepo interface {
	GetTwitterUserTimeline(userScreenName string) []string
}

type DiscordRepo interface {
	GetDiscordInfo(code, redirectURI string) (*model.DiscordUser, error)
	SendDiscordMessage(channelId string, request *DiscordSendMessageRequest) (*DiscordMessageResponse, error)
}

type PeopleLandContractRepo interface {
	BalanceOf(address string) (*big.Int, error)
	GetEthBalanceAt(blockNumber int64) (string, error)
}

type MintRecordRepo interface {
	CreateMintRecord(ctx context.Context, mintAddress string, x string, y string, userid string) (*model.MintRecord, error)
	FindLastMintRecord(ctx context.Context, mintAddress string, x string, y string, mintTimestamp int64) (*model.MintRecord, error)
}

type OpenerRecordRepo interface {
	GetListPaginateFirstPage(ctx context.Context, pageSize int64) (list []*model.OpenerRecord, afterTokenId int64, err error)
	GetListPaginateAfter(ctx context.Context, pageSize int64, inputAfterTokenId int64) (list []*model.OpenerRecord, beforeTokenId int64, afterTokenId int64, err error)
	GetListPaginateBefore(ctx context.Context, pageSize int64, inputBeforeTokenId int64) (list []*model.OpenerRecord, beforeTokenId int64, afterTokenId int64, err error)
	SaveOpenerRecord(ctx context.Context, tokenId int64, data *model.OpenerRecord) (*model.OpenerRecord, error)
	CreateOpenerRecord(ctx context.Context, tokenId int64, data *model.OpenerRecord) (*model.OpenerRecord, error)
	UpdateOpenerRecord(ctx context.Context, tokenId int64, data *model.OpenerRecord) (*model.OpenerRecord, error)
	GetNewest(ctx context.Context) (*model.OpenerRecord, error)
	GetOpenerRecordByTokenId(ctx context.Context, tokenId int64) (*model.OpenerRecord, error)
	GetTotalCount(ctx context.Context) (int64, error)
}

type OpenerGameRoundInfoRepo interface {
	GetByRoundNumber(ctx context.Context, roundNumber int64) (*model.OpenerGameRoundInfo, error)
	Create(ctx context.Context, roundNumber int64, data *model.OpenerGameRoundInfo) (*model.OpenerGameRoundInfo, error)
	Update(ctx context.Context, roundNumber int64, data *model.OpenerGameRoundInfo) (*model.OpenerGameRoundInfo, error)
	Save(ctx context.Context, roundNumber int64, data *model.OpenerGameRoundInfo) (*model.OpenerGameRoundInfo, error)
	UpdateEth(ctx context.Context, roundNumber int64, eth string) (*model.OpenerGameRoundInfo, error)
}

type PeopleLandContractTheGraphRepo interface {
	GetTokenInfoListByFromTokenId(fromTokenId int64) (list []*PeopleLandTokenInfo, blockTimestamp int64, err error)
	GetTokenInfoListByFromTimestamp(fromTimestamp int64) (list []*PeopleLandTokenInfo, blockTimestamp int64, err error)
}
