package main

import (
	"pure-init/api/server"
	"pure-init/lib/config"
	"pure-init/lib/db"
	"pure-init/lib/log"
	"pure-init/lib/redis"
	_ "pure-init/docs"
)

// @title wukong
// @version 1.0

// @contact.name 赵晓
// @contact.email zhaoxiao248@jd.com

// @tag.name hello
// @tag.description 健康检查

func main() {
	log.InitZapLog()
	defer log.Sync()

	if err := config.LoadCofig(); err != nil {
		return
	}
	redisConf, _ := config.Config.Get("redis")
	redis.InitRedisWithConfig(redisConf)

	dbConf, _ := config.Config.Get("database")
	db.InitDBWithConfig(dbConf)

	server.StartHttp("0.0.0.0:8023")
}
