package routers

import (
	"github.com/astaxie/beego"
	"github.com/matthew/beego/hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/upload", &controllers.UploadController{})
}
