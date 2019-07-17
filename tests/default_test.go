package test

import (
	"fmt"
	"github.com/astaxie/beego"
	"my_blog/Db"
	_ "my_blog/routers"
	"my_blog/service"
	"path/filepath"
	"runtime"
	"testing"
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

func TestService2(t *testing.T) {
	es := service.NewEmpService()
	//emp:=models.Emp{
	//	Empno:7799,
	//	Ename:"Jack",
	//	Mgr:7788,
	//	Job:"CLERK",
	//	Hiredate:time.Now(),
	//	Sal:1200,
	//	Deptno:20,
	//}
	res := es.DeleteEmp(7799)
	fmt.Println(res)
}
