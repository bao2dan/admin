package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"

	"fmt"
	"html/template"
	//"strings"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Create() {
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplNames = "category/create.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	var err error
	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	this.Data["json"] = result
	this.ServeJson()
	return
}

//分类列表
func (this *CategoryController) List() {
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplNames = "category/list.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	var err error
	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	fileds := []string{"", "account", "role", "email", "create_time", "update_time", "login_time", "lock", ""}
	table := dateTableCondition(this.Ctx, fileds)

	rows := []interface{}{}
	list, count, err := models.AdminList(table)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		seHtml := `<label>
		                <input type="checkbox" class="ace" />
		                <span class="lbl"></span>
		            </label>`
		statusHtmlStr := `<span class="label label-sm status %s">%s</span>`
		opHtmlStr := `<div class="visible-md visible-lg hidden-sm hidden-xs action-buttons" account="%s">
			                <a class="green updateBtn" title="编辑" href="javascript:void(0);">
			                    <i class="icon-pencil bigger-130"></i>
			                </a>
			                <a class="blue unLockBtn" title="%s" href="javascript:void(0);">
			                    <i class="%s bigger-130"></i>
			                </a>
			                <a class="red delBtn" title="删除" href="javascript:void(0);">
			                    <i class="icon-trash bigger-130"></i>
			                </a>
			            </div>`

		for _, row := range list {
			lock, _ := row["lock"]
			status := "已激活"
			title := "锁定"
			statusClass := "label-success"
			btnClass := "icon-unlock"
			if "1" == lock {
				status = "已锁定"
				title = "解锁"
				statusClass = "label-warning"
				btnClass = "icon-lock"
			}
			statusHtml := template.HTML(fmt.Sprintf(statusHtmlStr, statusClass, status))
			opHtml := template.HTML(fmt.Sprintf(opHtmlStr, row["account"], title, btnClass))
			line := []interface{}{seHtml, row["account"], row["role"], row["email"], row["create_time"], row["update_time"], row["login_time"], statusHtml, opHtml}
			rows = append(rows, line)
		}
	}
	result["iTotalDisplayRecords"] = count
	result["iTotalRecords"] = count
	result["aaData"] = rows
	result["succ"] = 1

	this.Data["json"] = result
	this.ServeJson()
	return
}
