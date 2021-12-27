package data

import (
	"backend/app/user/internal/conf"
	"log"

	"github.com/fauna/faunadb-go/v4/faunadb"
)

type Data struct {
	faunaClient *faunadb.FaunaClient
}

func NewData(conf *conf.Config, logger *log.Logger) (*Data, error) {
	fc := faunadb.NewFaunaClient(conf.FaunaDBSecret)
	logger.Println("connect fauna client ...")
	return &Data{faunaClient: fc}, nil
}
