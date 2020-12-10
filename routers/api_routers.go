package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	//API
	//beego.Router("/api/member", &controllers.UserController{}, "GET:MemberAll")                  // 查询成员
	//beego.Router("/api/member/:id", &controllers.UserController{}, "GET:UserQueryOne")           // 查询成员详情
	beego.Router("/api/article", &controllers.ApiController{}, "GET:QueryArticleList")        //查询文章列表（按日期）
	beego.Router("/api/article/hot", &controllers.ApiController{}, "GET:QueryArticleListHot") //查询文章列表（按热度值）
	beego.Router("/api/article/:id", &controllers.ApiController{}, "GET:QueryOneArticle")     // 查询文章详情
	beego.Router("/api/appbar", &controllers.ApiController{}, "GET:GetAppBar")                // 获取导航栏信息
}
