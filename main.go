package main

import (
	"fmt"
	"pure-init/api/server"
	"pure-init/lib/config"
	"pure-init/lib/db"
	"pure-init/lib/log"
	"pure-init/lib/redis"
)

func main() {

	if err := config.LoadCofig(); err != nil {
		return
	}
	redisConf, _ := config.Config.Get("redis")
	redis.InitRedisWithConfig(redisConf)
	fmt.Println("redis is ok")
	dbConf, _ := config.Config.Get("database")
	if err := log.InitLogWithConfig("api"); err != nil {
		return
	}
	db.InitDBWithConfig(dbConf)
	fmt.Println("mysql is ok")
	fmt.Println("welcome use pure-init")
	server.StartHttp("0.0.0.0:8022")
	fmt.Println("8022")
}
