package models

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/scrypt"
)

// Client 客户端信息
type User struct {
	Username    string
	Passwd      string
	Salt        string
	Id          int
	Create_time string
	Appid       string
	Secret      string
	Role_id     int
	Email       string
}

const pwHashBytes = 64

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

//验证帐号密码

func (u *User) CheckPass(pass string) (ok bool, err error) {
	hash, err := generatePassHash(pass, u.Salt)
	if err != nil {
		return false, err
	}

	return u.Passwd == hash, nil
}

func (u *User) FindByID(name string) (result int, err error) {
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
}

func NewUser(r *RegisterForm, appid string, secret string) (u *User, err error) {
	salt, err := generateSalt()
	regDate := time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		return nil, err
	}
	hash, err := generatePassHash(r.Password, salt)
	if err != nil {
		return nil, err
	}

	user := User{
		Username:    r.Username,
		Email:       r.Email,
		Passwd:      hash,
		Salt:        salt,
		Create_time: regDate,
		Appid:       appid,
		Secret:      secret,
		Role_id:     1,
	}

	return &user, nil
}

func (u *User) Insert() {
	o := orm.NewOrm()
	o.Using("default")

	fmt.Println("sssss")
	fmt.Println(u.Appid)
	//_, err = o.Insert(user)

	_, err := o.Raw("insert into user (username,passwd,email,status,create_time,role_id,appid,secret,salt) values (?,?,?,?,?,?,?,?,?)", u.Username, u.Passwd, u.Email, 1, u.Create_time, u.Role_id, u.Appid, u.Secret, u.Salt).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
}

func Update_user_role(role_id, username string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("update user set role_id = ? where username = ?", role_id, username).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
	return err
}
