package controllers

import (
	"chatapp/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/ikeikeikeike/gopkg/convert"
)

type BaseController struct {
	beego.Controller

	Userinfo *models.User
	IsLogin  bool
	Session  session.Store
}

func (c *BaseController) Prepare() {
	c.SetParams()
	if c.Session == nil {
		c.Session = c.StartSession()
	}
	c.IsLogin = c.Session.Get("Userinfo") != nil
	if c.IsLogin {
		c.Userinfo = c.GetLogin()
	}

}

func (c *BaseController) GetLogin() *models.User {
	u := c.Session.Get("Userinfo").(*models.User)
	err := u.Read()
	if err != nil {
		return nil
	}
	return u
}

func (c *BaseController) DelLogin() {
	_ = c.Session.Delete("Userinfo")
}

func (c *BaseController) SetLogin(user *models.User) {
	_ = c.Session.Set("Userinfo", user)
	c.Userinfo = user
	c.IsLogin = true
}

func (c *BaseController) LoginPath() string {
	return c.URLFor("LoginController.Login")
}

func (c *BaseController) SetParams() {
	c.Data["Params"] = make(map[string]string)
	for k, v := range c.Input() {
		c.Data["Params"].(map[string]string)[k] = v[0]
	}
}

func (c *BaseController) BuildRequestUrl(uri string) string {
	if uri == "" {
		uri = c.Ctx.Input.URI()
	}
	return fmt.Sprintf("%s:%s%s",
		c.Ctx.Input.Site(), convert.ToStr(c.Ctx.Input.Port()), uri)
}
