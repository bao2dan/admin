package routers

import (
	"admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//defult
	beego.Router("/", &controllers.SiteController{}, "*:Index")

	//site
	beego.Router("/site/login", &controllers.SiteController{}, "*:Login")
	beego.Router("/site/logout", &controllers.SiteController{}, "*:Logout")
	beego.Router("/site/register", &controllers.SiteController{}, "*:Register")
	beego.Router("/site/activate", &controllers.SiteController{}, "get:Activate")
	beego.Router("/site/noauth", &controllers.SiteController{}, "get:NoAuth")
	beego.Router("/site/menu", &controllers.SiteController{}, "*:Menu")

	//admin
	beego.Router("/admin/list", &controllers.AdminController{}, "*:List")
	beego.Router("/admin/update", &controllers.AdminController{}, "*:Update")
	beego.Router("/admin/del", &controllers.AdminController{}, "*:Del")
	beego.Router("/admin/lock", &controllers.AdminController{}, "*:Lock")
	beego.Router("/admin/unlock", &controllers.AdminController{}, "*:Unlock")

	//category
	beego.Router("/category/list", &controllers.CategoryController{}, "get:List")
	beego.Router("/category/create", &controllers.CategoryController{}, "*:Create")

	beego.Router("/admin/up", &controllers.AdminController{}, "*:Up")
}
