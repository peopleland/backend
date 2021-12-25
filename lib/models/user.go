package models

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type User struct {
	Name    string `fauna:"name" json:"name"`
	Address string `fauna:"address" json:"address"`
	Twitter string `fauna:"twitter" json:"twitter"`
}

type UserDb struct {
	Ref  f.RefV `fauna:"ref" json:"ref"`
	Ts   int64  `fauna:"ts" json:"ts"`
	Data User   `fauna:"data" json:"data"`
}

func CreateUser(client *f.FaunaClient, address string) (*UserDb, error) {
	result, err := client.Query(f.Create(f.Collection("users"), f.Obj{"data": User{Address: address}}))
	if err != nil {
		return nil, err
	}
	var userdb UserDb
	err1 := result.At(f.ObjKey("ts")).Get(&userdb.Ts)
	if err1 != nil {
		return nil, err1
	}
	err2 := result.At(f.ObjKey("data")).Get(&userdb.Data)
	if err2 != nil {
		return nil, err2
	}
	err3 := result.At(f.ObjKey("ref")).Get(&userdb.Ref)
	if err3 != nil {
		return nil, err1
	}
	return &userdb, nil
}

func GetOneUserByAddress(client *f.FaunaClient, address string) (*UserDb, error) {
	result, err := client.Query(
		f.Get(
			f.MatchTerm(
				f.Index("users_by_address"),
				address,
			),
		))

	if err != nil {
		return nil, err
	}
	var userdb UserDb
	err1 := result.At(f.ObjKey("ts")).Get(&userdb.Ts)
	if err1 != nil {
		return nil, err1
	}
	err2 := result.At(f.ObjKey("data")).Get(&userdb.Data)
	if err2 != nil {
		return nil, err2
	}
	err3 := result.At(f.ObjKey("ref")).Get(&userdb.Ref)
	if err3 != nil {
		return nil, err1
	}
	return &userdb, nil
}

func FindOrCreateUser(client *f.FaunaClient, address string) (*UserDb, error) {
	user, _ := GetOneUserByAddress(client, address)
	if user != nil {
		return user, nil
	}

	newUser, err := CreateUser(client, address)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func UpdateUserByAddress(client *f.FaunaClient, address string, updateData map[string]interface{}) (*UserDb, error) {
	get := f.Get(
		f.MatchTerm(
			f.Index("users_by_address"),
			address,
		),
	)
	result, err := client.Query(
		f.Update(
			f.Select([]string{"ref"}, get),
			f.Obj{"data": updateData},
		),
	)
	if err != nil {
		return nil, err
	}
	var userdb UserDb
	err1 := result.At(f.ObjKey("ts")).Get(&userdb.Ts)
	if err1 != nil {
		return nil, err1
	}
	err2 := result.At(f.ObjKey("data")).Get(&userdb.Data)
	if err2 != nil {
		return nil, err2
	}
	err3 := result.At(f.ObjKey("ref")).Get(&userdb.Ref)
	if err3 != nil {
		return nil, err1
	}
	return &userdb, nil
}
