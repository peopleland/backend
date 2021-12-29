package model

import (
	"fmt"
	"testing"

	"github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/stretchr/testify/assert"
)

func TestParseResult(t *testing.T) {
	fc := faunadb.NewFaunaClient("fnAEbfjifeACVMELXa_tc0wdOe5SqgdXDdJd-zUR")
	get := faunadb.Get(
		faunadb.MatchTerm(faunadb.Index(TelegramVerifyByUserIdIndex), "111"))
	record, err := fc.Query(get)
	if err != nil {

	}
	data := TelegramVerifyData{}
	meta := FaunadbCommon{}
	err = ParseResult(record, &meta, &data)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, meta.Ref.ID, "319276042109846100")
	assert.Equal(t, data.Code, "222")

	get, err = fc.Query(faunadb.Exists(faunadb.MatchTerm(faunadb.Index(TelegramVerifyByUserIdIndex), "111")))
	if err != nil {

	}
	fmt.Println(get)

}
