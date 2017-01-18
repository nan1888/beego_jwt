package role

import (
	"jwt_demo/common"
	"jwt_demo/models"

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

type AddzoneController struct {
	beego.Controller
}

type DeletezoneController struct {
	beego.Controller
}

type ChangezoneController struct {
	beego.Controller
}
type ListzoneController struct {
	beego.Controller
}

func (c *AddroleController) Post() {

	form := models.AddroleForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	err := models.Role_insert(form.Rolename, form.Roleremark)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *ChangeroleController) Post() {

	form := models.ChangeroleForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	err := models.Role_update(form.Id, form.Appid, form.Zoneid)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *DeleteroleController) Post() {

	form := models.DeleteroleForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	err := models.Role_delete(form.Id)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *ListroleController) Get() {

	roles, err := models.Role_all()
	if err == nil {
		c.Data["json"] = roles
		c.ServeJSON()
	} else {
		c.Data["json"] = user_encode.Err404
		c.ServeJSON()
	}
}

func (c *AddzoneController) Post() {

	form := models.AddzoneForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	err := models.Zone_insert(form.Zonename, form.Zoneremark)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *ChangezoneController) Post() {
	form := models.ChangezoneForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	err := models.Zone_update(form.Id, form.Zonename, form.Zoneremark)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *DeletezoneController) Post() {

	form := models.DeletezoneForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	err := models.Zone_delete(form.Id)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *ListzoneController) Get() {

	zones, err := models.Zone_all()
	if err == nil {
		c.Data["json"] = zones
		c.ServeJSON()
	} else {
		c.Data["json"] = user_encode.Err404
		c.ServeJSON()
	}
}
