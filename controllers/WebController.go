package controllers

import (
	"github.com/astaxie/beego"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) Home() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "首页.tpl"
}
