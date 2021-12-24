package models

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type User struct {
	Name    string `fauna:"name"`
	Address string `fauna:"address"`
}

type UserDb struct {
	ref  f.RefV `fauna:"ref"`
	ts   int64  `fauna:"ts"`
	data User   `fauna:"data"`
}

func CreateUser(client *f.FaunaClient, address string) (*UserDb, error) {
	result, err := client.Query(f.Create(f.Collection("users"), f.Obj{"data": User{Address: address}}))
	if err != nil {
		return nil, err
	}
	var userdb UserDb
	err1 := result.At(f.ObjKey("ts")).Get(&userdb.ts)
	if err1 != nil {
		return nil, err1
	}
	err2 := result.At(f.ObjKey("data")).Get(&userdb.data)
	if err2 != nil {
		return nil, err2
	}
	err3 := result.At(f.ObjKey("ref")).Get(&userdb.ref)
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
	err1 := result.At(f.ObjKey("ts")).Get(&userdb.ts)
	if err1 != nil {
		return nil, err1
	}
	err2 := result.At(f.ObjKey("data")).Get(&userdb.data)
	if err2 != nil {
		return nil, err2
	}
	err3 := result.At(f.ObjKey("ref")).Get(&userdb.ref)
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
