package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
}

func NewEnv() *Env {
	return new(Env)
}

func (e *Env) LoadEnvWithReplace(key, old, new string) {
	res := os.Getenv(key)
	res = strings.ReplaceAll(res, old, new)
	viper.SetDefault(key, res)
}

func (e *Env) LoadEnv(key string) {
	viper.SetDefault(key, os.Getenv(key))
}

func (e *Env) LoadFile(filePath string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filePath)
	if err := viper.MergeInConfig(); err != nil {
		fmt.Println(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (e *Env) Read(config interface{}) error {
	return viper.Unmarshal(config)
}
