package model

type User struct {
	FaunadbCommon

	Name       string `fauna:"name,omitempty" json:"name"`
	Address    string `fauna:"address,omitempty" json:"address"`
	Twitter    string `fauna:"twitter,omitempty" json:"twitter"`
	Telegram   string `fauna:"telegram,omitempty" json:"telegram"`
	Discord    string `fauna:"discord,omitempty" json:"discord"`
	VerifyCode string `fauna:"verify_code,omitempty" json:"verify_code"`
}

type TelegramVerify struct {
	FaunadbCommon

	Userid string `fauna:"userid" json:"userid"`
	Code   string `fauna:"code" json:"code"`
}
