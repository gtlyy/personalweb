package routers

import (
	"personalweb/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.HomeController{})
	web.Router("/article", &controllers.ArticleController{})
	web.Router("/article/:id", &controllers.ArticleController{}, "get:Detail")
	web.Router("/games", &controllers.GameController{})
	web.Router("/game/:id", &controllers.GameController{}, "get:Play")
	web.Router("/tools", &controllers.ToolController{})
	web.Router("/tool/:id", &controllers.ToolController{}, "get:Use")

	web.Router("/admin/login", &controllers.AdminController{}, "get:Login;post:DoLogin")
	web.Router("/admin/logout", &controllers.AdminController{}, "get:Logout")
	web.Router("/admin/index", &controllers.AdminController{}, "get:Index")
	web.Router("/admin/password", &controllers.AdminController{}, "get:PasswordPage;post:ChangePassword")

	web.Router("/admin/article/add", &controllers.AdminController{}, "get:Add;post:DoAdd")
	web.Router("/admin/article/edit/:id", &controllers.AdminController{}, "get:Edit;post:DoEdit")
	web.Router("/admin/article/del/:id", &controllers.AdminController{}, "get:Del")

	web.Router("/admin/game", &controllers.AdminController{}, "get:GameList")
	web.Router("/admin/game/add", &controllers.AdminController{}, "get:GameAdd;post:GameDoAdd")
	web.Router("/admin/game/edit/:id", &controllers.AdminController{}, "get:GameEdit;post:GameDoEdit")
	web.Router("/admin/game/del/:id", &controllers.AdminController{}, "get:GameDel")

	web.Router("/admin/tool", &controllers.AdminController{}, "get:ToolList")
	web.Router("/admin/tool/add", &controllers.AdminController{}, "get:ToolAdd;post:ToolDoAdd")
	web.Router("/admin/tool/edit/:id", &controllers.AdminController{}, "get:ToolEdit;post:ToolDoEdit")
	web.Router("/admin/tool/del/:id", &controllers.AdminController{}, "get:ToolDel")
}
