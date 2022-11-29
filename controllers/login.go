package controllers

import (
	"chatapp/lib"
	"chatapp/models"
	"encoding/json"
	"fmt"
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
	fmt.Println(c.Session.Get("Userinfo"))
	if c.IsLogin {
		c.Data["json"] = JsonResponse{true, reduce(c.Userinfo), 200, "OK"}
		c.ServeJSON()
		return
	}

	if !c.Ctx.Input.IsPost() {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: "No Last Login Found"}
		c.ServeJSON()
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
		c.ServeJSON()
		return
	}

	email := response.Email
	password := response.Password

	user, err := lib.Authenticate(email, password)
	if err != nil || user.Id < 1 {
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		c.ServeJSON()
		return
	}

	c.SetLogin(user)

	c.Data["json"] = JsonResponse{c.IsLogin, reduce(c.Userinfo), 200, "OK"}
	c.ServeJSON()

}

func (c *LoginController) Logout() {
	c.DelLogin()

	c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 200, Message: "OK"}
	c.ServeJSON()

}

func (c *LoginController) Signup() {
	type res struct {
		Username   string `json:"uname"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Repassword string `json:"repassword"`
	}

	var response res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &response)

	fmt.Println(string(c.Ctx.Input.RequestBody))

	u := &models.User{Username: response.Username, Email: response.Email, Password: response.Password, Repassword: response.Repassword}
	if err != nil {
		fmt.Println("0")
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		c.ServeJSON()
		return
	}
	if err = models.IsValid(u); err != nil {
		fmt.Println("1")
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		c.ServeJSON()
		return
	}

	id, err := lib.SignupUser(u)
	if err != nil || id < 1 {
		fmt.Println("2")
		c.Data["json"] = JsonResponse{IsLogin: false, Userinfo: nil, Status: 401, Message: err.Error()}
		c.ServeJSON()
		return
	}

	c.SetLogin(u)

	c.Data["json"] = JsonResponse{IsLogin: true, Userinfo: reduce(u), Status: 200, Message: "OK"}
	c.ServeJSON()
}
