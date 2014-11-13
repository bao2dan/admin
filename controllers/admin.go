package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"

	"fmt"
	"html/template"
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
	}

	var rows []interface{}
	seHtml := `<label>
	                <input type="checkbox" class="ace" />
	                <span class="lbl"></span>
	            </label>`
	statusHtmlStr := `<span class="label label-sm %s">%s</span>`
	opHtmlStr := `<div class="visible-md visible-lg hidden-sm hidden-xs action-buttons">
	                <a class="blue" href="#">
	                    <i class="%s bigger-130"></i>
	                </a>
	                <a class="green" href="#">
	                    <i class="icon-pencil bigger-130"></i>
	                </a>
	                <a class="red" href="#">
	                    <i class="icon-trash bigger-130"></i>
	                </a>
	            </div>`

	for _, row := range list {
		lock, _ := row["lock"]
		status := "已激活"
		statusClass := "label-success"
		btnClass := "icon-unlock"
		if "1" == lock {
			status = "已锁定"
			statusClass = "label-warning"
			btnClass = "icon-lock"
		}
		statusHtml := template.HTML(fmt.Sprintf(statusHtmlStr, statusClass, status))
		opHtml := template.HTML(fmt.Sprintf(opHtmlStr, btnClass))
		line := []interface{}{seHtml, row["account"], row["role"], row["email"], row["create_time"], row["update_time"], row["login_time"], statusHtml, opHtml}
		rows = append(rows, line)
	}
	result["iTotalDisplayRecords"] = len(rows)
	result["iTotalRecords"] = len(rows)
	result["aaData"] = rows
	result["sEcho"] = 1
	result["succ"] = 1

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
