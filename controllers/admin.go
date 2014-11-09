package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) List() {
	var err error
	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	info, err := models.Test("584143515@qq.com")
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	} else {
		this.Data["json"] = info
		this.ServeJson()
		return
	}

}

func (this *AdminController) Update() {
	role, _ := this.GetSession("role").(string)
	menu, err := getMenu(role)
	if nil != err {
		this.Ctx.WriteString(err.Error())
	}
	this.Data["Menu"] = menu
	this.Data["Version"] = menu
	this.Layout = "layout.html"
	this.TplNames = "index.tpl"
	this.Render()
}

func (this *AdminController) Del() {
	this.Ctx.WriteString("xxxxxxxxxxxccccccccccccccccccc")
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
