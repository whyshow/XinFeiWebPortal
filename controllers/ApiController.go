package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"time"
)

type ApiController struct {
	beego.Controller
}

//查询文章详情
func (c *ApiController) QueryOneArticle() {
	t := time.Now()
	if article, err := models.ArticleSeleteOne(c.Ctx.Input.Param(":id")); err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
		c.ServeJSON()
	}
}

//查询文章列表（按日期）
func (c *ApiController) QueryArticleList() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, 6, "DESC", false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, 6, "DESC", false); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

//查询文章列表（按热度值）
func (c *ApiController) QueryArticleListHot() {
	t := time.Now()
	if p, err := c.GetInt("p"); err == nil {
		if article, nums, err := models.ArticleSeleteAll(p, 6, "DESC", true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	} else {
		if article, nums, err := models.ArticleSeleteAll(1, 6, "DESC", true); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "time": utils.Millisecond(time.Since(t)), "result": article, "page": utils.Paginator(p, 6, nums)}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询失败", "time": utils.Millisecond(time.Since(t)), "result": err}
			c.ServeJSON()
		}
	}
}

var ab = models.AppBar{}

// 获取导航栏信息
func (c *ApiController) GetAppBar() {
	t := time.Now()
	if ab.Appbar == nil {
		var appb = models.AppBar{}
		func() {}()
		{

			f, err := os.Open("./conf/appbar.json")
			defer f.Close()
			if err != nil {
				return
			}
			bytes, _ := ioutil.ReadAll(f)
			err = json.Unmarshal(bytes, &appb)
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				return
			}
			ab = appb
		}
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "查询成功", "data": ab, "time": utils.Millisecond(time.Since(t))}
	c.ServeJSON()
}
