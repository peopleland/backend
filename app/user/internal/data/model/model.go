package model

type User struct {
	Name    string `fauna:"name" json:"name"`
	Address string `fauna:"address" json:"address"`
	Twitter string `fauna:"twitter" json:"twitter"`
}
