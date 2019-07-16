package test

import (
	"fmt"
	"my_blog/Db"
	_ "my_blog/routers"
	"my_blog/service"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	Db.NewInit()      //Mysql数据库连接
	Db.RedisConnect() //Redis连接
}

func TestService(t *testing.T) {
	es := service.NewEmpService()
	list := es.SelectAll()
	fmt.Println(list)
	fmt.Println("------------------------------")
	ds := service.NewDeptService()
	result := ds.SelectAll()
	fmt.Println(result)
}
