package model

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

const (
	UserCollectionName  = "users"
	UsersByAddressIndex = "users_by_address"

	TelegramVerifyCollectionName = "telegram_verify"
	TelegramVerifyByUserIdIndex  = "telegram_verify_by_userid"
)

type UserDb struct {
	Ref  f.RefV `fauna:"ref" json:"ref"`
	Ts   int64  `fauna:"ts" json:"ts"`
	Data User   `fauna:"data" json:"data"`
}

type FaunadbCommon struct {
	Ref f.RefV `fauna:"_" json:"ref"`
	Ts  int64  `fauna:"_" json:"ts"`
}

func ParseResult(result f.Value, meta *FaunadbCommon, data interface{}) error {
	err := result.At(f.ObjKey("ts")).Get(&meta.Ts)
	if err != nil {
		return err
	}
	err = result.At(f.ObjKey("ref")).Get(&meta.Ref)
	if err != nil {
		return err
	}
	err = result.At(f.ObjKey("data")).Get(data)
	if err != nil {
		return err
	}
	return nil
}
