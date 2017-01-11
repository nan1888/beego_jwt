package routers

import (
	"jwt_demo/controllers"
	"jwt_demo/controllers/action"
	"jwt_demo/controllers/role"
	"jwt_demo/controllers/token"
	"jwt_demo/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	/*token*/
	beego.Router("/access_token", &token.AccesstokenController{})
	/*user*/
	beego.Router("/login", &user.LoginController{})
	beego.Router("/register", &user.RegisterController{})
	beego.Router("/findpasswd", &user.FindpasswdController{})
	beego.Router("/loginout", &user.LoginoutController{})
	beego.Router("/changepwsswd", &user.ChangepasswdController{})
	/*role*/
	beego.Router("/addrole", &role.AddroleController{})
	beego.Router("/deleterole", &role.DeleteroleController{})
	beego.Router("/changerole", &role.ChangeroleController{})
	beego.Router("/listrole", &role.ListroleController{})
	/*action*/
	beego.Router("/read_user_list", &action.ReaduserlistController{})
	beego.Router("/read_one_user", &action.ReadoneuserController{})
}
