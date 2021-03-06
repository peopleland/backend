package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/data/model"
	"backend/app/user/pkg"
	"context"
	"errors"
	"log"
	"regexp"
	"strings"

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

func (r *userRepo) CreateUser(ctx context.Context, address string) (*model.User, error) {
	address = strings.ToLower(address)
	result, err := r.data.faunaClient.Query(f.Create(f.Collection(model.UserCollectionName), f.Obj{"data": map[string]string{
		"address": address,
	}}))
	if err != nil {
		return nil, err
	}
	var userdb model.User
	err = model.ParseResult(result, &userdb)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}

func (r *userRepo) GetUser(ctx context.Context, userid string) (*model.User, error) {
	result, err := r.data.faunaClient.Query(
		f.Get(
			f.Ref(
				f.Collection(model.UserCollectionName),
				userid,
			),
		))
	if err != nil {
		return nil, err
	}
	var userdb model.User
	err = model.ParseResult(result, &userdb)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}

func (r *userRepo) GetOneUserByAddress(ctx context.Context, address string) (*model.User, error) {
	address = strings.ToLower(address)
	result, err := r.data.faunaClient.Query(
		f.Get(
			f.MatchTerm(
				f.Index(model.UsersByAddressIndex),
				address,
			),
		))
	if err != nil {
		return nil, err
	}
	var userdb model.User
	err = model.ParseResult(result, &userdb)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}

func (r *userRepo) GetUserListByAddressList(_ context.Context, addressList []string) ([]*model.User, error) {
	arr := make([]f.Expr, 0)
	for _, address := range addressList {
		arr = append(arr, f.MatchTerm(f.Index(model.UsersByAddressIndex), strings.ToLower(address)))
	}

	if len(arr) == 0 {
		list := make([]*model.User, 0)
		return list, nil
	}

	result, err := r.data.faunaClient.Query(
		f.Map(
			f.Paginate(
				f.Union(arr),
				f.Size(len(arr)),
			),
			f.Lambda(
				"ref",
				f.Get(f.Var("ref")),
			),
		),
	)
	if err != nil {
		return nil, err
	}

	var data f.ArrayV
	err = result.At(f.ObjKey("data")).Get(&data)
	if err != nil {
		return nil, err
	}

	list := make([]*model.User, 0)
	if len(data) == 0 {
		return list, nil
	}

	for _, item := range data {
		var record model.User
		err = model.ParseResult(item, &record)
		list = append(list, &record)
	}

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *userRepo) FindOrCreateUser(ctx context.Context, address string) (*model.User, error) {
	address = strings.ToLower(address)
	user, err := r.GetOneUserByAddress(ctx, address)
	if user != nil {
		return user, nil
	}

	_, ok := err.(f.NotFound)
	if !ok {
		return nil, err
	}

	newUser, err := r.CreateUser(ctx, address)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, userid string, updateData map[string]interface{}) (*model.User, error) {
	result, err := r.data.faunaClient.Query(
		f.Update(
			f.Ref(
				f.Collection(model.UserCollectionName),
				userid,
			),
			f.Obj{"data": updateData},
		),
	)
	if err != nil {
		return nil, err
	}
	var userdb model.User
	err = model.ParseResult(result, &userdb)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}

func (r *userRepo) UpdateUserByAddress(ctx context.Context, address string, updateData map[string]interface{}) (*model.User, error) {
	address = strings.ToLower(address)
	get := f.Get(
		f.MatchTerm(
			f.Index(model.UsersByAddressIndex),
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
		err1, ok := err.(f.BadRequest)
		if ok {
			has, _ := regexp.MatchString("document is not unique", err1.Error())
			if has {
				return nil, errors.New("db.document.not_unique")
			}
		}
		return nil, err
	}
	var userdb model.User
	err = model.ParseResult(result, &userdb)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}

func (r *userRepo) CreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error) {
	data := model.TelegramVerify{
		Userid: userid,
		Code:   pkg.RandomString(10),
	}
	result, err := r.data.faunaClient.
		Query(
			f.Create(
				f.Collection(model.TelegramVerifyCollectionName),
				f.Obj{"data": data}))
	if err != nil {
		return nil, err
	}
	err = model.ParseResult(result, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *userRepo) GetOrCreateTelegramVerifyCode(ctx context.Context, userid string) (*model.TelegramVerify, error) {
	var existed bool
	get, err := r.data.faunaClient.Query(f.Exists(f.MatchTerm(f.Index(model.TelegramVerifyByUserIdIndex), userid)))
	if err != nil {
		return nil, err
	}
	err = get.Get(&existed)
	if err != nil {
		return nil, err
	}
	if existed {
		value, err := r.data.faunaClient.Query(f.Get(
			f.MatchTerm(f.Index(model.TelegramVerifyByUserIdIndex), userid)))
		if err != nil {
			return nil, err
		}
		var dbData model.TelegramVerify
		err = model.ParseResult(value, &dbData)
		if err != nil {
			return nil, err
		}
		return &dbData, err
	} else {
		return r.CreateTelegramVerifyCode(ctx, userid)
	}
}

func (r *userRepo) GetUserByTelegramVerifyCode(ctx context.Context, code string) (string, error) {
	value, err := r.data.faunaClient.Query(f.Get(
		f.MatchTerm(f.Index(model.TelegramVerifyByCodeIndex), code)))
	if err != nil {
		return "", err
	}
	var dbData model.TelegramVerify
	err = model.ParseResult(value, &dbData)
	if err != nil {
		return "", err
	}
	return dbData.Userid, err
}

func (r *userRepo) GenVerifyCode(ctx context.Context, userid string) (string, error) {
	user, err := r.GetUser(ctx, userid)
	if err != nil {
		return "", err
	}
	if user.VerifyCode != "" {
		return user.VerifyCode, nil
	}

	tryCount := 3
	for true {
		if tryCount <= 0 {
			break
		}
		tryCount -= 1
		newCode := pkg.RandomString(8)
		newUser, err := r.UpdateUser(ctx, userid, map[string]interface{}{
			"verify_code": newCode,
		})
		if err == nil {
			return newUser.VerifyCode, nil
		}
	}

	return "", errors.New("verify_code.gen.error")
}

func (r *userRepo) GetUserByVerifyCode(_ context.Context, verifyCode string) (*model.User, error) {
	result, err := r.data.faunaClient.Query(
		f.Get(
			f.MatchTerm(
				f.Index(model.UsersByVerifyCodeIndex),
				verifyCode,
			),
		))
	if err != nil {
		return nil, err
	}
	var userdb model.User
	err = model.ParseResult(result, &userdb)
	if err != nil {
		return nil, err
	}
	return &userdb, nil
}
