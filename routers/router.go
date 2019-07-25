package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"my_blog/controllers"
)

func init() {

	// 跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/emp",
			beego.NSInclude(
				&controllers.EmpController{},
			),
		),
		beego.NSNamespace("/dept",
			beego.NSInclude(&controllers.DeptController{}),
		),
	)
	beego.AddNamespace(ns)
}
