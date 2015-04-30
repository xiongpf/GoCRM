package main

import (
	_ "CoCRM/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)
    orm.RegisterDataBase("default", "mysql", "root:ohana1011@/go_crm?charset=utf8")
}

func main() {
	beego.Run()
}

