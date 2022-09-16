package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root1234@/cloud-drive?charset=utf8mb4")
	if err != nil {
		log.Printf("init xorm engine failed, err: %v", err)
		return nil
	}
	return engine

}
