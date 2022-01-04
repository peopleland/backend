package biz

import (
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data/model"
	"context"
	"errors"
	"log"
)

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

type PeopleLandTokenInfo struct {
	TokenId            int32
	X                  int32
	Y                  int32
	MintedAddress      string
	OwnerAddress       string
	GivedAtTimestamp   int32
	GivedAtBlockNumber int32
}

type OpenerGameCase struct {
	userRepo                       UserRepo
	mintRecordRepo                 MintRecordRepo
	openerRecordRepo               OpenerRecordRepo
	openerGameRoundInfoRepo        OpenerGameRoundInfoRepo
	peopleLandContractTheGraphRepo PeopleLandContractTheGraphRepo
	logger                         *log.Logger
	conf                           *conf.Config
}

func NewOpenerGameCase(userRepo UserRepo, mintRecordRepo MintRecordRepo, openerRecordRepo OpenerRecordRepo, openerGameRoundInfoRepo OpenerGameRoundInfoRepo, peopleLandContractTheGraphRepo PeopleLandContractTheGraphRepo, conf *conf.Config, logger *log.Logger) *OpenerGameCase {
	return &OpenerGameCase{
		userRepo:                       userRepo,
		mintRecordRepo:                 mintRecordRepo,
		openerRecordRepo:               openerRecordRepo,
		openerGameRoundInfoRepo:        openerGameRoundInfoRepo,
		peopleLandContractTheGraphRepo: peopleLandContractTheGraphRepo,
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

func (ogc *OpenerGameCase) GetOpenerRecordList(ctx context.Context, pageSize int64, afterTokenId int64, beforeTokenId int64) ([]*OpenerRecordWithUserName, error) {
	if afterTokenId != 0 && beforeTokenId != 0 {
		return nil, errors.New("query.params.error")
	}

	var data []*model.OpenerRecord
	var err error
	if afterTokenId != 0 {
		data, err = ogc.openerRecordRepo.GetListPaginateAfter(ctx, pageSize, afterTokenId)
	}
	if beforeTokenId != 0 {
		data, err = ogc.openerRecordRepo.GetListPaginateBefore(ctx, pageSize, beforeTokenId)
	}
	if afterTokenId == 0 && beforeTokenId == 0 {
		data, err = ogc.openerRecordRepo.GetListPaginateFirstPage(ctx, pageSize)
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
	return list, nil
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

func (ogc *OpenerGameCase) SyncOpenerRecord(_ context.Context) {
	// TODO sync opner record
}
