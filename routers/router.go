package routers

import (
	"chatapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginController{}, "get:Login")
	beego.Router("/user/", &controllers.UsersController{}, "get,post:Index")
	beego.Router("/login/", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout/", &controllers.LoginController{}, "post:Logout")
	beego.Router("/signup/", &controllers.LoginController{}, "post:Signup")

	// WebSocket.
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
}
