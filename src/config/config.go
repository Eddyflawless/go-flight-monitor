package config

import (
	"os"
)

type ConfigVar struct {
	Port          string
	ApiKey        string
	RedisPort     string
	RedisHost     string
	RedisPassword string
	ApiUrlLive    string
	ApiUrlDev     string
}

func GetVariables() (c *ConfigVar) {

	return &ConfigVar{
		Port:          os.Getenv("PORT"),
		ApiKey:        os.Getenv("AVIATION_STACK_API_KEY"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		ApiUrlLive:    os.Getenv("AVIATION_STACK_URL_LIVE"),
		ApiUrlDev:     os.Getenv("AVIATION_STACK_URL_DEV"),
	}

}
