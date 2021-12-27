package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/data/model"
	"context"
	"log"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type userRepo struct {
	data   *Data
	logger *log.Logger
}

func NewUserRepo(data *Data, logger *log.Logger) biz.UserRepo {
	return &userRepo{
		data:   data,
		logger: logger,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, address string) (*model.UserDb, error) {
	result, err := r.data.faunaClient.Query(f.Create(f.Collection("users"), f.Obj{"data": model.User{Address: address}}))
	if err != nil {
		return nil, err
	}
	var userdb model.UserDb
	err = result.At(f.ObjKey("ts")).Get(&userdb.Ts)
	if err != nil {
		return nil, err
	}
	err = result.At(f.ObjKey("data")).Get(&userdb.Data)
	if err != nil {
		return nil, err
	}
	err = result.At(f.ObjKey("ref")).Get(&userdb.Ref)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}

func (r *userRepo) GetOneUserByAddress(ctx context.Context, address string) (*model.UserDb, error) {
	result, err := r.data.faunaClient.Query(
		f.Get(
			f.MatchTerm(
				f.Index("users_by_address"),
				address,
			),
		))

	if err != nil {
		return nil, err
	}
	var userdb model.UserDb
	err = result.At(f.ObjKey("ts")).Get(&userdb.Ts)
	if err != nil {
		return nil, err
	}
	err = result.At(f.ObjKey("data")).Get(&userdb.Data)
	if err != nil {
		return nil, err
	}
	err = result.At(f.ObjKey("ref")).Get(&userdb.Ref)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}

func (r *userRepo) FindOrCreateUser(ctx context.Context, address string) (*model.UserDb, error) {
	user, _ := r.GetOneUserByAddress(ctx, address)
	if user != nil {
		return user, nil
	}

	newUser, err := r.CreateUser(ctx, address)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *userRepo) UpdateUserByAddress(ctx context.Context, address string, updateData map[string]interface{}) (*model.UserDb, error) {
	get := f.Get(
		f.MatchTerm(
			f.Index("users_by_address"),
			address,
		),
	)
	result, err := r.data.faunaClient.Query(
		f.Update(
			f.Select([]string{"ref"}, get),
			f.Obj{"data": updateData},
		),
	)
	if err != nil {
		return nil, err
	}
	var userdb model.UserDb
	err = result.At(f.ObjKey("ts")).Get(&userdb.Ts)
	if err != nil {
		return nil, err
	}
	err = result.At(f.ObjKey("data")).Get(&userdb.Data)
	if err != nil {
		return nil, err
	}
	err = result.At(f.ObjKey("ref")).Get(&userdb.Ref)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}
