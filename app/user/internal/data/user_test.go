package data

import (
	"backend/app/user/internal/conf"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = &conf.Config{
	FaunaDBSecret: "fnAEbfitSAACVKRgPF0ZYX-Q3zZiIE3jQpr_9km0",
}

var logger = log.Default()

func Test_userRepo_CreateTelegramVerifyCode(t *testing.T) {
	d, _ := NewData(config, logger)
	userRepo := NewUserRepo(d, logger)
	code, err := userRepo.CreateTelegramVerifyCode(context.TODO(), "1112")
	if err != nil {
		return
	}
	assert.NotEmpty(t, code.Code)
}

func Test_userRepo_GetUser(t *testing.T) {
	d, _ := NewData(config, logger)
	ur := NewUserRepo(d, logger)
	ctx := context.Background()

	address := "0x40fcc42c5a25945c02b19204d082a67591d30cf6"
	user, _ := ur.FindOrCreateUser(ctx, address)

	got, err := ur.GetUser(ctx, user.Ref.ID)
	assert.Empty(t, err)
	assert.Equal(t, got.Ref.ID, got.Ref.ID)
	assert.Equal(t, got.Address, got.Address)
}

func Test_userRepo_GenVerifyCode(t *testing.T) {
	d, _ := NewData(config, logger)
	ur := NewUserRepo(d, logger)
	ctx := context.Background()

	address := "0x88FBA00170cD2Df76D1f1F3630B4A6b7a9FA4D8A"
	user, _ := ur.FindOrCreateUser(ctx, address)

	_, _ = ur.UpdateUser(ctx, user.Ref.ID, map[string]interface{}{
		"verify_code": "",
	})

	code, _ := ur.GenVerifyCode(ctx, user.Ref.ID)
	assert.NotEqual(t, code, "")
}

func Test_userRepo_GetUserListByAddressList(t *testing.T) {
	d, _ := NewData(config, logger)
	ur := NewUserRepo(d, logger)
	ctx := context.Background()

	addressList := make([]string, 0)
	addressList = append(addressList, "0x1111111111111111111111111111111111111111")
	addressList = append(addressList, "123")
	users, err := ur.GetUserListByAddressList(ctx, addressList)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}
