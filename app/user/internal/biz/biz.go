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
	"regexp"
	"strings"
)

const (
	twitterReg string = "I am verifying my identity as (\\S*) on peopleland"
)

type UserRepo interface {
	CreateUser(ctx context.Context, address string) (*model.UserDb, error)
	GetOneUserByAddress(ctx context.Context, address string) (*model.UserDb, error)
	FindOrCreateUser(ctx context.Context, address string) (*model.UserDb, error)
	UpdateUserByAddress(ctx context.Context, address string, updateData map[string]interface{}) (*model.UserDb, error)
	CreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error)
	GetOrCreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error)
}

type TwitterRepo interface {
	GetTwitterUserTimeline(userScreenName string) []string
}

type UserUseCase struct {
	repo        UserRepo
	twitterRepo TwitterRepo
	logger      *log.Logger
	conf        *conf.Config
}

func NewUserUseCase(repo UserRepo, twitterRepo TwitterRepo, conf *conf.Config, logger *log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:        repo,
		twitterRepo: twitterRepo,
		logger:      logger,
		conf:        conf,
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
	return &userDb.Data, nil
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

	return &user.Data, nil
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
			if name == user.Data.Name {
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
