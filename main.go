package main

import (
	"./controllers"
	"./models"
	_ "./routers"
	"./utils"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

func init() {
	url := beego.AppConfig.String("jdbc.url")
	port := beego.AppConfig.String("jdbc.port")

	databasename := beego.AppConfig.String("jdbc.databasename")
	username := beego.AppConfig.String("jdbc.username")
	password := beego.AppConfig.String("jdbc.password")
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/"+databasename+"?charset=utf8")
	orm.RegisterModel(new(models.Xinfei_Admin))
	orm.RegisterModel(new(models.Xinfei_Day))
	orm.RegisterModel(new(models.Xinfei_user))
	orm.RegisterModel(new(models.Xinfei_article))
}

func main() {
	//Config()
	//注册新的错误设置 404/401/502
	beego.ErrorController(&controllers.ErrorController{})
	//注册过滤器  判断在用户访问admin后台时是否处于登录状态，如果没有登录，则跳转到 登录页面
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/*", beego.BeforeRouter, Record)
	//自动更新up pu量
	//models.TimerUpPUTask()
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}

//管理后台的登录过滤中间件
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("admin_id").(string)
	ok2 := strings.Contains(ctx.Request.RequestURI, "/admin/login")
	if !ok && !ok2 {
		ctx.Redirect(302, "/admin/login")
	}
}

var Record = func(ctx *context.Context) {
	if ok2 := strings.Contains(ctx.Request.RequestURI, "/admin/*"); !ok2 {
		//models.AddPv(1)
	}
}

func Config() {
	// 配置文件
	// 判断配置文件是否存在
	if is, _ := utils.PathExists("conf/init.ini"); !is {
		os.Create("../conf/init.ini")
	}
	if _, err := goconfig.LoadConfigFile("conf/init.ini"); err != nil {
		fmt.Println(err)
	} else {
		//cfg.SetValue("", "AdminLogin", "true")
		//goconfig.SaveConfigFile(cfg, "conf/init.ini")
	}
}
func ConfigINIT() *goconfig.ConfigFile {
	config, _ := goconfig.LoadConfigFile("conf/init.ini")
	return config
}