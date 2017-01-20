package controllers

import (
	"fmt"
	"jwt_demo/common"
	"jwt_demo/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	token := c.GetString("token")
	appid, err := user_encode.Token_auth(token, "secret")
	if err != nil {

		c.Data["json"] = user_encode.ErrExpired
		c.ServeJSON()
		return
	}
	roleid, err := models.Auth_role("2321fd", appid)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = user_encode.ErrPermission
		c.ServeJSON()
		return
	}
	fmt.Println(roleid)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
