package models

import (
	"cloud-drive/core/define"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", define.GetDatabaseInfo())
	if err != nil {
		log.Printf("init xorm engine failed, err: %v", err)
		return nil
	}
	return engine

}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     define.GetRedisInfo(), // RDB host
		Password: "",                    // no password set
		DB:       0,                     // use default DB
	})
}
