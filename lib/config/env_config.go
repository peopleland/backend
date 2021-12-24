package env_config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	JwtRsaPrivateKeyPem string `mapstructure:"PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM"`
}

func loadEnvWithReplace(key string, old string, new string) {
	res := os.Getenv(key)
	res = strings.ReplaceAll(res, old, new)
	viper.Set(key, res)
}

func BuildConfig(config *Config) {
	loadEnvWithReplace("PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM", "\\n", "\n")
	viper.Unmarshal(config)
}
