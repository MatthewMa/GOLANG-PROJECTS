package routers

import (
	"github.com/astaxie/beego"
	"github.com/matthew/beego/sportsstore/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
}
