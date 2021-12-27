package model

import f "github.com/fauna/faunadb-go/v4/faunadb"

type UserDb struct {
	Ref  f.RefV `fauna:"ref" json:"ref"`
	Ts   int64  `fauna:"ts" json:"ts"`
	Data User   `fauna:"data" json:"data"`
}
