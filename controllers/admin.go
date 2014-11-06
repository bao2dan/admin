package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) ResetPasswd() {
	this.Ctx.WriteString("xxxxxxxxxxxccccccccccccccccccc")
}

func (this *AdminController) List() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.comcccccccccccccccccccccc"
	this.Layout = "layout.html"
	this.TplNames = "user/login.tpl"
}

func (this *AdminController) View() {
	this.Ctx.WriteString("xxxxxxxxxxxccccccccccccccccccc")
}

func (this *AdminController) Lock() {
	this.Ctx.WriteString("xxxxxxxxxxxccccccccccccccccccc")
}

func (this *AdminController) Unlock() {
	this.Ctx.WriteString("xxxxxxxxxxxccccccccccccccccccc")
}
