package env_config

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	JwtRsaPrivateKeyPem string `mapstructure:"PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM"`
	JwtRsaPublicKeyPem  string `mapstructure:"PEOPLELAND_JWT_RSA_PUBLIC_KEY_PEM"`
	FaunadbSecret       string `mapstructure:"PEOPLELAND_FAUNADB_SECRET"`
}

var FaunadbClient *f.FaunaClient

var Conf Config

func loadEnvWithReplace(key string, old string, new string) {
	res := os.Getenv(key)
	res = strings.ReplaceAll(res, old, new)
	viper.Set(key, res)
}

func loadEnv(key string) {
	viper.Set(key, os.Getenv(key))
}

func InitFaunadbClient() {
	FaunadbClient = f.NewFaunaClient(Conf.FaunadbSecret)
}

func BuildConfig() {
	loadEnvWithReplace("PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM", "\\n", "\n")
	loadEnvWithReplace("PEOPLELAND_JWT_RSA_PUBLIC_KEY_PEM", "\\n", "\n")
	loadEnv("PEOPLELAND_FAUNADB_SECRET")
	viper.Unmarshal(&Conf)
}
