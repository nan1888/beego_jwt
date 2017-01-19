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
	token_status := user_encode.Token_auth(token, "secret")
	if token_status == 1 {
		c.Data["json"] = user_encode.ErrExpired
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
