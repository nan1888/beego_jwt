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
	Id         string
	Rolename   string
	Roleremark string
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

func Role_insert(rolename, roleremark string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into oauth_roleid (role_name,role_remark) values (?,?)", rolename, roleremark).Exec()

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

func Auth_role(appid, zoneid string) (int, error) {
	o := orm.NewOrm()
	o.Using("default")
	var roles []Role
	_, err := o.Raw("select id from oauth_roleid where zoneid = ? adn appid = ?", zoneid, appid).QueryRow(&roles)

	return roles, err
}
