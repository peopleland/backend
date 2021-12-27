package service

import (
	api "backend/api/user/v1"
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/pkg/jwt"
	"context"
	"fmt"
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

func (u *UserService) Login(ctx context.Context, load *api.LoginPayLoad) (*api.LoginResponse, error) {
	jwtStr, err := u.uc.GetJWT(ctx, load)
	if err != nil {
		return nil, err
	}
	return &api.LoginResponse{Jwt: jwtStr}, nil
}

func (u *UserService) GetProfile(ctx context.Context) (*api.UserProfile, error) {
	jwtStr := ctx.Value("authorization").(string)
	jwtMap, err := jwt.DecodeJwt(jwtStr, u.conf.JwtRsaPublicKeyPem)
	if err != nil {
		return nil, err
	}
	fmt.Println(*jwtMap)
	address := (*jwtMap)["address"].(string)
	profile, err := u.uc.GetProfile(ctx, address)
	if err != nil {
		return nil, err
	}
	return &api.UserProfile{
		Address: profile.Address,
		Discord: "",
		Name:    profile.Name,
		Twitter: profile.Twitter,
	}, nil
}

func (u *UserService) PutProfile(ctx context.Context, load *api.PutProfilePayLoad) (*api.UserProfile, error) {
	jwtStr := ctx.Value("authorization").(string)
	jwtMap, err := jwt.DecodeJwt(jwtStr, u.conf.JwtRsaPublicKeyPem)
	if err != nil {
		return nil, err
	}
	address := (*jwtMap)["address"].(string)

	updateData := map[string]string{}
	if load.Twitter != "" {
		updateData["twitter"] = load.Twitter
	}
	if load.Name != "" {
		updateData["name"] = load.Name
	}

	profile, err := u.uc.UpdateProfile(ctx, address, updateData)
	if err != nil {
		return nil, err
	}
	return &api.UserProfile{
		Address: profile.Address,
		Discord: "",
		Name:    profile.Name,
		Twitter: profile.Twitter,
	}, nil
}
