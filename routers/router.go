package routers

import (
	"jwt_demo/controllers"
	"jwt_demo/controllers/action"
	"jwt_demo/controllers/apps"
	"jwt_demo/controllers/role"
	"jwt_demo/controllers/token"
	"jwt_demo/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	/*token*/
	beego.Router("/access_token", &token.AccesstokenController{}) //获取token
	/*user*/
	beego.Router("/login", &user.LoginController{})               //用户登录
	beego.Router("/register", &user.RegisterController{})         //用户注册
	beego.Router("/findpasswd", &user.FindpasswdController{})     //找回密码
	beego.Router("/loginout", &user.LoginoutController{})         //用户退出登录
	beego.Router("/changepwsswd", &user.ChangepasswdController{}) //用户更改密码
	/*role*/
	beego.Router("/addrole", &role.AddroleController{})       //添加模块到用户组
	beego.Router("/deleterole", &role.DeleteroleController{}) //删除用户组的模块
	beego.Router("/changerole", &role.ChangeroleController{}) //改变用户的用户组
	beego.Router("/listrole", &role.ListroleController{})     //列出所有用户组的模块
	/*action*/
	beego.Router("/read_user_list", &action.ReaduserlistController{}) //查看所有用户的信息
	beego.Router("/read_one_user", &action.ReadoneuserController{})   //查看单独用户的信息
	/*zone*/
	beego.Router("/addzone", &role.AddzoneController{})       //添加用户组
	beego.Router("/deletezone", &role.DeletezoneController{}) //删除用户组
	beego.Router("/changezone", &role.ChangezoneController{}) //改变用户组
	beego.Router("/listzone", &role.ListzoneController{})     //列出所有用户组
	/*apps*/
	beego.Router("/addapps", &apps.AddappsController{})       //添加模块
	beego.Router("/deleteapps", &apps.DeleteappsController{}) //删除模块
	beego.Router("/changeapps", &apps.ChangeappsController{}) //改变模块的信息
	beego.Router("/listapps", &apps.ListappsController{})     //列出所有模块
}
