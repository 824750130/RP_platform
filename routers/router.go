package routers

import (
	"RP_platform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/user/register", &controllers.RegisterController{})
	beego.Router("/user/login", &controllers.LoginController{})
	beego.Router("/user/info", &controllers.UserController{})
}
