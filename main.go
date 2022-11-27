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
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/chatapp?charset=utf8")
	fmt.Println("MySQL DataBase Connected")
	beego.Run("localhost")
}
