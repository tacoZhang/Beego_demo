/*
data: 2019-07-16
author: Geek_张
*/
package Db

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

var (
	MariaConf *gorm.DB
	OnceMutex sync.Once
)

func NewInit() {
	ConnectMariaDB()
}

//连接mysql <单例模式>
func ConnectMariaDB() {
	OnceMutex.Do(func() {
		beego.Info("connect to Mysql ", beego.AppConfig.String("dbHost"), beego.AppConfig.String("dbPort"))
		url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", //添加parseTime=true解决时间问题
			beego.AppConfig.String("dbUser"), beego.AppConfig.String("dbPassword"),
			beego.AppConfig.String("dbHost"), beego.AppConfig.String("dbPort"),
			beego.AppConfig.String("dbName"))
		db, err := gorm.Open("mysql", url)
		if err != nil {
			beego.Error("DataBase Connect Fail!", err.Error())
			return
		}
		db.SingularTable(true)
		db.DB().SetConnMaxLifetime(1 * time.Second)
		db.DB().SetMaxIdleConns(20)   //最大打开的连接数
		db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
		db.SingularTable(true)        //表生成结尾不带s
		// 启用Logger，显示详细日志
		res, _ := beego.AppConfig.Bool("dbSQL")
		db.LogMode(res)
		beego.Info("Mysql had connected")
		MariaConf = db
	})
}
