package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultConfigPath     = "config"
	defaultConfigFileName = "application"
	defaultConfigFileType = "yaml"
	ymlTagName            = "mapstructure"
)

type (
	ServerConfig struct {
		Port  string   `mapstructure:"port"`
		Cosrs []string `mapstructure:"cors"`
	}

	ElasticConfig struct {
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	}

	Config struct {
		Server  ServerConfig  `mapstructure:"app-server"`
		Elastic ElasticConfig `mapstructure:"elastic"`
	}
)

func NewConfig() *Config {
	return ReadConfig(&Config{})
}

var ReadConfig = func(c *Config) *Config {
	fmt.Println("Konfigrasyon dosyası okunuyor...")
	v := viper.New()
	v = readFromLocalAppYml(v)

	if err := v.Unmarshal(c); err != nil {
		panic("Konfigrasyon dosyası okunurken hata alındı" + err.Error())
	}
	return c
}
