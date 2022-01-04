package server

import (
	v1 "backend/api/user/v1"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/service"
	"backend/pkg/http"
	"log"
)

func NewLambdaServer(_ *conf.Config, userServ *service.UserService, _ *log.Logger) *http.Server {
	serv := http.NewServer()

	v1.RegisterUserLambdaServer(serv, userServ)
	return serv
}
