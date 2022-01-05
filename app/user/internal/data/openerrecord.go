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

type openerRecordRepo struct {
	data   *Data
	logger *log.Logger
}

func NewOpenerRecordRepo(data *Data, logger *log.Logger) biz.OpenerRecordRepo {
	return &openerRecordRepo{
		data:   data,
		logger: logger,
	}
}

func (repo *openerRecordRepo) GetOpenerRecordByTokenId(_ context.Context, tokenId int64) (*model.OpenerRecord, error) {
	result, err := repo.data.faunaClient.Query(
		f.Get(
			f.MatchTerm(
				f.Index(model.OpenerRecordByTokenId),
				tokenId,
			),
		))
	if err != nil {
		return nil, err
	}
	var record model.OpenerRecord
	err = model.ParseResult(result, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (repo *openerRecordRepo) CreateOpenerRecord(_ context.Context, tokenId int64, data *model.OpenerRecord) (*model.OpenerRecord, error) {
	data.TokenId = tokenId
	result, err := repo.data.faunaClient.Query(
		f.Create(f.Collection(model.OpenerRecordCollectionName),
			f.Obj{"data": data}))
	if err != nil {
		return nil, err
	}
	var record model.OpenerRecord
	err = model.ParseResult(result, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (repo *openerRecordRepo) UpdateOpenerRecord(_ context.Context, tokenId int64, data *model.OpenerRecord) (*model.OpenerRecord, error) {
	data.TokenId = tokenId
	result, err := repo.data.faunaClient.Query(
		f.Update(
			f.Select(
				[]string{"ref"},
				f.Get(
					f.MatchTerm(
						f.Index(model.OpenerRecordByTokenId),
						tokenId,
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
	var record model.OpenerRecord
	err = model.ParseResult(result, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (repo *openerRecordRepo) SaveOpenerRecord(ctx context.Context, tokenId int64, data *model.OpenerRecord) (*model.OpenerRecord, error) {
	record, err := repo.GetOpenerRecordByTokenId(ctx, tokenId)

	if err == nil {
		record, err := repo.UpdateOpenerRecord(ctx, tokenId, data)
		if err != nil {
			return nil, err
		}
		return record, nil
	}

	_, ok := err.(f.NotFound)
	if !ok {
		return nil, err
	}

	record, err = repo.CreateOpenerRecord(ctx, tokenId, data)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (repo *openerRecordRepo) GetNewest(ctx context.Context) (*model.OpenerRecord, error) {
	list, _, _, err := repo.getListPaginate(ctx, f.Size(1))
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

func (repo *openerRecordRepo) GetTotalCount(_ context.Context) (int64, error) {
	result, err := repo.data.faunaClient.Query(
		f.Count(
			f.Paginate(
				f.Match(f.Index(model.OpenerRecordSortTokenIdDescIndex)),
				f.Size(100_000),
			),
		),
	)

	if err != nil {
		return 0, err
	}

	var data f.ArrayV
	err = result.At(f.ObjKey("data")).Get(&data)
	if err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return 0, nil
	}

	var count int64
	err = data[0].Get(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetListPaginateFirstPage first page
func (repo *openerRecordRepo) GetListPaginateFirstPage(ctx context.Context, pageSize int64) (list []*model.OpenerRecord, afterTokenId int64, err error) {
	list, _, afterTokenId, err = repo.getListPaginate(ctx, f.Size(pageSize))
	return list, afterTokenId, err
}

// GetListPaginateAfter next page
func (repo *openerRecordRepo) GetListPaginateAfter(ctx context.Context, pageSize int64, inputAfterTokenId int64) (list []*model.OpenerRecord, beforeTokenId int64, afterTokenId int64, err error) {
	record, err := repo.GetOpenerRecordByTokenId(ctx, inputAfterTokenId)
	if err != nil {
		return nil, 0, 0, err
	}

	return repo.getListPaginate(
		ctx,
		f.Size(pageSize),
		f.After(
			f.Arr{
				inputAfterTokenId,
				f.Ref(f.Collection(model.OpenerRecordCollectionName), record.Ref.ID),
			}),
	)
}

// GetListPaginateBefore prev page
func (repo *openerRecordRepo) GetListPaginateBefore(ctx context.Context, pageSize int64, inputBeforeTokenId int64) (list []*model.OpenerRecord, beforeTokenId int64, afterTokenId int64, err error) {
	record, err := repo.GetOpenerRecordByTokenId(ctx, inputBeforeTokenId)
	if err != nil {
		return nil, 0, 0, err
	}

	return repo.getListPaginate(
		ctx,
		f.Size(pageSize),
		f.Before(f.Arr{
			inputBeforeTokenId,
			f.Ref(f.Collection(model.OpenerRecordCollectionName), record.Ref.ID),
		}),
	)
}

func (repo *openerRecordRepo) getListPaginate(_ context.Context, paginateOptions ...f.OptionalParameter) (list []*model.OpenerRecord, beforeTokenId int64, afterTokenId int64, err error) {
	result, err := repo.data.faunaClient.Query(
		f.Map(
			f.Paginate(
				f.Match(f.Index(model.OpenerRecordSortTokenIdDescIndex)),
				paginateOptions...,
			),
			f.Lambda(
				f.Arr{"token_id", "ref"},
				f.Get(f.Var("ref")),
			),
		),
	)

	if err != nil {
		return nil, 0, 0, err
	}

	var data f.ArrayV
	err = result.At(f.ObjKey("data")).Get(&data)
	if err != nil {
		return nil, 0, 0, err
	}

	resultList := make([]*model.OpenerRecord, 0)
	if len(data) == 0 {
		return resultList, 0, 0, nil
	}

	for _, item := range data {
		var record model.OpenerRecord
		err = model.ParseResult(item, &record)
		list = append(list, &record)
	}

	if err != nil {
		return nil, 0, 0, err
	}

	afterTokenId = 0
	var after f.ArrayV
	result.At(f.ObjKey("after")).Get(&after)
	if after != nil {
		_ = after[0].Get(&afterTokenId)
	}

	beforeTokenId = 0
	var before f.ArrayV
	result.At(f.ObjKey("before")).Get(&before)
	if before != nil {
		_ = before[0].Get(&beforeTokenId)
	}

	return list, beforeTokenId, afterTokenId, nil
}
