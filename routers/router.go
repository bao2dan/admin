package routers

import (
	"admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//defult
	beego.Router("/", &controllers.IndexController{}, "*:Index")

	//site
	beego.Router("/site/login", &controllers.SiteController{}, "*:Login")
	beego.Router("/site/logout", &controllers.SiteController{}, "*:Logout")
	beego.Router("/site/register", &controllers.SiteController{}, "*:Register")
	beego.Router("/site/activate", &controllers.SiteController{}, "get:Activate")

	//admin
	beego.Router("/admin/resetpasswd", &controllers.AdminController{}, "*:ResetPasswd")
	beego.Router("/admin/list", &controllers.AdminController{}, "*:List")
	beego.Router("/admin/view", &controllers.AdminController{}, "*:View")
	beego.Router("/admin/lock", &controllers.AdminController{}, "*:Lock")
	beego.Router("/admin/unlock", &controllers.AdminController{}, "*:Unlock")

	//category
	beego.Router("/category/list", &controllers.CategoryController{}, "get:List")
	beego.Router("/category/create", &controllers.CategoryController{}, "*:Create")
}
