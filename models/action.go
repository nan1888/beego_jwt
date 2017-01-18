package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//  客户端信息
type Action_user struct {
	Username    string
	Create_time string
	Appid       string
	Secret      string
	Role_id     int
	Email       string
}

func Select_all() (u []Action_user, err error) {
	o := orm.NewOrm()
	o.Using("default")
	var action_users []Action_user
	_, err = o.Raw("select username,create_time,appid,secrec,email from oauth_token").QueryRows(&action_users)

	return action_users, err
}

func (u *Action_user) Select_one(username string) (one_user *Action_user, err error) {
	o := orm.NewOrm()
	o.Using("default")
	err = o.Raw("select create_time,express_in from oauth_token where token = ?", username).QueryRow(&u)

	return u, err
}
