package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	config     = new(Config)
	ConfigPath = "./"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Rate int    `yaml:"rate"`
	} `yaml:"server"`

	MySQL struct {
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`

	Redis struct {
		Addr         string `yaml:"addr"`
		Pass         string `yaml:"pass"`
		Db           int    `yaml:"db"`
		MaxRetries   int    `yaml:"maxRetries"`
		PoolSize     int    `yaml:"poolSize"`
		MinIdleConns int    `yaml:"minIdleConns"`
	} `yaml:"redis"`

	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`

	WX struct {
		AppID     string `yaml:"appID"`
		AppSecret string `yaml:"appSecret"`
	}
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
}

func Get() Config {
	return *config
}
