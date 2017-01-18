package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Client 客户端信息
type Apps struct {
	Id         string
	App_name   string
	App_remark string
}

func NewApp(r *AddappsForm) (u *Apps, err error) {

	apps := Apps{

		App_name:   r.Appname,
		App_remark: r.Appremark,
	}

	return &apps, nil
}

func (u *Apps) Insert() {
	o := orm.NewOrm()
	o.Using("default")
	//_, err = o.Insert(user)
	create_time := time.Now().Format("2006-01-02 15:04:05")
	_, err := o.Raw("insert into oauth_modelid (app_name,app_remark,create_time,rewrite_time) values (?,?,?,?)", u.App_name, u.App_remark, create_time, create_time).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
}

func Apps_update(id, appname, appremark string) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	//_, err = o.Insert(user)
	rewrite_time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(appname)
	_, err = o.Raw("update oauth_modelid set app_name = ?,app_remark = ?,rewrite_time = ? where id = ?", appname, appremark, rewrite_time, id).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
	return
}

func Apps_delete(id string) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	_, err = o.Raw("delete from oauth_modelid where id = ?", id).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete apps ok")
	}
	return
}

func Applist() ([]Apps, error) {
	o := orm.NewOrm()
	o.Using("default")
	var apps []Apps
	_, err := o.Raw("select id,app_name,app_remark from oauth_modelid").QueryRows(&apps)

	return apps, err
}
