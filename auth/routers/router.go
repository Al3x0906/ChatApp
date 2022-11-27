package routers

import (
	"chatapp/auth/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/", &controllers.UsersController{}, "get,post:Index")
	beego.Router("/login/", &controllers.LoginController{}, "get,post:Login")
}
