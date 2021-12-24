package models

import (
	"fmt"
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	faunadbClient := f.NewFaunaClient("fnAEbMoNoqACTLHGcQsRsF5FqEL9YDyoNefP2ozf")
	address := "0x3946d96a4b46657ca95CBE85d8a60b822186Ad1f"
	user, _ := GetOneUserByAddress(faunadbClient, address)

	if user != nil {
		fmt.Println("have user")
		assert.Equal(t, user.data.Address, address)
	} else {
		fmt.Println("create user")
		user2, err2 := CreateUser(faunadbClient, address)
		if err2 != nil {
			t.Fail()
		}
		assert.Equal(t, user2.data.Address, address)
	}
}
