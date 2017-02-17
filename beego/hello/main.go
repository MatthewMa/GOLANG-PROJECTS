package main

import (
	"github.com/astaxie/beego"
	"github.com/matthew/beego/hello/controllers"
	_ "github.com/matthew/beego/hello/routers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
