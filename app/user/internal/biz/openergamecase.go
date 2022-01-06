package biz

import (
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data/model"
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"log"
	"strconv"
)

var WinnerTimeCon int64 = 24 * 60 * 60

type OpenerRecordWithUserName struct {
	MintAddress             string
	MintUserName            string
	TokenId                 int64
	X                       string
	Y                       string
	BlockNumber             int64
	BlockTimestamp          int64
	InvitedAddress          string
	InvitedUserName         string
	NextTokenBlockTimestamp int64
}

type OpenerRecordListPaginate struct {
	List          []*OpenerRecordWithUserName
	TotalCount    int64
	AfterTokenId  int64
	BeforeTokenId int64
}

type PeopleLandTokenInfo struct {
	TokenId            int64
	X                  string
	Y                  string
	MintedAddress      string
	OwnerAddress       string
	GivedAtTimestamp   int64
	GivedAtBlockNumber int64
}

type OpenerGameCase struct {
	userRepo                       UserRepo
	mintRecordRepo                 MintRecordRepo
	openerRecordRepo               OpenerRecordRepo
	openerGameRoundInfoRepo        OpenerGameRoundInfoRepo
	peopleLandContractRepo         PeopleLandContractRepo
	peopleLandContractTheGraphRepo PeopleLandContractTheGraphRepo
	logger                         *log.Logger
	conf                           *conf.Config
}

func NewOpenerGameCase(userRepo UserRepo, mintRecordRepo MintRecordRepo, openerRecordRepo OpenerRecordRepo, openerGameRoundInfoRepo OpenerGameRoundInfoRepo, peopleLandContractTheGraphRepo PeopleLandContractTheGraphRepo, peopleLandContractRepo PeopleLandContractRepo, conf *conf.Config, logger *log.Logger) *OpenerGameCase {
	return &OpenerGameCase{
		userRepo:                       userRepo,
		mintRecordRepo:                 mintRecordRepo,
		openerRecordRepo:               openerRecordRepo,
		openerGameRoundInfoRepo:        openerGameRoundInfoRepo,
		peopleLandContractTheGraphRepo: peopleLandContractTheGraphRepo,
		peopleLandContractRepo:         peopleLandContractRepo,
		logger:                         logger,
		conf:                           conf,
	}
}

func (ogc *OpenerGameCase) CreateMintRecord(ctx context.Context, mintAddress string, x string, y string, verifyCode string) (*model.MintRecord, error) {
	user, err := ogc.userRepo.GetUserByVerifyCode(ctx, verifyCode)
	if err != nil {
		return nil, err
	}
	mintRecord, err := ogc.mintRecordRepo.CreateMintRecord(ctx, mintAddress, x, y, user.Ref.ID)
	if err != nil {
		return nil, err
	}

	return mintRecord, nil
}

func (ogc *OpenerGameCase) FindInvitedUser(ctx context.Context, mintAddress string, x string, y string, mintTimestamp int64) (*model.User, error) {
	mr, err := ogc.mintRecordRepo.FindLastMintRecord(ctx, mintAddress, x, y, mintTimestamp)
	if err != nil || mr == nil {
		return nil, nil
	}
	user, err := ogc.userRepo.GetUser(ctx, mr.InviteUserid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ogc *OpenerGameCase) GetOpenerRecordList(ctx context.Context, pageSize int64, inputAfterTokenId int64, inputBeforeTokenId int64) (*OpenerRecordListPaginate, error) {
	if inputAfterTokenId != 0 && inputBeforeTokenId != 0 {
		return nil, errors.New("query.params.error")
	}

	var err error
	totalCount, err := ogc.openerRecordRepo.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	var data []*model.OpenerRecord
	var afterTokenId int64
	var beforeTokenId int64

	if inputAfterTokenId != 0 {
		data, beforeTokenId, afterTokenId, err = ogc.openerRecordRepo.GetListPaginateAfter(ctx, pageSize, inputAfterTokenId)
	}
	if inputBeforeTokenId != 0 {
		data, beforeTokenId, afterTokenId, err = ogc.openerRecordRepo.GetListPaginateBefore(ctx, pageSize, inputBeforeTokenId)
	}
	if inputAfterTokenId == 0 && inputBeforeTokenId == 0 {
		data, afterTokenId, err = ogc.openerRecordRepo.GetListPaginateFirstPage(ctx, pageSize)
	}
	if err != nil {
		return nil, err
	}

	addressMap := make(map[string]bool)
	for _, item := range data {
		addressMap[item.MintAddress] = true
		addressMap[item.InvitedAddress] = true
	}
	addressList := make([]string, 0)
	for address, _ := range addressMap {
		addressList = append(addressList, address)
	}

	users, err := ogc.userRepo.GetUserListByAddressList(ctx, addressList)
	if err != nil {
		return nil, err
	}

	var userAddress2Name = make(map[string]string)
	for _, user := range users {
		if user.Name != "" {
			userAddress2Name[user.Address] = user.Name
		}
	}

	list := make([]*OpenerRecordWithUserName, 0)
	for _, item := range data {
		list = append(list, &OpenerRecordWithUserName{
			MintAddress:             item.MintAddress,
			MintUserName:            userAddress2Name[item.MintAddress],
			TokenId:                 item.TokenId,
			X:                       item.X,
			Y:                       item.Y,
			BlockNumber:             item.BlockNumber,
			BlockTimestamp:          item.BlockTimestamp,
			InvitedAddress:          item.InvitedAddress,
			InvitedUserName:         userAddress2Name[item.InvitedAddress],
			NextTokenBlockTimestamp: item.NextTokenBlockTimestamp,
		})
	}
	return &OpenerRecordListPaginate{
		List:          list,
		TotalCount:    totalCount,
		AfterTokenId:  afterTokenId,
		BeforeTokenId: beforeTokenId,
	}, nil
}

func (ogc *OpenerGameCase) GetOpenerGameRoundInfo(ctx context.Context, roundNumber int64) (*model.OpenerGameRoundInfo, *OpenerRecordWithUserName, error) {
	var info *model.OpenerGameRoundInfo
	var record *model.OpenerRecord
	var err error
	info, err = ogc.openerGameRoundInfoRepo.GetByRoundNumber(ctx, roundNumber)
	if err != nil {
		return nil, nil, err
	}
	if info.HasWinner {
		record, err = ogc.openerRecordRepo.GetOpenerRecordByTokenId(ctx, info.WinnerTokenId)
	} else {
		record, err = ogc.openerRecordRepo.GetNewest(ctx)
	}
	if err != nil {
		return nil, nil, err
	}
	if record == nil {
		return info, nil, nil
	}

	addressList := make([]string, 0)
	if record.MintAddress != "" {
		addressList = append(addressList, record.MintAddress)
	}
	if record.InvitedAddress != "" {
		addressList = append(addressList, record.InvitedAddress)
	}

	users, err := ogc.userRepo.GetUserListByAddressList(ctx, addressList)
	if err != nil {
		return nil, nil, err
	}

	var userAddress2Name = make(map[string]string)
	for _, user := range users {
		if user.Name != "" {
			userAddress2Name[user.Address] = user.Name
		}
	}

	return info, &OpenerRecordWithUserName{
		MintAddress:             record.MintAddress,
		MintUserName:            userAddress2Name[record.MintAddress],
		TokenId:                 record.TokenId,
		X:                       record.X,
		Y:                       record.Y,
		BlockNumber:             record.BlockNumber,
		BlockTimestamp:          record.BlockTimestamp,
		InvitedAddress:          record.InvitedAddress,
		InvitedUserName:         userAddress2Name[record.InvitedAddress],
		NextTokenBlockTimestamp: record.NextTokenBlockTimestamp,
	}, nil
}

func (ogc *OpenerGameCase) SyncOpenerRecord(ctx context.Context) {
	/*
		获取 game info
		if game info 不存在
			退出同步
		if 已经有获胜者
			退出同步
	*/
	info, err := ogc.openerGameRoundInfoRepo.GetByRoundNumber(ctx, 1)
	if err != nil {
		ogc.logger.Println("1.error", err)
		return
	}
	if info == nil {
		ogc.logger.Println("round_info.not_exists")
		return
	}
	if info.HasWinner {
		ogc.logger.Println("round_info.has_winner")
		return
	}

	/*
		从 db 获取最新 opener_record  newest_token_id
		if 不存在
			获取 game info 开始时间
			从开始时间查询 the graph api 获取 un_sync_token_list
		else
			从 newest_token_id + 1 查询 the graph api 获取 un_sync_token_list

		获得 un_sync_token_list 是 id 正序的数组
		if 数组为空
			if newest_token 存在
				判断是否有获胜者，并更新 info
			退出同步
	*/
	var unSyncTokenList []*PeopleLandTokenInfo
	var thegraphTimestamp int64

	newestRecord, err := ogc.openerRecordRepo.GetNewest(ctx)
	if err != nil {
		ogc.logger.Println("2.error", err)
		return
	}
	if newestRecord == nil {
		unSyncTokenList, thegraphTimestamp, err = ogc.peopleLandContractTheGraphRepo.GetTokenInfoListByFromTimestamp(info.StartTimestamp)
		if err != nil {
			ogc.logger.Println("3.error", err)
			return
		}
	} else {
		unSyncTokenList, thegraphTimestamp, err = ogc.peopleLandContractTheGraphRepo.GetTokenInfoListByFromTokenId(newestRecord.TokenId + 1)
		if err != nil {
			ogc.logger.Println("4.error", err)
			return
		}
	}
	if (len(unSyncTokenList)) == 0 {
		ogc.logger.Println("un_sync_token_list.empty")
		if newestRecord != nil {
			if thegraphTimestamp-newestRecord.BlockTimestamp >= WinnerTimeCon {
				err := ogc.setWinner(ctx, 1, info, newestRecord)
				if err != nil {
					ogc.logger.Println("5.error", err)
				}
			}
		}
		return
	}

	/*
		if newestRecord 存在
			把 un_sync_token_list[0].block info 更新到 opener_record(newest_token_id)
			if opener_record(newest_token_id) 获胜
				更新 info 信息，更新获胜者
				退出同步
	*/
	if newestRecord != nil {
		if unSyncTokenList[0].TokenId != newestRecord.TokenId+1 {
			ogc.logger.Println("sync.next_token_block_timestamp.error")
			return
		}
		newestRecord.NextTokenBlockTimestamp = unSyncTokenList[0].GivedAtTimestamp
		_, err = ogc.openerRecordRepo.UpdateOpenerRecord(ctx, newestRecord.TokenId, newestRecord)
		if err != nil {
			ogc.logger.Println("6.error", err)
			return
		}
		if newestRecord.NextTokenBlockTimestamp-newestRecord.BlockTimestamp >= WinnerTimeCon {
			err := ogc.setWinner(ctx, 1, info, newestRecord)
			if err != nil {
				ogc.logger.Println("7.error", err)
				return
			}
			return
		}
	}

	/*
		遍历 un_sync_token_list
			if current 不是数组最后一个
				current merge un_sync_token_list[current_index+1]。block_info
			sync current to db
			if opener_record(current) 获胜
				更新 info 信息，更新获胜者
				退出同步
	*/
	total := len(unSyncTokenList)
	for index, item := range unSyncTokenList {
		mintRecord, err := ogc.mintRecordRepo.FindLastMintRecord(
			ctx,
			item.OwnerAddress,
			item.X,
			item.Y,
			item.GivedAtTimestamp,
		)
		if err != nil {
			ogc.logger.Println("8.error", err)
			return
		}
		invitedAddress := item.MintedAddress
		if mintRecord != nil {
			user, err := ogc.userRepo.GetUser(ctx, mintRecord.InviteUserid)
			if err != nil {
				ogc.logger.Println("9.error", err)
				return
			}
			invitedAddress = user.Address
		}
		record := &model.OpenerRecord{
			MintAddress:    item.OwnerAddress,
			TokenId:        item.TokenId,
			X:              item.X,
			Y:              item.Y,
			BlockNumber:    item.GivedAtBlockNumber,
			BlockTimestamp: item.GivedAtTimestamp,
			InvitedAddress: invitedAddress,
		}
		if index != total-1 {
			record.NextTokenBlockTimestamp = unSyncTokenList[index+1].GivedAtTimestamp
		}
		_, err = ogc.openerRecordRepo.CreateOpenerRecord(ctx, record.TokenId, record)
		if err != nil {
			ogc.logger.Println("10.error", err)
			return
		}
		ogc.logger.Println("sync.token." + strconv.FormatInt(record.TokenId, 10) + ".success")

		if index != total-1 {
			if record.NextTokenBlockTimestamp-record.BlockTimestamp >= WinnerTimeCon {
				err := ogc.setWinner(ctx, 1, info, record)
				if err != nil {
					ogc.logger.Println("11.error", err)
					return
				}
				return
			}
		}

	}
}

func (ogc *OpenerGameCase) setWinner(ctx context.Context, roundNumber int64, info *model.OpenerGameRoundInfo, winnerRecord *model.OpenerRecord) error {
	info.HasWinner = true
	info.WinnerTokenId = winnerRecord.TokenId
	info.EndTimestamp = winnerRecord.BlockTimestamp + WinnerTimeCon
	_, err := ogc.openerGameRoundInfoRepo.Update(ctx, roundNumber, info)
	if err != nil {
		return err
	}
	ogc.logger.Println("sync.token.set_winner." + strconv.FormatInt(info.WinnerTokenId, 10) + ".success")
	return nil
}

func (ogc *OpenerGameCase) SyncRoundInfoEth(ctx context.Context) {
	info, err := ogc.openerGameRoundInfoRepo.GetByRoundNumber(ctx, 1)
	if err != nil {
		ogc.logger.Println("12.error", err)
		return
	}
	if info == nil {
		ogc.logger.Println("round_info.not_exists")
		return
	}

	var record *model.OpenerRecord
	record, err = ogc.openerRecordRepo.GetNewest(ctx)
	if err != nil {
		ogc.logger.Println("13.error", err)
		return
	}
	if record == nil {
		ogc.logger.Println("opener_record.newest.nil")
		return
	}
	if info.HasWinner {
		record, err = ogc.openerRecordRepo.GetOpenerRecordByTokenId(ctx, info.WinnerTokenId)
		if err != nil {
			ogc.logger.Println("14.error", err)
			return
		}
		if record == nil {
			ogc.logger.Println("opener_record.winner_token.nil")
			return
		}
	}

	ethStr, err := ogc.peopleLandContractRepo.GetEthBalanceAt(record.BlockNumber)
	if err != nil {
		ogc.logger.Println("15.error", err)
		return
	}

	ethDecimal, err := decimal.NewFromString(ethStr)
	if err != nil {
		ogc.logger.Println("16.error", err)
		return
	}
	subEthDecimal, err := decimal.NewFromString("1.32")
	if err != nil {
		ogc.logger.Println("17.error", err)
		return
	}
	divEthDecimal, err := decimal.NewFromString("2")
	if err != nil {
		ogc.logger.Println("18.error", err)
		return
	}
	result := ethDecimal.Sub(subEthDecimal)
	resultStr := result.Div(divEthDecimal).String()

	if info.EthAmount == resultStr {
		return
	}
	_, err = ogc.openerGameRoundInfoRepo.UpdateEth(ctx, info.RoundNumber, resultStr)
	if err != nil {
		ogc.logger.Println("19.error", err)
		return
	}
	return
}
