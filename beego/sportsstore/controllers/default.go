package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type AdminController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *AdminController) Get() {
	c.TplName = "admin.html"
}
