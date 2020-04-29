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
	beego.Router("/member", &controllers.UserController{}, "GET:MemberAll")                                    //

	// 文章
	//beego.Router("article/member", &controllers.ArticleController{},"GET:ArticleQueryOne") //
	beego.Router("/article/:id", &controllers.ArticleController{}, "GET:ArticleQueryOne")
	beego.Router("/article/add", &controllers.ArticleController{}, "POST:ArticleAddOne")
	beego.Router("/article/all", &controllers.ArticleController{}, "GET:ArticleQueryAll")

	// 重启系统
	beego.Router("admin/reboot_system", &controllers.AdminController{}, "GET:RebootSystem")

	//API 返回JSON数据
	beego.Router("admin/getsystemstatus", &controllers.AdminController{}, "GET:GetSystemStatus")
}
