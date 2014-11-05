package routers

import (
	"admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//defult
	beego.Router("/", &controllers.IndexController{}, "*:Index")

	//admin
	beego.Router("/admin/login", &controllers.AdminController{}, "*:Login")
	beego.Router("/admin/logout", &controllers.AdminController{}, "*:Logout")
	beego.Router("/admin/register", &controllers.AdminController{}, "*:Register")
	beego.Router("/admin/activate", &controllers.AdminController{}, "get:Activate")
	beego.Router("/admin/getpwd", &controllers.AdminController{}, "*:Getpwd")
	beego.Router("/admin/list", &controllers.AdminController{}, "*:List")
	beego.Router("/admin/view", &controllers.AdminController{}, "*:View")
	beego.Router("/admin/lock", &controllers.AdminController{}, "*:Lock")
	beego.Router("/admin/unlock", &controllers.AdminController{}, "*:Unlock")

	//category
	beego.Router("/category/list", &controllers.CategoryController{}, "get:List")
	beego.Router("/category/create", &controllers.CategoryController{}, "*:Create")
}
