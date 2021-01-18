package controllers

import "github.com/astaxie/beego"

type HonorController struct {
	beego.Controller
}

// 获奖管理页面
func (c *HonorController) HonorHonePage() {
	c.TplName = "admin/获奖管理页面.html"
}
