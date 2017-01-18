package user

import (
	//	"encoding/hex"
	//"strings"
	"jwt_demo/common"
	"jwt_demo/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}
type LoginoutController struct {
	beego.Controller
}
type FindpasswdController struct {
	beego.Controller
}
type ChangepasswdController struct {
	beego.Controller
}
type ChangeuserroleController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	form := models.LoginForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	beego.Debug("ParseLoginForm:", &form)
	user := models.User{}
	if _, err := user.FindByID(form.Username); err != nil {
		beego.Error("FindUserById:", err)
		c.Data["json"] = user_encode.ErrNoUser
		c.ServeJSON()
		return
	}
	beego.Debug("UserInfo:", &user)
	if ok, err := user.CheckPass(form.Password); err != nil {
		beego.Error("CheckUserPass:", err)
		c.Data["json"] = user_encode.ErrPass
		c.ServeJSON()
		return
	} else if !ok {
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = &models.LoginInfo{Code: 0, UserInfo: &user}
	c.ServeJSON()
}

func (c *RegisterController) Post() {

	form := models.RegisterForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("errParseRegsiterForm:", err)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	beego.Debug("ParseRegsiterForm:", &form)
	appid := user_encode.To_md5(form.Username)
	secret := user_encode.To_md5(form.Email)
	user, err := models.NewUser(&form, appid, secret)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	beego.Debug("NewUser:", user)
	user.Insert()
	/*if _, err := user.Insert(); err != nil {
		beego.Error("InsertUser:", err)
		c.Data["json"] = "系统错误"
		c.ServeJSON()
	}*/
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}

func (c *LoginoutController) Post() {
	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}

func (c *FindpasswdController) Post() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}

func (c *ChangepasswdController) Post() {

	c.Data["json"] = user_encode.Err404
	c.ServeJSON()
}

func (c *ChangeuserroleController) Post() {
	form := models.ChangeuserroleForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		//c.Data["json"] = common.NewErrorInfo(ErrInputData)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	err := models.Update_user_role(form.Id, form.Username)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	c.Data["json"] = user_encode.Actionsuccess
	c.ServeJSON()
}
