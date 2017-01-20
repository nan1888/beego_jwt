package token

import (
	"jwt_demo/common"
	"jwt_demo/models"
	//"fmt"
	//"reflect"
	"strconv"

	"github.com/astaxie/beego"
)

type AccesstokenController struct {
	beego.Controller
}

type RefreshtokenController struct {
	beego.Controller
}

type Usertoken struct {
	Token      string
	Appid      string
	Secret     string
	Express_in int64
}

func (c *AccesstokenController) Get() {
	/*验证appid 和 secret，下发token*/
	form := models.CreateTokenForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("errParseRegsiterForm:", err)
		c.Data["json"] = user_encode.ErrInputData
		c.ServeJSON()
		return
	}
	var T Usertoken
	T.Token, T.Express_in = user_encode.Create_token(form.Appid, form.Secret)
	express_in := strconv.FormatInt(T.Express_in, 10)
	token_model, err := models.NewToken(&form, T.Token, express_in)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = user_encode.ErrSystem
		c.ServeJSON()
		return
	}
	beego.Debug("NewUser:", token_model)
	token_model.Insert()
	T.Appid = form.Appid
	T.Secret = form.Secret
	//fmt.Println("type:", reflect.TypeOf(express_in))
	c.Data["json"] = &T
	c.ServeJSON()
}
