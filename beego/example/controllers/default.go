package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type DirectiveController struct {
	beego.Controller
}

type FormController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *DirectiveController) Get() {
	c.TplName = "directives.html"
}

func (c *FormController) Get() {
	c.TplName = "form.html"
}
