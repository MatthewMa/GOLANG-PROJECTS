package routers

import (
	"github.com/matthew/beego/lessonsbasket/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
