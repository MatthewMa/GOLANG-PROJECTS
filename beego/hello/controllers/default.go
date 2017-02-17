package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

type UploadController struct {
	beego.Controller
}

type ErrorController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *UserController) Get() {
	c.Ctx.WriteString("Hello,users!")
}

func (c *UploadController) Get() {
	c.TplName = "upload.html"
}

func (c *UploadController) Post() {
	username := c.GetString("username", "Matthew")
	password := c.GetString("password", "123666")
	_, h, err := c.GetFile("uploadfile")
	if err != nil {
		c.Ctx.WriteString("File submitting error!")
	} else {
		c.SaveToFile("uploadfile", "static/upload/"+h.Filename)
		str := "Username:" + username + ";Password:" + password + ";FileName:" + h.Filename
		c.Ctx.WriteString(str)
	}
}

func (c *ErrorController) Error404() {
	c.TplName = "Error.html"
}
