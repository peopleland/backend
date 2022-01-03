package service

import (
	api "backend/api/user/v1"
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data/model"
	"backend/pkg/jwt"
	"context"
	"errors"
	"log"
)

type UserService struct {
	uc     *biz.UserUseCase
	logger *log.Logger
	conf   *conf.Config
}

func NewUserService(uc *biz.UserUseCase, conf *conf.Config, logger *log.Logger) *UserService {
	return &UserService{
		uc:     uc,
		logger: logger,
		conf:   conf,
	}
}

// 暂时这些写，后续抽取成 Middleware
func parseAuthorization(ctx context.Context, conf *conf.Config) (address string, userid string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("unauthorized")
		}
	}()
	jwtStr := ctx.Value("authorization").(string)
	jwtMap, err := jwt.DecodeJwt(jwtStr, conf.JwtRsaPublicKeyPem)
	if err != nil {
		panic(err)
	}
	address = (*jwtMap)["address"].(string)
	userid = (*jwtMap)["userid"].(string)
	return address, userid, err
}

func (u *UserService) Login(ctx context.Context, load *api.LoginPayLoad) (*api.LoginResponse, error) {
	jwtStr, err := u.uc.GetJWT(ctx, load)
	if err != nil {
		return nil, err
	}
	return &api.LoginResponse{Jwt: *jwtStr}, nil
}

func (u *UserService) parseProfile(profile *model.User) *api.UserProfile {
	var discord api.DiscordInfo
	var telegram api.TelegramInfo
	if profile.Discord != nil {
		discord = api.DiscordInfo{
			Id:            profile.Discord.ID,
			Username:      profile.Discord.Username,
			Discriminator: profile.Discord.Discriminator,
			Avatar:        profile.Discord.Avatar,
		}
	}
	if profile.Telegram != nil {
		telegram = api.TelegramInfo{
			Id:           profile.Telegram.ID,
			FirstName:    profile.Telegram.FirstName,
			Username:     profile.Telegram.Username,
			LanguageCode: profile.Telegram.LanguageCode,
		}
	}
	return &api.UserProfile{
		Address:  profile.Address,
		Discord:  &discord,
		Name:     profile.Name,
		Twitter:  profile.Twitter,
		Telegram: &telegram,
		Id:       profile.Ref.ID,
	}
}

func (u *UserService) GetProfile(ctx context.Context, load *api.GetProfilePayLoad) (*api.UserProfile, error) {
	address, _, err := parseAuthorization(ctx, u.conf)
	if err != nil {
		return nil, err
	}
	profile, err := u.uc.GetProfile(ctx, address)
	if err != nil {
		return nil, err
	}
	return u.parseProfile(profile), nil
}

func (u *UserService) PutProfile(ctx context.Context, load *api.PutProfilePayLoad) (*api.UserProfile, error) {
	address, _, err := parseAuthorization(ctx, u.conf)
	if err != nil {
		return nil, err
	}

	updateData := map[string]string{}
	if load.Name != "" {
		updateData["name"] = load.Name
	}

	profile, err := u.uc.UpdateProfile(ctx, address, updateData)
	if err != nil {
		return nil, err
	}
	return u.parseProfile(profile), nil
}

func (u *UserService) ConnectTwitter(ctx context.Context, load *api.ConnectTwitterPayLoad) (*api.UserProfile, error) {
	address, _, err := parseAuthorization(ctx, u.conf)
	if err != nil {
		return nil, err
	}

	profile, err := u.uc.ConnectTwitter(ctx, address, load)
	if err != nil {
		return nil, err
	}

	return u.parseProfile(profile), nil
}

func (u *UserService) ConnectTelegram(ctx context.Context, load *api.ConnectTelegramPayLoad) (*api.ConnectTelegramResponse, error) {
	_, userid, err := parseAuthorization(ctx, u.conf)
	if err != nil {
		return nil, err
	}

	code, err := u.uc.GetTelegramVerifyCode(ctx, userid)
	if err != nil {
		return nil, err
	}
	return &api.ConnectTelegramResponse{Code: code}, nil
}

func (u *UserService) GenVerifyCode(ctx context.Context, load *api.GenVerifyCodePayLoad) (*api.GenVerifyCodeResponse, error) {
	_, userid, err := parseAuthorization(ctx, u.conf)
	if err != nil {
		return nil, err
	}

	verifyCode, err := u.uc.GenVerifyCode(ctx, userid)
	if err != nil {
		return nil, err
	}

	return &api.GenVerifyCodeResponse{VerifyCode: verifyCode}, nil
}

func (u *UserService) ConnectDiscord(ctx context.Context, load *api.ConnectDiscordPayLoad) (*api.ConnectDiscordResponse, error) {
	_, userid, err := parseAuthorization(ctx, u.conf)
	if err != nil {
		return nil, err
	}
	_, err = u.uc.ConnectDiscord(ctx, userid, load.Code, load.RedirectUri)
	if err != nil {
		return nil, err
	}
	return &api.ConnectDiscordResponse{}, nil
}

func (u *UserService) DisconnectSocial(ctx context.Context, load *api.DisconnectSocialPayLoad) (*api.DisconnectSocialResponse, error) {
	_, userid, err := parseAuthorization(ctx, u.conf)
	if err != nil {
		return nil, err
	}
	err = u.uc.DisconnectSocial(ctx, userid, load.Social)
	if err != nil {
		return nil, err
	}
	return &api.DisconnectSocialResponse{}, nil
}

func (u *UserService) TelegramBotDMWebhooks(ctx context.Context, load *api.TelegramBotDMWebhooksPayLoad) (*api.TelegramBotDMWebhooksResponse, error) {
	if load.Message.From.IsBot {
		return nil, errors.New("cannot response bot")
	}
	_, err := u.uc.ConnectTelegram(ctx, load.Message.Text, &model.TelegramUser{
		ID:           load.Message.From.Id,
		FirstName:    load.Message.From.FirstName,
		Username:     load.Message.From.Username,
		LanguageCode: load.Message.From.LanguageCode,
	})
	if err != nil {
		return nil, err
	}
	return &api.TelegramBotDMWebhooksResponse{}, nil
}
