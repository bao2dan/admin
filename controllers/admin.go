package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"

	"fmt"
	"html/template"
	"strings"
)

type AdminController struct {
	beego.Controller
}

//管理员账号列表
func (this *AdminController) List() {
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplNames = "admin/list.tpl"
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

	fileds := []string{"", "account", "role", "name", "phone", "email", "add_time", "update_time", "login_time", "lock", ""}
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
		statusHtmlStr := `<span class="label status %s">%s</span>`
		opHtmlStr := `<div class="action-buttons" account="%s" >
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
			line := []interface{}{seHtml, row["account"], row["role"], row["name"], row["phone"], row["add_time"], row["login_time"], statusHtml, opHtml}
			rows = append(rows, line)
		}
	}
	result["iTotalDisplayRecords"] = count
	result["iTotalRecords"] = count
	result["aaData"] = rows
	result["succ"] = 1
	result["msg"] = "成功"

	this.Data["json"] = result
	this.ServeJson()
	return
}

//修改管理员账号
func (this *AdminController) Update() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//获取账号和角色
	account := this.GetString("account")
	myaccount, _ := this.GetSession("account").(string)
	isAdmin := false
	if "" != account && account != myaccount {
		isAdmin = true
		//只能是超级管理员才能编辑其他管理员
		myrole, _ := this.GetSession("role").(string)
		if "root" != myrole {
			result["msg"] = "无权编辑其他管理员"
			this.Data["json"] = result
			this.ServeJson()
			return
		}
	} else {
		account = myaccount
	}

	//连接mongodb
	var err error
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	if !this.IsAjax() {
		//获取管理员信息
		info, err := models.GetAdmin(account)
		if nil != err {
			result["msg"] = err.Error()
			this.Data["json"] = result
			this.ServeJson()
			return
		}

		this.Data["IsAdmin"] = isAdmin
		this.Data["Role"] = strings.Split(info["role"], ",")
		this.Data["Info"] = info
		this.Layout = "layout.html"
		this.TplNames = "admin/update.tpl"
		this.Render()
		return
	}

	//获取参数并校验
	passwd := this.GetString("passwd")
	name := this.GetString("name")
	phone := this.GetString("phone")
	email := this.GetString("email")
	sex := this.GetString("sex")
	role := this.GetString("role")

	hasErr := false
	if "" == account || !isEmail(account) {
		result["msg"] = "账号有误"
		hasErr = true
	}
	if "" != passwd && !isPasswd(passwd) {
		result["msg"] = "密码有误"
		hasErr = true
	}
	if "" == name || len(name) < 2 || len(name) > 12 {
		result["msg"] = "姓名有误"
		hasErr = true
	}
	if "" == phone || !isPhone(phone) {
		result["msg"] = "手机号码有误"
		hasErr = true
	}
	if hasErr {
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	if "" != passwd {
		//md5 加密
		passwd = md5Encode(passwd + PASSWD_SECURITY)
	}

	err = models.UpdateAdmin(account, passwd, name, phone, email, sex, role, nowTime)
	result["succ"] = 1
	result["msg"] = "编辑成功"
	this.Data["json"] = result
	this.ServeJson()
	return

}

//删除管理员账号
func (this *AdminController) Del() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	account := this.GetString("account")
	if "" == account {
		result["msg"] = "参数不能为空"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//只能是超级管理员才能编辑其他管理员
	myrole, _ := this.GetSession("role").(string)
	if "root" != myrole {
		result["msg"] = "无权删除其他管理员"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	var err error
	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	err = models.DelAdmin(account)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["succ"] = 1
		result["msg"] = "删除成功"
	}

	this.Data["json"] = result
	this.ServeJson()
}

//查看管理员账号信息
func (this *AdminController) View() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	account, _ := this.GetSession("account").(string)

	//连接mongodb
	var err error
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	//获取管理员信息
	info, err := models.GetAdmin(account)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	this.Data["Role"] = strings.Split(info["role"], ",")
	this.Data["Info"] = info
	this.Layout = "layout.html"
	this.TplNames = "admin/view.tpl"
	this.Render()
}

//锁定管理员账号
func (this *AdminController) Lock() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	account := this.GetString("account")
	if "" == account {
		result["msg"] = "参数不能为空"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	var err error
	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	err = models.LockAdmin(account, nowTime)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["succ"] = 1
		result["msg"] = "锁定成功"
	}

	this.Data["json"] = result
	this.ServeJson()
}

//解锁(激活)管理员账号
func (this *AdminController) Unlock() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	account := this.GetString("account")
	if "" == account {
		result["msg"] = "参数不能为空"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	var err error
	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Data["json"] = err.Error()
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	err = models.UnlockAdmin(account, nowTime)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["succ"] = 1
		result["msg"] = "解锁成功"
	}

	this.Data["json"] = result
	this.ServeJson()
}
