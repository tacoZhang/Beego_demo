/*
date:2019-07-16
author:Geek_张
*/
package Db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"sync"
)

var (
	RedisEngine *redis.Client
	Once        sync.Once
)

func RedisInit() {
	RedisConnect()
}

//Redis连接 <单例模式>
func RedisConnect() {
	Once.Do(func() {
		db, _ := beego.AppConfig.Int("Db")
		engine := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", beego.AppConfig.String("Host"), beego.AppConfig.String("Port")),
			DB:       db,
			Password: beego.AppConfig.String("Pwd"),
			PoolSize: 20,
		})
		if engine != nil {
			beego.Info("redis DB1 had connected")
		} else {
			beego.Error("redis start fail")
		}
		RedisEngine = engine
	})
}
