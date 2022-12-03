package controllers

import (
	"chatapp/lib"
	"chatapp/models"
	"encoding/json"
)

type LoginController struct {
	BaseController
}

type LoginUser struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type JsonResponse struct {
	IsLogin  bool       `json:"isLogin"`
	Userinfo *LoginUser `json:"userinfo"`
	Status   int        `json:"status"`
	Message  string     `json:"message"`
}

func reduce(user *models.User) *LoginUser {
	return &LoginUser{Id: user.Id, Username: user.Username, Email: user.Email}
}

func (c *LoginController) Login() {
	defer c.ServeJSON()
	if c.IsLogin {
		c.Data["json"] = JsonResponse{true, reduce(c.Userinfo), 200, "OK"}
		return
	}

	if !c.Ctx.Input.IsPost() {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: "No Last Login Found"}
		return
	}

	type res struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var response res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &response)
	if !(err == nil) {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		return
	}

	email := response.Email
	password := response.Password

	user, err := lib.Authenticate(email, password)
	if err != nil || user.Id < 1 {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		return
	}

	c.SetLogin(user)

	c.Data["json"] = JsonResponse{c.IsLogin, reduce(c.Userinfo), 200, "OK"}

}

func (c *LoginController) Logout() {
	c.DelLogin()

	c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 200, Message: "OK"}
	c.ServeJSON()

}

type res struct {
	Username   string `json:"uname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

func (c *LoginController) Signup() {
	defer c.ServeJSON()

	var response res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &response)

	u := &models.User{Username: response.Username, Email: response.Email, Password: response.Password, Repassword: response.Repassword}
	if err != nil {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		return
	}
	if err = models.IsValid(u); err != nil {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		return
	}

	id, err := lib.SignupUser(u)
	if err != nil || id < 1 {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		return
	}

	c.SetLogin(u)

	c.Data["json"] = JsonResponse{IsLogin: true, Userinfo: reduce(u), Status: 200, Message: "OK"}
}
