package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {
	uname := this.GetSession("uname")

	if nil == uname {
		this.Redirect("/admin/login", 302)
	}

	this.Data["Version"] = "1.0"
	this.Layout = "layout.html"
	this.TplNames = "index.tpl"
	this.Render()
}
