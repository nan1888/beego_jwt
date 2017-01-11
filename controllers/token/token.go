package token

import (
	"jwt_demo/common"
	"jwt_demo/models"
	//"fmt"
	//"reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type AccesstokenController struct {
	beego.Controller
}

type RefreshtokenController struct {
	beego.Controller
}

type Claims struct {
	Username string `json:"username"`
	// recommended having
	jwt.StandardClaims
}
type Usertoken struct {
	Token      string
	Appid      string
	Secret     string
	Express_in int64
}

func create_token(appid string, secret string) (token string) {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claims := Claims{
		appid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    appid,
		},
	}

	// Create the token using your claims
	c_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := c_token.SignedString([]byte(secret))

	return signedToken
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
	T.Express_in = time.Now().Add(time.Hour * 2).Unix()
	T.Token = create_token(form.Appid, form.Secret)
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
