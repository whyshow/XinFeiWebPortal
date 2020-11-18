package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	// Web
	beego.Router("/", &controllers.WebController{}, "GET:Home")

	// Admin
	beego.Router("admin/", &controllers.AdminController{}, "GET:Home")
	beego.Router("admin/welcome", &controllers.AdminController{}, "GET:Welcome")
	beego.Router("admin/login", &controllers.AdminController{}, "GET:LoginPage")
	beego.Router("manage/request_login", &controllers.AdminController{}, "POST:AdminLogin")
	beego.Router("admin/exit", &controllers.AdminController{}, "GET:ExitLogin")
	// 成员
	beego.Router("admin/user/home", &controllers.UserController{}, "GET:UserHome")                             // 用户的首页
	beego.Router("admin/user/page/add", &controllers.UserController{}, "GET:UserAddPage")                      // 添加用户页面
	beego.Router("admin/user/add", &controllers.UserController{}, "POST:UserAdd")                              // 单个添加用户
	beego.Router("admin/user/activate", &controllers.UserController{}, "GET:UserActivate")                     // 单个用户的激活或者禁用
	beego.Router("admin/user/delete", &controllers.UserController{}, "POST:UserDelete")                        // 单个用户的删除
	beego.Router("admin/user/page", &controllers.UserController{}, "GET:UserListPage")                         // 获取分页 (待用)
	beego.Router("admin/user/alter/page/password", &controllers.UserController{}, "GET:UserAlterPasswordPage") // 修改用户密码页面
	beego.Router("admin/user/alter/password", &controllers.UserController{}, "POST:UserAlterPassword")         // 修改用户密码
	beego.Router("admin/user/alter/page/user", &controllers.UserController{}, "GET:UserAlterInfoPage")         // 修改用户信息页面
	beego.Router("admin/user/alter/user", &controllers.UserController{}, "POST:UserAlterInfo")                 // 修改用户信息
	beego.Router("admin/user/query/user", &controllers.UserController{}, "POST:UserQuery")                     // 条件查询
	// 文章
	beego.Router("admin/article/home", &controllers.ArticleController{}, "GET:ArticleHome")
	beego.Router("admin/article/query", &controllers.ArticleController{}, "GET:ArticleQueryAllDate")
	beego.Router("article/add", &controllers.ArticleController{}, "POST:ArticleAddOne")

	beego.Router("/article/:id", &controllers.ArticleController{}, "GET:ArticleQueryOne")

	//API
	beego.Router("/member", &controllers.UserController{}, "GET:MemberAll")                  // 查询成员
	beego.Router("/member/:id", &controllers.UserController{}, "GET:UserQueryOne")           // 查询成员详情
	beego.Router("/article", &controllers.ArticleController{}, "GET:ArticleQueryAllDate")    //查询文章列表（按日期）
	beego.Router("/article/hot", &controllers.ArticleController{}, "GET:ArticleQueryAllHot") //查询文章列表（按热度值）
	beego.Router("/article/:id", &controllers.ArticleController{}, "GET:ArticleQueryOne")    // 查询文章详情

	beego.Router("admin/getsystemstatus", &controllers.AdminController{}, "GET:GetSystemStatus") //返回JSON数据
	beego.Router("admin/reboot_system", &controllers.AdminController{}, "GET:RebootSystem")      //重启系统

}
