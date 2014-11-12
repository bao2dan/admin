package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) List() {
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplNames = "admin/list.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": "", "list": ""}
	account := this.GetString("account")

	var err error
	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	list, err := models.AdminList(account)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["list"] = list
	}
	this.Data["json"] = result
	this.ServeJson()
	return
}

func (this *AdminController) Update() {
	this.Data["Version"] = "1.1"
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
