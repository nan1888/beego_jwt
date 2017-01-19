package main

import (
	_ "jwt_demo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:newpasswd@tcp(qmyk.shang-wifi.com:33306)/userrole?charset=utf8", 30, 30)
	orm.Debug = true
	beego.Run()
}
