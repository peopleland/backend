package model

type User struct {
	FaunadbCommon `fauna:"-" json:"-"`

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
	FaunadbCommon `fauna:"-" json:"-"`

	Userid string `fauna:"userid,omitempty" json:"userid"`
	Code   string `fauna:"code,omitempty" json:"code"`
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

type OpenerGameRoundInfo struct {
	FaunadbCommon `fauna:"-" json:"-"`

	RoundNumber        int64  `fauna:"round_number" json:"round_number"`
	BuilderTokenAmount string `fauna:"builder_token_amount" json:"builder_token_amount"`
	EthAmount          string `fauna:"eth_amount" json:"eth_amount"`
	StartTimestamp     int64  `fauna:"start_timestamp" json:"start_timestamp"`
	EndTimestamp       int64  `fauna:"end_timestamp" json:"end_timestamp"`
	HasWinner          bool   `fauna:"has_winner" json:"has_winner"`
	WinnerTokenId      int64  `fauna:"winner_token_id" json:"winner_token_id"`
}
