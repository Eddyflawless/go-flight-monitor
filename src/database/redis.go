package database

import (
	"github.com/go-redis/redis"
	config "github.com/Eddyflawless/flight-monitor/config"
	"log"
	"fmt"
)

var configVar *config.ConfigVar 
var DB *redis.Client


func New() (*redis.Client, error)  {

	configVar = config.GetVariables()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", configVar.RedisHost, configVar.RedisPort), 
		Password: configVar.RedisPassword, 
		DB: 0,//default db
	})

	pong, err := rdb.Ping().Result()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println(pong)

	return rdb, nil

}

