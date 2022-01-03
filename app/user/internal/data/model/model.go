package model

type User struct {
	FaunadbCommon `fauna:"-" json:"-"`

	Name       string `fauna:"name,omitempty" json:"name"`
	Address    string `fauna:"address,omitempty" json:"address"`
	Twitter    string `fauna:"twitter,omitempty" json:"twitter"`
	Telegram   string `fauna:"telegram,omitempty" json:"telegram"`
	Discord    string `fauna:"discord,omitempty" json:"discord"`
	VerifyCode string `fauna:"verify_code,omitempty" json:"verify_code"`
}

type TelegramVerify struct {
	FaunadbCommon `fauna:"-" json:"-"`

	Userid string `fauna:"userid" json:"userid"`
	Code   string `fauna:"code" json:"code"`
}

type MintRecord struct {
	FaunadbCommon `fauna:"-" json:"-"`

	MintAddress  string `fauna:"mint_address" json:"mint_address"`
	X            string `fauna:"x" json:"x"`
	Y            string `fauna:"y" json:"y"`
	InviteUserid string `fauna:"invited_userid" json:"invited_userid"`
}

type OpenerRecord struct {
	FaunadbCommon `fauna:"-" json:"-"`

	MintAddress             string `fauna:"mint_address" json:"mint_address"`
	TokenId                 int64  `fauna:"token_id" json:"token_id"`
	X                       string `fauna:"x" json:"x"`
	Y                       string `fauna:"y" json:"y"`
	BlockNumber             int64  `fauna:"block_number" json:"block_number"`
	BlockTimestamp          int64  `fauna:"block_timestamp" json:"block_timestamp"`
	InvitedAddress          string `fauna:"invited_address" json:"invited_address"`
	NextTokenBlockTimestamp int64  `fauna:"next_token_block_timestamp" json:"next_token_block_timestamp"`
}
