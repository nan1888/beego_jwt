package action

import (
	"jwt_demo/common"

	"github.com/astaxie/beego"
)

type ReaduserlistController struct {
	beego.Controller
}

type ReadoneuserController struct {
	beego.Controller
}

func (c *ReaduserlistController) Get() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}

func (c *ReadoneuserController) Get() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}
