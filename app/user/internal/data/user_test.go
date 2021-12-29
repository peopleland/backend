package data

import (
	"backend/app/user/internal/conf"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = &conf.Config{
	FaunaDBSecret: "fnAEbfjifeACVMELXa_tc0wdOe5SqgdXDdJd-zUR",
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
