package main

import (
	"api/server"
	"fmt"
	"lib/config"
	"lib/db"
	"lib/log"
	"lib/redis"
)

func main() {

	if err := config.LoadCofig(); err != nil {
		return
	}
	redisConf, _ := config.Config.Get("redis")
	redis.InitRedisWithConfig(redisConf)
	dbConf, _ := config.Config.Get("database")
	if err := log.InitLogWithConfig("api"); err != nil {
		return
	}
	db.InitDBWithConfig(dbConf)
	fmt.Println("welcome use pure-init")
	server.StartHttp("0.0.0.0:8022")
	fmt.Println("8022")
}
