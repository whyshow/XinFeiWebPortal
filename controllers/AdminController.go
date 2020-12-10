package controllers

import (
	"../models"
	"../utils"
	"github.com/astaxie/beego"
	"log"
	"os/exec"
	"time"
)

type AdminController struct {
	beego.Controller
}

// 后台管理首页
func (c *AdminController) Home() {
	id := c.GetSession("admin_id")
	c.Data["info"] = models.QueryAdmin(id)
	c.TplName = "admin/后台首页页面.html"
}
func (c *AdminController) RedirectHome() {
	c.Redirect("/admin", 302)

}

// 基础信息统计显示页面
func (c *AdminController) Welcome() {
	t := time.Now()
	sy := models.GetSystemStatus()
	c.Data["UserCount"] = models.GetUserCount()
	c.Data["ArticleCount"] = models.GetArticleCount()

	c.Data["System"] = sy
	c.Data["time"] = utils.Millisecond(time.Since(t))
	c.TplName = "admin/系统信息显示页面.html"
}

// 获取系统状态信息
func (c *AdminController) GetSystemStatus() {
	t := time.Now()
	sy := models.GetSystemStatus()
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": sy}
	c.ServeJSON()
}

// 登录页面，判断是否数据库有admin信息，如果没有则跳转到注册页面
func (c *AdminController) LoginPage() {
	//

	c.TplName = "admin/后台登录页面.tpl"
}

// 管理员登录操作
func (c *AdminController) AdminLogin() {
	admin := models.Xinfei_Admin{Admin_name: c.GetString("Admin_name"), Admin_password: c.GetString("Admin_password")}
	models.AdminLogin(admin)
	if c.GetString("Admin_id") == "" && c.GetString("Admin_password") == "" {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "请输入内容", "result": nil}
		c.ServeJSON()
	}
	if bl, message, admin := models.AdminLogin(admin); bl {
		//加入session
		c.SetSession("admin_id", admin.Admin_id)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": message, "result": "/admin/"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": message, "result": nil}
		c.ServeJSON()
	}
}

// 退出登录
func (c *AdminController) ExitLogin() {
	c.DelSession("admin_id")
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "退出成功", "result": "/admin/login"}
	c.ServeJSON()

}

// 重新启动系统
// 测试中
func (c *AdminController) RebootSystem() {
	cmd := exec.Command("/bin/bash", "-c", "sudo kill -9 $(pidof XinFeiWebPortal)\ncd /www/wwwroot/Web/XinFeiWebPortal/\nnohup ./XinFeiWebPortal")
	_, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	c.TplName = "admin/后台首页页面.html"
}
