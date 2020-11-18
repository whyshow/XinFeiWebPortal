package models

import (
	"../utils"
	"github.com/astaxie/beego/orm"
)

/**
 * 管理员数据库模型
 */
//管理员处理
type Xinfei_Admin struct {
	Admin_id         string `orm:"column(admin_id);pk"`
	Admin_name       string
	Admin_password   string
	Admin_permission int
	Admin_icon       string
}

//管理员登录
func AdminLogin(admin Xinfei_Admin) (bool, string, Xinfei_Admin) {
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM xinfei_admin WHERE admin_name = ? AND admin_password = ? ", admin.Admin_name, utils.MD5(admin.Admin_password)).QueryRow(&admin)
	if err == nil {
		return true, "登录成功", admin
	} else {
		if err := o.Raw("SELECT * FROM xinfei_admin WHERE admin_name = ? ", admin.Admin_name).QueryRow(&admin); err == nil {
			return false, "密码不正确", admin
		} else {
			return false, "账号不存在", admin
		}
	}
}

//查询管理员信息
func QueryAdmin(admin_id interface{}) Xinfei_Admin {
	o := orm.NewOrm()
	admin := Xinfei_Admin{}
	if admin_id == nil {
		if err := o.Raw("SELECT * FROM xinfei_admin WHERE admin_name = 'admin'").QueryRow(&admin); err == nil {
			return admin
		} else {
			return admin
		}
	} else {
		if err := o.Raw("SELECT xinfei_admin.admin_id,xinfei_admin.admin_name,xinfei_admin.admin_permission,xinfei_admin.admin_icon FROM xinfei_admin WHERE admin_id = ? ", admin_id).QueryRow(&admin); err == nil {
			return admin
		} else {
			return admin
		}
	}

}

//注册超级管理员
func AdminRegister() {

}
