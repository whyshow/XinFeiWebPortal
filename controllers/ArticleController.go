package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// 查询一篇文章
func (c *ArticleController) ArticleQueryOne() {
	t := time.Now()
	if article, err := models.ArticleSeleteOne(c.Ctx.Input.Param(":id")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 查询所有文章
func (c *ArticleController) ArticleQueryAll() {
	t := time.Now()
	if article, err := models.ArticleSeleteAll(); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

// 添加一篇文章
func (c *ArticleController) ArticleAddOne() {
	t := time.Now()
	article := models.Xinfei_article{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &article); err == nil {
		if err := models.ArticleInsertOne(article); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "错误", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "参数错误", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

//修改一篇文章
func (c *ArticleController) ArticleAlterOne() {

}

//删除一篇文章
func (c *ArticleController) ArticleDeleteOne() {
	t := time.Now()
	if _, err := models.ArticleDeleteOne(c.Ctx.Input.Param("id")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}
