package model

import (
	"fmt"
	"testing"

	"github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/stretchr/testify/assert"
)

func TestParseResult(t *testing.T) {
	fc := faunadb.NewFaunaClient("fnAEbfitSAACVKRgPF0ZYX-Q3zZiIE3jQpr_9km0")
	get := faunadb.Get(
		faunadb.MatchTerm(faunadb.Index(TelegramVerifyByUserIdIndex), "1112"))
	record, err := fc.Query(get)
	if err != nil {
		panic(err)
	}
	data := TelegramVerify{}
	err = ParseResult(record, &data)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, data.Ref.ID, "319400721764057684")
	assert.NotEmpty(t, data.Code)

	get, err = fc.Query(faunadb.Exists(faunadb.MatchTerm(faunadb.Index(TelegramVerifyByUserIdIndex), "111")))
	if err != nil {

	}
	fmt.Println(get)
}
