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
}
