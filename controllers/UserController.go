package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	beego.Controller
}

// 用户或者成员的首页
// 使用GET的形式进行请求 返回前端的数据包括json格式的用户列表，分页信息。
func (c *UserController) UserHome() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		//查询用户列表打包成json数据
		if user, num, err := models.UserGetAllInfo("", p, 6, ""); err == nil {
			//制作分页数据
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": user, "page": utils.Paginator(p, 6, num)}
		}
	} else {
		//查询用户列表打包成json数据
		if user, num, err := models.UserGetAllInfo("", 1, 6, ""); err == nil {
			//制作分页数据
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": user, "page": utils.Paginator(1, 6, num)}
			//显示前端页面

		}
	}
	c.TplName = "admin/member-list.html"
}

// 添加单个用户的页面
func (c *UserController) UserListPage() {

}

// 内容：添加用户
// 接收正确的json数据后会自动解析数据并添加到数据库中 返回json格式的数据
func (c *UserController) UserAdd() {
	t := time.Now()
	user := models.Xinfei_user{}
	//json数据封装到user对象中
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		if err := models.UserInsertOne(user); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "增加成功", "time": utils.Millisecond(time.Since(t)), "result": "/admin/member"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "增加失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "增加失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 内容：根据学号删除指定的人的信息
//
func (c *UserController) UserDelete() {
	t := time.Now()
	if num, err := models.UserDeleteOne(c.GetString("account")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": num, "message": "删除成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 添加单个用户的页面
func (c *UserController) UserAddPage() {
	c.TplName = "admin/member-add.html"
}

// 修改用户密码页面
func (c *UserController) UserAlterPasswordPage() {
	c.Data["json"] = map[string]interface{}{"code": 1, "message": c.GetString("u")}
	c.TplName = "admin/member-password.html"
}

// 修改用户密码
func (c *UserController) UserAlterPassword() {
	t := time.Now()
	if num, err := models.UserUpdatePassword(c.GetString("account"), c.GetString("password")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": num, "message": "修改成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 用户的账号激活或者禁用
func (c *UserController) UserActivate() {
	t := time.Now()
	activate, _ := c.GetBool("activate", false)
	account := c.GetString("account")
	if code, err := models.UserActivate(account, activate); code == 1 {
		c.Data["json"] = map[string]interface{}{"code": code, "message": "处理成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": code, "message": "处理失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

//修改用户信息页面
func (c *UserController) UserAlterInfoPage() {
	t := time.Now()
	if user, err := models.UserGetInfoOne(c.GetString("u")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
	}
	c.TplName = "admin/member-edit.html"
}

// 内容：修改用户信息
func (c *UserController) UserAlterInfo() {
	t := time.Now()
	user := models.Xinfei_user{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		if num, err := models.UserUpdateOne(user); err == nil {
			c.Data["json"] = map[string]interface{}{"code": num, "message": "修改成功", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": num, "message": "修改失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 内容：根据条件查询用户
func (c *UserController) UserQuery() {
	t := time.Now()
	if user, _, err := models.UserGetWhereAllInfo(c.GetString("u")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 内容：前端查询所有用户  返回 所有用户的数据（不包含密码）
func (c *UserController) MemberAll() {
	t := time.Now()
	activate := c.GetString("activate")
	desc := c.GetString("grade") //降序
	account := c.GetString("account")
	if account != "" {
		if user, err := models.UserGetInfoOne(account); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": user}
			c.ServeJSON()
		}
	} else if activate != "" {
		adc := "ASC"
		if desc == "desc" {
			adc = "DESC"
		} else if desc == "asc" {
			adc = "ASC"
		}
		if user, num, err := models.UserGetAllInfo(activate, 0, 0, adc); err == nil {
			if num > 1 {
				c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": user}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"code": 0, "message": "未查询到", "time": utils.Millisecond(time.Since(t)), "result": user}
				c.ServeJSON()
			}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "参数错误", "time": utils.Millisecond(time.Since(t)), "result": nil}
		c.ServeJSON()
	}
}
