package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Client 客户端信息
type Zone struct {
	Id         string
	Zonename   string
	Zoneremark string
}

type Role struct {
	Id      string
	App_id  string
	Zone_id string
}

func Zone_insert(zonename, zoneremark string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into oauth_zone (zone_name,zone_remark) values (?,?)", zonename, zoneremark).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
	return err
}

func Role_insert(appid, zoneid string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into oauth_roleid (app_id,zone_id) values (?,?)", appid, zoneid).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
	return err
}

func Zone_update(id, zonename, zoneremark string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("update oauth_zone set zone_name = ?,zone_remark = ?,rewrite_time = ? where id = ?", zonename, zoneremark, id).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
	return err
}

func Role_update(id, appid, zoneid string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("update oauth_roleid set app_id = ?,zone_id = ? where id = ?", appid, zoneid, id).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
	return err
}

func Role_delete(id string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("delete from oauth_roleid where id = ?", id).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete ok")
	}
	return err
}

func Zone_delete(id string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("delete form oauth_zone where id = ?", id).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete ok")
	}
	return err
}

func Zone_all() ([]Zone, error) {
	o := orm.NewOrm()
	o.Using("default")
	var zones []Zone
	_, err := o.Raw("select *from oauth_token").QueryRows(&zones)

	return zones, err
}

func Role_all() ([]Role, error) {
	o := orm.NewOrm()
	o.Using("default")
	var roles []Role
	_, err := o.Raw("select *from oauth_token").QueryRows(&roles)

	return roles, err
}

func Auth_role(app_name, appid string) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	var role Role
	err := o.Raw("select oauth_roleid.id  from oauth_roleid inner join oauth_modelid on (oauth_roleid.app_id=oauth_modelid.id) where oauth_modelid.app_name=? and oauth_roleid.zone_id = (select role_id from oauth_user where appid=?)", app_name, appid).QueryRow(&role)
	//fmt.Println(role.Id)
	return role.Id, err
}
