package biz

import (
	v1 "backend/api/user/v1"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data/model"
	"backend/app/user/pkg"
	"backend/pkg/jwt"
	"context"
	"errors"
	"log"
	"math/big"
	"regexp"
	"strings"
)

const (
	twitterReg string = "I am verifying my identity as (\\S*) on peopleland"
)

type UserRepo interface {
	CreateUser(ctx context.Context, address string) (*model.User, error)
	GetUser(ctx context.Context, userid string) (*model.User, error)
	GetOneUserByAddress(ctx context.Context, address string) (*model.User, error)
	FindOrCreateUser(ctx context.Context, address string) (*model.User, error)
	UpdateUserByAddress(ctx context.Context, address string, updateData map[string]interface{}) (*model.User, error)
	UpdateUser(ctx context.Context, userid string, updateData map[string]interface{}) (*model.User, error)
	CreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error)
	GetOrCreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error)
	GenVerifyCode(ctx context.Context, userid string) (string, error)
	GetUserByVerifyCode(ctx context.Context, verifyCode string) (*model.User, error)
}

type TwitterRepo interface {
	GetTwitterUserTimeline(userScreenName string) []string
}

type PeopleLandContractRepo interface {
	BalanceOf(address string) (*big.Int, error)
}

type MintRecordRepo interface {
	CreateMintRecord(ctx context.Context, mintAddress string, x string, y string, userid string) (*model.MintRecord, error)
	FindLastMintRecord(_ context.Context, mintAddress string, x string, y string, mintTimestamp int64) (*model.MintRecord, error)
}

type UserUseCase struct {
	repo                   UserRepo
	twitterRepo            TwitterRepo
	peopleLandContractRepo PeopleLandContractRepo
	logger                 *log.Logger
	conf                   *conf.Config
}

func NewUserUseCase(repo UserRepo, twitterRepo TwitterRepo, peopleLandContractRepo PeopleLandContractRepo, conf *conf.Config, logger *log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                   repo,
		twitterRepo:            twitterRepo,
		peopleLandContractRepo: peopleLandContractRepo,
		logger:                 logger,
		conf:                   conf,
	}
}

func (u *UserUseCase) GetJWT(ctx context.Context, load *v1.LoginPayLoad) (*string, error) {
	verified := pkg.VerifyEip191Sig(load.Address, load.OriginMessage, load.Signature)
	if !verified {
		return nil, errors.New("request.verify.error")
	}
	address := strings.ToLower(load.Address)

	dbData, err := u.repo.FindOrCreateUser(ctx, address)
	if err != nil {
		return nil, errors.New("request.db.error")
	}

	claims := jwt.NewMapClaims(dbData.Ref.ID, address)
	jwtStr, err := jwt.EncodeJwt(claims, u.conf.JwtRsaPrivateKeyPem, int64(86400))
	if err != nil {
		return nil, errors.New("request.jwt.error")
	}
	return &jwtStr, nil
}

func (u *UserUseCase) GetProfile(ctx context.Context, address string) (*model.User, error) {
	userDb, err := u.repo.GetOneUserByAddress(ctx, address)
	if err != nil {
		return nil, err
	}
	return userDb, nil
}

func (u *UserUseCase) UpdateProfile(ctx context.Context, address string, updateData map[string]string) (*model.User, error) {
	filterUpdateData := map[string]interface{}{}
	for key, value := range updateData {
		if key == "name" || key == "twitter" {
			filterUpdateData[key] = value
		}
	}

	user, err := u.repo.UpdateUserByAddress(ctx, address, filterUpdateData)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) ConnectTwitter(ctx context.Context, address string, load *v1.ConnectTwitterPayLoad) (*model.User, error) {
	user, err := u.repo.GetOneUserByAddress(ctx, address)
	if err != nil {
		return nil, err
	}
	textList := u.twitterRepo.GetTwitterUserTimeline(load.Twitter)
	var hasT bool
	for _, text := range textList {
		reg := regexp.MustCompile(twitterReg)
		if reg == nil {
			return nil, errors.New("reg.error")
		}
		result := reg.FindAllStringSubmatch(text, -1)
		if len(result) != 0 {
			name := result[0][1]
			if name == user.Name {
				hasT = true
			}
		}
	}

	if !hasT {
		return nil, errors.New("request.twitter.error")
	}

	updateData := map[string]string{
		"twitter": load.Twitter,
	}
	profile, err := u.UpdateProfile(ctx, address, updateData)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (u *UserUseCase) GetTelegramVerifyCode(ctx context.Context, userid string) (code string, err error) {
	verifyCode, err := u.repo.GetOrCreateTelegramVerifyCode(ctx, userid)
	if err != nil {
		return
	}
	return verifyCode.Code, nil
}

func (u *UserUseCase) GenVerifyCode(ctx context.Context, userid string) (verifyCode string, err error) {
	user, err := u.repo.GetUser(ctx, userid)
	if user.VerifyCode != "" {
		return user.VerifyCode, nil
	}

	address := user.Address
	count, err := u.peopleLandContractRepo.BalanceOf(address)
	if err != nil {
		return "", errors.New("request.nft.error.unknown")
	}
	if count.Cmp(big.NewInt(0)) != 1 {
		return "", errors.New("request.nft.error.none")
	}

	verifyCode, err = u.repo.GenVerifyCode(ctx, user.Ref.ID)
	if err != nil {
		return "", errors.New("request.gen.error")
	}

	return verifyCode, err
}

type OpenerGameCase struct {
	userRepo       UserRepo
	mintRecordRepo MintRecordRepo
	logger         *log.Logger
	conf           *conf.Config
}

func NewOpenerGameCase(userRepo UserRepo, mintRecordRepo MintRecordRepo, conf *conf.Config, logger *log.Logger) *OpenerGameCase {
	return &OpenerGameCase{
		userRepo:       userRepo,
		mintRecordRepo: mintRecordRepo,
		logger:         logger,
		conf:           conf,
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
