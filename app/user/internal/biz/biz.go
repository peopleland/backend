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
	"strings"
)

type UserRepo interface {
	CreateUser(ctx context.Context, address string) (*model.UserDb, error)
	GetOneUserByAddress(ctx context.Context, address string) (*model.UserDb, error)
	FindOrCreateUser(ctx context.Context, address string) (*model.UserDb, error)
	UpdateUserByAddress(ctx context.Context, address string, updateData map[string]interface{}) (*model.UserDb, error)
}

type UserUseCase struct {
	repo   UserRepo
	logger *log.Logger
	conf   *conf.Config
}

func NewUserUseCase(repo UserRepo, conf *conf.Config, logger *log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		logger: logger,
		conf:   conf,
	}
}

func (u *UserUseCase) GetJWT(ctx context.Context, load *v1.LoginPayLoad) (*string, error) {
	verified := pkg.VerifyEip191Sig(*load.Address, *load.OriginMessage, *load.Signature)
	if !verified {
		return nil, errors.New("request.verify.error")
	}
	address := strings.ToLower(*load.Address)

	_, err := u.repo.FindOrCreateUser(ctx, address)
	if err != nil {
		return nil, errors.New("request.db.error")
	}

	claims := jwt.NewMapClaims(address)
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
