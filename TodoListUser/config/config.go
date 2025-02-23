package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Init() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("get work dir failed: %v", err))
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(workDir, "config", GetEnv()))

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config failed: %v", err))
	}
}

func GetEnv() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}

	return env
}
