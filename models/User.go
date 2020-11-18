package models

import (
	"../utils"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

/**
 * 用户表增删改查数据模型
 */
type Xinfei_user struct {
	User_account    string `orm:"column(user_account);pk"` //账号
	User_name       string
	User_password   string
	User_grade      string
	User_class      string
	User_direction  string
	User_permission int
	User_phone      string
	User_qq         string
	User_motto      string
	User_icon       string
	User_activate   int
	Register_date   string
}

// 根据年级查询行数
func UserGetAllGrade(grade string, p int, num int) ([]Xinfei_user, int64, error) {
	o := orm.NewOrm()
	var user []Xinfei_user
	sql := " "

	if p > 1 {
		sql = "SELECT * FROM xinfei_user WHERE user_grade =" + "'" + grade + "'" + " limit " + strconv.Itoa(p*num-num) + "," + strconv.Itoa(num)
	} else {
		sql = "SELECT * FROM xinfei_user WHERE user_grade =" + "'" + grade + "'" + " limit 0" + "," + strconv.Itoa(int(num))
	}

	// 返回行数或者错误信息
	if _, err := o.Raw(sql).QueryRows(&user); err == nil {
		b := []Xinfei_user{}
		for _, v := range user {
			v.User_password = ""
			b = append(b, v)
		}
		user = nil

		cnt, _ := o.QueryTable("xinfei_user").Filter("user_grade", grade).Count()
		return b, cnt, err
	} else {
		return user, 0, err
	}
	return nil, 0, nil
}

//按页码获取user表中的用户及信息 return 查询的数据，受影响行数，错误
func UserGetAllInfo(activate string, p int, num int, adc string) ([]Xinfei_user, int64, error) {
	o := orm.NewOrm()
	user := []Xinfei_user{}
	sql := ""
	if p == 1 && activate == "" || activate == " " {
		sql = "SELECT * FROM xinfei_user" + " ORDER BY user_grade " + adc + " limit 0" + "," + strconv.Itoa(num)
	} else if p > 1 && activate == "" || activate == " " {
		s := p*num - num
		sql = "SELECT * FROM xinfei_user" + " ORDER BY user_grade " + adc + " limit " + strconv.Itoa(s) + "," + strconv.Itoa(num)
	} else {
		sql = "SELECT * FROM xinfei_user" + " WHERE user_activate = " + "'" + activate + "'" + " ORDER BY user_grade " + adc
	}
	// 返回行数或者错误信息
	if num, err := o.Raw(sql).QueryRows(&user); err == nil {
		b := []Xinfei_user{}
		for _, v := range user {
			v.User_password = ""
			b = append(b, v)
		}
		user = nil
		if p == 0 {
			return b, num, err
		} else {
			cnt, _ := o.QueryTable("xinfei_user").Count()
			return b, cnt, err
		}
	} else {
		return user, 0, err
	}
}

//获取用户表中的用户数量
func GetUserCount() int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("xinfei_user").Count()
	return cnt
}

// 根据条件获取user表中单个user信息 return 查询的数据，错误
func UserGetInfoOne(unknown string) (Xinfei_user, error) {
	o := orm.NewOrm()
	user := Xinfei_user{}
	sql := "SELECT * FROM xinfei_user WHERE user_account = " + "'" + unknown + "'" + ""
	if err := o.Raw(sql).QueryRow(&user); err == nil {
		user.User_password = ""
		return user, err
	} else {
		return user, err
	}
}

// 根据条件查询所有的user信息 return 查询的数据，受影响行数，错误
func UserGetWhereAllInfo(unknown string) ([]Xinfei_user, int64, error) {
	o := orm.NewOrm()
	user := []Xinfei_user{}
	sql := "SELECT * FROM xinfei_user WHERE user_account = " + "'" + unknown + "'" + " OR  user_name = " + "'" + unknown + "'" + " OR  user_grade = " + "'" + unknown + "'" + " OR  user_class = " + "'" + unknown + "'" + " OR  user_direction = " + "'" + unknown + "'" + " OR  user_phone = " + "'" + unknown + "'" + " OR  user_qq = " + "'" + unknown + "'" + "" + " OR  user_motto = " + "'" + unknown + "'" + ""
	if numb, err := o.Raw(sql).QueryRows(&user); err == nil {
		b := []Xinfei_user{}
		for _, v := range user {
			v.User_password = ""
			b = append(b, v)
		}
		user = nil
		return b, numb, err
	} else {
		return user, numb, err
	}
}

// 激活/禁用 用户 return 受影响行数，错误
func UserActivate(account string, ac bool) (int64, error) {
	o := orm.NewOrm()
	user := Xinfei_user{User_account: account}
	if err := o.Read(&user); err == nil {
		activate := 0
		if ac {
			activate = 1
		}
		user.User_activate = activate
		if _, err := o.Update(&user); err == nil {
			return 1, err
		}
	} else {
		return 0, err
	}
	return 0, nil
}

// 根据账号删除单个用户 return 受影响行数，错误
func UserDeleteOne(account string) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Xinfei_user{User_account: account}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

// 根据用户账号更新(修改)用户密码 return 受影响行数，错误
func UserUpdatePassword(account string, password string) (int64, error) {
	o := orm.NewOrm()
	user := Xinfei_user{User_account: account}
	if err := o.Read(&user); err == nil {
		user.User_password = utils.MD5(password)
		if num, err := o.Update(&user); err == nil {
			return num, err
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}
}

// 插入单个用户 return 错误
func UserInsertOne(user Xinfei_user) error {
	o := orm.NewOrm()
	//补充数据 user_permission、user_activate
	user.User_permission = 0
	user.User_activate = 1
	user.User_password = utils.MD5(user.User_password)
	user.Register_date = time.Now().Format("2006-01-02")
	_, err := o.Insert(&user)
	if err == nil {
		return nil
	} else {
		return err
	}
}

// 修改单个用户信息 return 受影响行数，错误
func UserUpdateOne(user Xinfei_user) (int64, error) {
	o := orm.NewOrm()
	u := Xinfei_user{User_account: user.User_account}
	if err := o.Read(&u); err == nil {
		user.User_password = u.User_password
		user.User_activate = u.User_activate
		if num, err := o.Update(&user); err == nil {
			return num, err
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}
}
