package config

import(
	"os"
)

type ConfigVar struct {
	Port string
	ApiKey string
}

func GetVariables()(c *ConfigVar){

	return &ConfigVar{
		Port : os.Getenv("PORT"),
		ApiKey : os.Getenv("AVIATION_STACK_API_KEY"),
	}

}