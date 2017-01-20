package action

import (
	"jwt_demo/common"
	"jwt_demo/models"

	"github.com/astaxie/beego"
)

type ReaduserlistController struct {
	beego.Controller
}

type ReadoneuserController struct {
	beego.Controller
}

func (c *ReaduserlistController) Get() {
	token := c.GetString("token")
	appid, err := user_encode.Token_auth(token, "secret")
	if err != nil {
		c.Data["json"] = user_encode.ErrExpired
		c.ServeJSON()
		return
	}
	_, err = models.Auth_role(user_encode.Select_all_user, appid)
	if err != nil {
		c.Data["json"] = user_encode.ErrPermission
		c.ServeJSON()
		return
	}
	users, err := models.Select_all()
	if err == nil {
		c.Data["json"] = users
		c.ServeJSON()
	} else {
		c.Data["json"] = user_encode.Err404
		c.ServeJSON()
	}

}

func (c *ReadoneuserController) Get() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}
