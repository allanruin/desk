package routers

import (
	"desk/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wxapi", &controllers.APIController{})
	beego.Router("/register", &controllers.RegisterController{})
}
