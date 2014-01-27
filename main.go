package main

import (
	"desk/controllers"
	"github.com/astaxie/beego"
	"time"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wxapi", &controllers.APIController{})
	timefix := time.Now().Format("2006-01-02")
	logFileName := "logs/dev-" + timefix + ".log"
	beego.SetLogger("file", `{"filename":"`+logFileName+`"}`)
	beego.Router("/register", &controllers.RegisterController{})
	beego.Run()
}
