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
