package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User_token token message
type User_token struct {
	Token       string
	Express_in  string
	Appid       string
	Create_time string
}

/*
func (u *User_token) FindByID(name string) (result int, err error) {
	o := orm.NewOrm()
	o.Using("default")
	err = o.Raw("select * from user where username = ?", name).QueryRow(&u)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.Passwd)
		result = u.Id
	}
	return
}*/

func NewToken(r *CreateTokenForm, token string, express_in string) (u *User_token, err error) {
	regDate := time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		return nil, err
	}

	user := User_token{
		Appid:       r.Appid,
		Token:       token,
		Express_in:  express_in,
		Create_time: regDate,
	}

	return &user, nil
}

func (u *User_token) Insert() {
	o := orm.NewOrm()
	o.Using("default")

	fmt.Println("sssss")
	fmt.Println(u.Appid)
	//_, err = o.Insert(user)

	_, err := o.Raw("replace into oauth_token (id,token,create_time,express_in) values (?,?,?,?)", u.Appid, u.Token, u.Create_time, u.Express_in).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
}
