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

type UserUseCase struct {
	repo                   UserRepo
	twitterRepo            TwitterRepo
	discordRepo            DiscordRepo
	peopleLandContractRepo PeopleLandContractRepo
	logger                 *log.Logger
	conf                   *conf.Config
}

func NewUserUseCase(repo UserRepo, twitterRepo TwitterRepo, discordRepo DiscordRepo, peopleLandContractRepo PeopleLandContractRepo, conf *conf.Config, logger *log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                   repo,
		twitterRepo:            twitterRepo,
		discordRepo:            discordRepo,
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

func (u *UserUseCase) ConnectDiscord(ctx context.Context, userid, code, redirectUri string) (*model.User, error) {
	_, err := u.repo.GetUser(ctx, userid)
	if err != nil {
		return nil, err
	}
	discordUser, err := u.discordRepo.GetDiscordInfo(code, redirectUri)
	if err != nil {
		return nil, err
	}
	return u.repo.UpdateUser(ctx, userid, map[string]interface{}{"discord": discordUser})
}

func (u *UserUseCase) ConnectTelegram(ctx context.Context, code string, telegramUser *model.TelegramUser) (*model.User, error) {
	userid, err := u.repo.GetUserByTelegramVerifyCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return u.repo.UpdateUser(ctx, userid, map[string]interface{}{"telegram": telegramUser})
}

func (u *UserUseCase) DisconnectSocial(ctx context.Context, userid string, socialType v1.SocialType) error {
	_, err := u.repo.UpdateUser(ctx, userid, map[string]interface{}{strings.ToLower(socialType.String()): nil})
	if err != nil {
		return err
	}
	return nil
}
