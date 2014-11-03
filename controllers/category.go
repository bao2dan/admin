package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) List() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.comcccccccccccccccccccccc"
	this.TplNames = "index.tpl"
}

func (this *CategoryController) Create() {
	this.Ctx.WriteString("xxxxxxxxxxxccccccccccccccccccc")
}
