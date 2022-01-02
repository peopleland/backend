package model

import (
	"reflect"
	"sync"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

const (
	UserCollectionName     = "users"
	UsersByAddressIndex    = "users_by_address"
	UsersByVerifyCodeIndex = "users_by_verify_code"

	TelegramVerifyCollectionName = "telegram_verify"
	TelegramVerifyByUserIdIndex  = "telegram_verify_by_userid"

	MintRecordCollectionName                        = "mint_records"
	MintRecordsByMintAddressAndXAndYSortTsDescIndex = "mint_records_by_mint_address_and_x_and_y_sort_ts_desc"

	OpenerRecordCollectionName       = "opener_records"
	OpenerRecordSortTokenIdDescIndex = "opener_records_sort_token_id_desc"
	OpenerRecordByTokenId            = "opener_records_by_token_id"
)

type FaunadbCommon struct {
	Ref f.RefV `fauna:"-" json:"ref"`
	Ts  int64  `fauna:"-" json:"ts"`
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(FaunadbCommon)
	},
}

func ParseResult(result f.Value, data interface{}) error {
	common := pool.Get().(*FaunadbCommon)

	err := result.At(f.ObjKey("ts")).Get(&common.Ts)
	if err != nil {
		return err
	}
	err = result.At(f.ObjKey("ref")).Get(&common.Ref)
	if err != nil {
		return err
	}
	err = result.At(f.ObjKey("data")).Get(data)
	if err != nil {
		return err
	}

	reflect.ValueOf(data).Elem().FieldByName("Ref").Set(reflect.ValueOf(common.Ref))
	reflect.ValueOf(data).Elem().FieldByName("Ts").SetInt(common.Ts)

	pool.Put(common)
	return nil
}
