package config

import (
	"encoding/json"
	"github.com/spf13/viper"
)

type Config struct {
	DBUsername         string `json:"db_username"`
	DBPassword         string `json:"db_password"`
	DBHost             string `json:"db_host"`
	DBPort             int    `json:"db_port"`
	DBName             string `json:"db_name"`
	RedisHost          string `json:"redis_host"`
	RedisDB            int    `json:"redis_db"`
	AuthRedisKeyPrefix string `json:"auth_redis_key_prefix"`
	RedisKeyPrefix     string `json:"redis_key_prefix"`
	RedisPort          int    `json:"redis_port"`
	LogLevel           string `json:"log_level"`
}

var GlobalConfig *Config

func Load(env string) {
	conf := Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	m := viper.GetStringMap(env)
	jsonBody, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonBody, &conf); err != nil {
		panic(err)
	}

	GlobalConfig = &conf
}
