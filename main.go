package main

import (
	_ "chatapp/auth/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
}

func main() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:root@/chatapp?charset=utf8")
	fmt.Println("MySQL DataBase Connected")
	beego.Run("localhost")
}
