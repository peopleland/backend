package conf

type Config struct {
	JwtRsaPrivateKeyPem string `mapstructure:"PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM"`
	JwtRsaPublicKeyPem  string `mapstructure:"PEOPLELAND_JWT_RSA_PUBLIC_KEY_PEM"`
	FaunaDBSecret       string `mapstructure:"PEOPLELAND_FAUNADB_SECRET"`
	Env                 string `mapstructure:"ENV"`
}
