package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type T struct {
	Website string `json:"Website"`
	Email   string `json:"Email"`
	Status  int    `json:"Status"`
}

func (c *MainController) Get() {
	var res T
	res = T{Website: "hia", Email: "hua", Status: 200}
	c.Data["json"] = res
	c.ServeJSON()
}
