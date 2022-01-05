package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/data/model"
	"context"
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"log"
	"strings"
)

type mintRecordRepo struct {
	data   *Data
	logger *log.Logger
}

func NewMintRecordRepo(data *Data, logger *log.Logger) biz.MintRecordRepo {
	return &mintRecordRepo{
		data:   data,
		logger: logger,
	}
}

func (repo *mintRecordRepo) CreateMintRecord(_ context.Context, mintAddress string, x string, y string, userid string) (*model.MintRecord, error) {
	mintAddress = strings.ToLower(mintAddress)
	data := model.MintRecord{
		MintAddress:  mintAddress,
		X:            x,
		Y:            y,
		InviteUserid: userid,
	}
	result, err := repo.data.faunaClient.Query(
		f.Create(
			f.Collection(model.MintRecordCollectionName),
			f.Obj{"data": data},
		))
	if err != nil {
		return nil, err
	}
	var mintRecord model.MintRecord
	err = model.ParseResult(result, &mintRecord)
	if err != nil {
		return nil, err
	}
	return &mintRecord, nil
}

func (repo *mintRecordRepo) FindLastMintRecord(_ context.Context, mintAddress string, x string, y string, mintTimestamp int64) (*model.MintRecord, error) {
	mintAddress = strings.ToLower(mintAddress)
	result, err := repo.data.faunaClient.Query(
		f.Map(
			f.Paginate(
				f.Range(
					f.MatchTerm(
						f.Index(model.MintRecordsByMintAddressAndXAndYSortTsDescIndex),
						f.Arr{mintAddress, x, y},
					),
					f.Arr{mintTimestamp * 1000000},
					f.Arr{},
				),
				f.Size(1),
			),
			f.Lambda(
				f.Arr{"ts", "x", "y", "ref"},
				f.Get(f.Var("ref")),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	var data f.ArrayV
	_ = result.At(f.ObjKey("data")).Get(&data)
	if len(data) == 0 {
		return nil, nil
	}

	var mintRecord model.MintRecord
	err = model.ParseResult(data[0], &mintRecord)
	if err != nil {
		return nil, err
	}
	return &mintRecord, nil
}
