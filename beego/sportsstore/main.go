package main

import (
	"github.com/astaxie/beego"
	_ "github.com/matthew/beego/sportsstore/routers"
)

func main() {
	beego.Run()
}
