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

var num = 10 //每页文章显示10行
/**
 * 文章列表首页
 */
func (c *ArticleController) ArticleHome() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, num, "DESC", false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, num, "DESC", false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "请求成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
		}
	}

	c.TplName = "admin/文章列表及管理页面.html"
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

//根据文章id删除一篇文章
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

//根据文章id修改一篇文章
func (c *ArticleController) ArticleAlterOne() {

}

// 查询文章详情内容
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

// 查询所有文章(按日期查询)
func (c *ArticleController) ArticleQueryAllDate() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, num, "DESC", false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, num, "DESC", false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

// 查询所有文章(按热度查询)
func (c *ArticleController) ArticleQueryAllHot() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, num, "DESC", true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, num, "DESC", true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, num, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

// 上下架一篇文章
//func (c * ArticleController) ArticleActivityOne() {
//	t := time.Now()
//
//	if mun,err:= models.ArticleActivityOne(c.Ctx.Input.Param(":id"));err == nil{
//		c.Data["json"] = map[string]interface{}{"code": mun, "message": "上架成功", "time": utils.Millisecond(time.Since(t)), "result": nil}
//		c.ServeJSON()
//	}else {
//		c.Data["json"] = map[string]interface{}{"code": mun, "message": "下架成功", "time": utils.Millisecond(time.Since(t)), "result": nil}
//		c.ServeJSON()
//	}
//
//}
