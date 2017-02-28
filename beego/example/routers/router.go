package routers

import (
	"github.com/astaxie/beego"
	"github.com/matthew/beego/example/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/directive", &controllers.DirectiveController{})
	beego.Router("/form", &controllers.FormController{})
}
