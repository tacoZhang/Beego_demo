package routers

import (
	"github.com/astaxie/beego"
	"my_blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
