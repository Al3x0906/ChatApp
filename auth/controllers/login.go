package controllers

import (
	"chatapp/auth/lib"
	"chatapp/auth/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

type LoginUser struct {
	IsLogin  bool         `json:"IsLogin"`
	Userinfo *models.User `json:"Userinfo"`
	Status   int          `json:"Status"`
}

func (c *LoginController) Login() {

	if c.IsLogin {
		c.Data["json"] = LoginUser{IsLogin: c.IsLogin, Userinfo: c.Userinfo, Status: 200}
		c.ServeJSON()
		return
	}

	c.Data["json"] = LoginUser{IsLogin: false, Userinfo: nil, Status: 401}

	if !c.Ctx.Input.IsPost() {
		c.ServeJSON()
		return
	}

	fmt.Println(c)
	email := c.GetString("Email")
	password := c.GetString("Password")

	user, err := lib.Authenticate(email, password)
	if err != nil || user.Id < 1 {
		c.Data["json"] = LoginUser{IsLogin: false, Userinfo: nil, Status: 401}
		c.ServeJSON()
		return
	}

	c.SetLogin(user)

	c.Data["json"] = LoginUser{IsLogin: c.IsLogin, Userinfo: c.Userinfo, Status: 200}
	c.ServeJSON()

}

func (c *LoginController) Logout() {
	c.DelLogin()

	c.Data["json"] = LoginUser{IsLogin: false, Userinfo: nil, Status: 200}
	c.ServeJSON()

}

func (c *LoginController) Signup() {
	fmt.Println(c.Data)

	var err error
	flash := beego.NewFlash()

	u := &models.User{}
	if err = c.ParseForm(u); err != nil {
		flash.Error("Signup invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(u); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	id, err := lib.SignupUser(u)
	if err != nil || id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Register user")
	flash.Store(&c.Controller)

	c.SetLogin(u)

	c.Redirect(c.URLFor("UsersController.Index"), 303)
}
