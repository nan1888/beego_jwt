package controllers

import (
	"jwt_demo/common"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*token := c.GetString("token")
	i := user_encode.Token_auth(token, "secret")
	if i == 1 {
		c.Data["json"] = user_encode.ErrExpired
		c.ServeJSON()
		return
	}*/
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
