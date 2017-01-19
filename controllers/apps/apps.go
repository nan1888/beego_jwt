package apps

import (
	"fmt"
	"jwt_demo/common"
	"jwt_demo/models"

	"github.com/astaxie/beego"
)

type AddappsController struct {
	beego.Controller
}

type ChangeappsController struct {
	beego.Controller
}

type DeleteappsController struct {
	beego.Controller
}

type ListappsController struct {
	beego.Controller
}

func (c *AddappsController) Post() {
	form := models.AddappsForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	token_status := user_encode.Token_auth(form.Token, "secret")
	if token_status == 1 {
		c.Data["json"] = user_encode.ErrExpired
		c.ServeJSON()
		return
	}
	app, err := models.NewApp(&form)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	beego.Debug("NewUser:", app)
	app.Insert()
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *ChangeappsController) Post() {
	form := models.ChangeappsForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	token_status := user_encode.Token_auth(form.Token, "secret")
	if token_status == 1 {
		c.Data["json"] = user_encode.ErrExpired
		c.ServeJSON()
		return
	}
	fmt.Println(form.Appname)
	err := models.Apps_update(form.Id, form.Appname, form.Appremark)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *DeleteappsController) Post() {

	form := models.ChangeappsForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	token_status := user_encode.Token_auth(form.Token, "secret")
	if token_status == 1 {
		c.Data["json"] = user_encode.ErrExpired
		c.ServeJSON()
		return
	}
	err := models.Apps_delete(form.Id)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *ListappsController) Get() {
	token := c.GetString("token")
	token_status := user_encode.Token_auth(token, "secret")
	if token_status == 1 {
		c.Data["json"] = user_encode.ErrExpired
		c.ServeJSON()
		return
	}
	apps, err := models.Applist()
	if err == nil {
		c.Data["json"] = apps
		c.ServeJSON()
	} else {
		c.Data["json"] = user_encode.Actionsuccess
		c.ServeJSON()
	}
}
