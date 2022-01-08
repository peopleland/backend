package conf

type Config struct {
	Env string `mapstructure:"ENV"`

	JwtRsaPrivateKeyPem string `mapstructure:"PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM"`
	JwtRsaPublicKeyPem  string `mapstructure:"PEOPLELAND_JWT_RSA_PUBLIC_KEY_PEM"`

	FaunaDBSecret string `mapstructure:"PEOPLELAND_FAUNADB_SECRET"`

	TwitterConsumerKey    string `mapstructure:"PEOPLELAND_TWITTER_CONSUMER_KEY"`
	TwitterConsumerSecret string `mapstructure:"PEOPLELAND_TWITTER_CONSUMER_SECRET"`
	TwitterToken          string `mapstructure:"PEOPLELAND_TWITTER_TOKEN"`
	TwitterTokenSecret    string `mapstructure:"PEOPLELAND_TWITTER_TOKEN_SECRET"`

	DiscordBotClientID     string `mapstructure:"PEOPLELAND_DISCORD_BOT_CLIENT_ID"`
	DiscordBotClientSecret string `mapstructure:"PEOPLELAND_DISCORD_BOT_CLIENT_SECRET"`
	DiscordBotToken        string `mapstructure:"PEOPLELAND_DISCORD_BOT_TOKEN"`

	EthClientRawUrl           string `mapstructure:"PEOPLELAND_ETH_CLIENT_RAW_URL"`
	PeopleLandContractAddress string `mapstructure:"PEOPLELAND_CONTRACT_ADDRESS"`

	PeopleLandContractTheGraphApiUrl string `mapstructure:"PEOPLELAND_CONTRACT_THE_GRAPH_API_URL"`
}
