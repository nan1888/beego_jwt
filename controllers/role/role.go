package role

import (
	"jwt_demo/common"

	"github.com/astaxie/beego"
)

type AddroleController struct {
	beego.Controller
}

type ChangeroleController struct {
	beego.Controller
}

type DeleteroleController struct {
	beego.Controller
}

type ListroleController struct {
	beego.Controller
}

func (c *AddroleController) Get() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}

func (c *ChangeroleController) Get() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}

func (c *DeleteroleController) Get() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}

func (c *ListroleController) Get() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}
