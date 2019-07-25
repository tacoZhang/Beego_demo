package main

import (
	"github.com/astaxie/beego"
	"my_blog/Db"
	_ "my_blog/routers"
)

func init() {
	Db.NewInit()      //Mysql数据库连接
	Db.RedisConnect() //Redis连接
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
