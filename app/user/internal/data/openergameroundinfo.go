package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/data/model"
	"context"
	"errors"
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"log"
	"regexp"
)

type OpenerGameRoundInfoRepo struct {
	data   *Data
	logger *log.Logger
}

func NewOpenerGameRoundInfoRepo(data *Data, logger *log.Logger) biz.OpenerGameRoundInfoRepo {
	return &OpenerGameRoundInfoRepo{
		data:   data,
		logger: logger,
	}
}

func (repo *OpenerGameRoundInfoRepo) GetByRoundNumber(_ context.Context, roundNumber int64) (*model.OpenerGameRoundInfo, error) {
	result, err := repo.data.faunaClient.Query(
		f.Get(
			f.MatchTerm(
				f.Index(model.OpenerGameRoundInfoByRoundNumberIndex),
				roundNumber,
			),
		))
	if err != nil {
		return nil, err
	}
	var record model.OpenerGameRoundInfo
	err = model.ParseResult(result, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (repo *OpenerGameRoundInfoRepo) Create(_ context.Context, roundNumber int64, data *model.OpenerGameRoundInfo) (*model.OpenerGameRoundInfo, error) {
	data.RoundNumber = roundNumber
	result, err := repo.data.faunaClient.Query(
		f.Create(f.Collection(model.OpenerGameRoundInfoCollectionName),
			f.Obj{"data": data}))
	if err != nil {
		return nil, err
	}
	var record model.OpenerGameRoundInfo
	err = model.ParseResult(result, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (repo *OpenerGameRoundInfoRepo) Update(_ context.Context, roundNumber int64, data *model.OpenerGameRoundInfo) (*model.OpenerGameRoundInfo, error) {
	data.RoundNumber = roundNumber
	result, err := repo.data.faunaClient.Query(
		f.Update(
			f.Select(
				[]string{"ref"},
				f.Get(
					f.MatchTerm(
						f.Index(model.OpenerGameRoundInfoByRoundNumberIndex),
						roundNumber,
					),
				),
			),
			f.Obj{"data": data},
		),
	)

	if err != nil {
		err1, ok := err.(f.BadRequest)
		if ok {
			has, _ := regexp.MatchString("document is not unique", err1.Error())
			if has {
				return nil, errors.New("db.document.not_unique")
			}
		}
		return nil, err
	}
	var record model.OpenerGameRoundInfo
	err = model.ParseResult(result, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (repo *OpenerGameRoundInfoRepo) Save(ctx context.Context, roundNumber int64, data *model.OpenerGameRoundInfo) (*model.OpenerGameRoundInfo, error) {
	record, err := repo.GetByRoundNumber(ctx, roundNumber)

	if err == nil {
		record, err := repo.Update(ctx, roundNumber, data)
		if err != nil {
			return nil, err
		}
		return record, nil
	}

	_, ok := err.(f.NotFound)
	if !ok {
		return nil, err
	}

	record, err = repo.Create(ctx, roundNumber, data)
	if err != nil {
		return nil, err
	}
	return record, nil
}
