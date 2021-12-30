package model

type User struct {
	Name       string `fauna:"name" json:"name"`
	Address    string `fauna:"address" json:"address"`
	Twitter    string `fauna:"twitter" json:"twitter"`
	Telegram   string `fauna:"telegram" json:"telegram"`
	Discord    string `fauna:"discord" json:"discord"`
	VerifyCode string `fauna:"verify_code" json:"verify_code"`
}

type TelegramVerify struct {
	FaunadbCommon
	TelegramVerifyData
}

type TelegramVerifyData struct {
	Userid string `fauna:"userid" json:"userid"`
	Code   string `fauna:"code" json:"code"`
}
