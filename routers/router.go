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
	beego.Router("/admin/del", &controllers.AdminController{}, "post:Del")
	beego.Router("/admin/lock", &controllers.AdminController{}, "post:Lock")
	beego.Router("/admin/unlock", &controllers.AdminController{}, "post:Unlock")
	beego.Router("/admin/view", &controllers.AdminController{}, "*:View")

	//category
	beego.Router("/category/list", &controllers.CategoryController{}, "*:List")
	beego.Router("/category/add", &controllers.CategoryController{}, "*:Add")
	beego.Router("/category/update", &controllers.CategoryController{}, "*:Update")
	beego.Router("/category/del", &controllers.CategoryController{}, "post:Del")

	//alimama
	beego.Router("/alimama/list", &controllers.AlimamaController{}, "*:List")
	beego.Router("/alimama/add", &controllers.AlimamaController{}, "*:Add")
	beego.Router("/alimama/update", &controllers.AlimamaController{}, "*:Update")
	beego.Router("/alimama/del", &controllers.AlimamaController{}, "post:Del")
}
