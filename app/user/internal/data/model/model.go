package model

type User struct {
	FaunadbCommon

	Name       string        `fauna:"name,omitempty" json:"name"`
	Address    string        `fauna:"address,omitempty" json:"address"`
	Twitter    string        `fauna:"twitter,omitempty" json:"twitter"`
	Telegram   *TelegramUser `fauna:"telegram,omitempty" json:"telegram"`
	Discord    *DiscordUser  `fauna:"discord,omitempty" json:"discord"`
	VerifyCode string        `fauna:"verify_code,omitempty" json:"verify_code"`
}

type DiscordUser struct {
	ID            string `fauna:"id,omitempty" json:"id,omitempty"`
	Username      string `fauna:"username,omitempty" json:"username,omitempty"`
	Discriminator string `fauna:"discriminator,omitempty" json:"discriminator,omitempty"`
	Avatar        string `fauna:"avatar,omitempty" json:"avatar,omitempty"`
}

type TelegramUser struct {
	ID           int64  `fauna:"id,omitempty" json:"id,omitempty"`
	FirstName    string `fauna:"first_name,omitempty" json:"first_name,omitempty"`
	Username     string `fauna:"username,omitempty" json:"username,omitempty"`
	LanguageCode string `fauna:"language_code,omitempty" json:"language_code,omitempty"`
}

type TelegramVerify struct {
	FaunadbCommon

	Userid string `fauna:"userid" json:"userid"`
	Code   string `fauna:"code" json:"code"`
}
