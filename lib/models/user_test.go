package models

import (
	"fmt"
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	faunadbClient := f.NewFaunaClient("fnAEbMoNoqACTLHGcQsRsF5FqEL9YDyoNefP2ozf")
	address := "0x3946d96a4b46657ca95CBE85d8a60b822186Ad1f"
	user, _ := GetOneUserByAddress(faunadbClient, address)

	if user != nil {
		fmt.Println("have user")
		assert.Equal(t, user.Data.Address, address)
	} else {
		fmt.Println("create user")
		user2, err2 := CreateUser(faunadbClient, address)
		if err2 != nil {
			t.Fail()
		}
		assert.Equal(t, user2.Data.Address, address)
	}
}

func TestUpdateUserByAddress(t *testing.T) {
	faunadbClient := f.NewFaunaClient("fnAEbMoNoqACTLHGcQsRsF5FqEL9YDyoNefP2ozf")
	address := "0x3946d96a4b46657ca95CBE85d8a60b822186Ad1f"

	name := "haha" + fmt.Sprint(time.Now().Unix())
	twitter := "yyy" + fmt.Sprint(time.Now().Unix())
	updateData := map[string]interface{}{
		"name":    name,
		"twitter": twitter,
	}
	user, err := UpdateUserByAddress(faunadbClient, address, updateData)
	if err != nil {
		t.Fail()
		return
	}
	assert.Equal(t, user.Data.Name, name)
	assert.Equal(t, user.Data.Twitter, twitter)
}
