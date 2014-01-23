package main

import (
	"desk/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wxapi", &controllers.APIController{})
	beego.Run()
}
