package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"

	"admin/models"

	"encoding/json"
	"fmt"
	"html/template"
	"strings"
)

type AlimamaController struct {
	beego.Controller
}

//管理员账号列表
func (this *AlimamaController) List() {
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplNames = "alimama/list.tpl"
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

	fileds := []string{"", "catid", "name", "oldPrice", "price", "url", "img", "addTime", "sort", "status", ""}
	table := dateTableCondition(this.Ctx, fileds)

	rows := []interface{}{}
	list, count, err := models.AlimamaList(table)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		seHtml := `<label>
		                <input type="checkbox" class="ace" />
		                <span class="lbl"></span>
		            </label>`
		statusHtmlStr := `<span class="%s">%s</span>`
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
			statusClass := "green"
			btnClass := "icon-unlock"
			if "1" == lock {
				status = "已锁定"
				title = "解锁"
				statusClass = "red"
				btnClass = "icon-lock"
			}
			statusHtml := template.HTML(fmt.Sprintf(statusHtmlStr, statusClass, status))
			opHtml := template.HTML(fmt.Sprintf(opHtmlStr, row["account"], title, btnClass))

			//角色值转角色名称
			r, _ := row["role"].(string)
			rs := strings.Split(r, ",")
			rnames := []string{}
			for _, v := range rolesKv {
				for rkey, rname := range v {
					if utils.InSlice(rkey, rs) {
						rnames = append(rnames, rname)
					}
				}
			}

			line := []interface{}{seHtml, row["account"], strings.Join(rnames, ","), row["name"], row["phone"], row["addTime"], row["loginTime"], statusHtml, opHtml}
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

//添加商品
func (this *AlimamaController) Add() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

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
		categoryTree, _ := models.CategoryTreeData("")
		jsonCateTree, _ := json.Marshal(categoryTree)
		this.Data["CategoryTree"] = string(jsonCateTree)
		this.Layout = "layout.html"
		this.TplNames = "alimama/add.tpl"
		this.Render()
		return
	}

	//获取参数并校验
	catid := this.GetString("catid")
	name := this.GetString("name")
	oldPrice := this.GetString("oldPrice")
	price := this.GetString("price")
	sort := this.GetString("sort")
	status := this.GetString("status")
	url := this.GetString("url")
	img := this.GetString("img")

	hasErr := false
	if "" == catid {
		result["msg"] = "分类ID有误"
		hasErr = true
	}
	if "" == name {
		result["msg"] = "名称有误"
		hasErr = true
	}
	if "" == oldPrice {
		result["msg"] = "原价有误"
		hasErr = true
	}
	if "" == price {
		result["msg"] = "现价有误"
		hasErr = true
	}
	if "" == sort {
		result["msg"] = "排序有误"
		hasErr = true
	}
	if "" == status {
		result["msg"] = "状态有误"
		hasErr = true
	}
	if "" == url {
		result["msg"] = "链接有误"
		hasErr = true
	}
	if "" == img {
		result["msg"] = "图片有误"
		hasErr = true
	}
	if hasErr {
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//添加分类
	err = models.AddAlimama(catid, name, oldPrice, price, sort, status, url, img, nowTime)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	result["succ"] = 1
	result["msg"] = "添加成功"
	this.Data["json"] = result
	this.ServeJson()
	return
}

//修改商品
func (this *AlimamaController) Update() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//商品ID
	id := this.GetString("id")
	if "" == id {
		result["msg"] = "参数不能为空"
		this.Data["json"] = result
		this.ServeJson()
		return
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

	//获取管理员信息
	info, err := models.GetAlimama(id)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	if !this.IsAjax() {
		//角色值转角色名称
		roleHtmlStr := `<label>
	                      <input name="role" value="%s" %s type="checkbox" class="ace ace-checkbox-2" />
	                      <span class="lbl">%s</span>
	                    </label>`
		roleHtml := ""
		role, _ := info["role"].(string)
		rs := strings.Split(role, ",")
		for _, v := range rolesKv {
			for rkey, rname := range v {
				if "root" == rkey {
					continue
				}
				if utils.InSlice(rkey, rs) {
					roleHtml += fmt.Sprintf(roleHtmlStr, rkey, "checked", rname)
				} else {
					roleHtml += fmt.Sprintf(roleHtmlStr, rkey, "", rname)
				}
			}
		}

		this.Data["RoleHtml"] = template.HTML(roleHtml)
		this.Data["Info"] = info
		this.Layout = "layout.html"
		this.TplNames = "alimama/update.tpl"
		this.Render()
		return
	}

	//获取参数并校验
	catid := this.GetString("catid")
	name := this.GetString("name")
	oldPrice := this.GetString("oldPrice")
	price := this.GetString("price")
	sort := this.GetString("sort")
	status := this.GetString("status")
	url := this.GetString("url")
	img := this.GetString("img")

	hasErr := false
	if "" == catid {
		result["msg"] = "分类ID有误"
		hasErr = true
	}
	if "" == name {
		result["msg"] = "名称有误"
		hasErr = true
	}
	if "" == oldPrice {
		result["msg"] = "原价有误"
		hasErr = true
	}
	if "" == price {
		result["msg"] = "现价有误"
		hasErr = true
	}
	if "" == sort {
		result["msg"] = "排序有误"
		hasErr = true
	}
	if "" == status {
		result["msg"] = "状态有误"
		hasErr = true
	}
	if "" == url {
		result["msg"] = "链接有误"
		hasErr = true
	}
	if "" == img {
		result["msg"] = "图片有误"
		hasErr = true
	}
	if hasErr {
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	err = models.UpdateAlimama(id, catid, name, oldPrice, price, sort, status, url, img, nowTime)
	result["succ"] = 1
	result["msg"] = "编辑成功"
	this.Data["json"] = result
	this.ServeJson()
	return
}

//删除管理员账号
func (this *AlimamaController) Del() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	id := this.GetString("id")
	if "" == id {
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

	err = models.DelAlimama(id)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["succ"] = 1
		result["msg"] = "删除成功"
	}

	this.Data["json"] = result
	this.ServeJson()
}

//查看商品信息
func (this *AlimamaController) View() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	id, _ := this.GetSession("id").(string)

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
	info, err := models.GetAlimama(id)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//角色值转角色名称
	roleHtmlStr := `<span class="lbl">%s </span>`
	roleHtml := ""
	role, _ := info["role"].(string)
	rs := strings.Split(role, ",")
	for _, v := range rolesKv {
		for rkey, rname := range v {
			if utils.InSlice(rkey, rs) {
				roleHtml += fmt.Sprintf(roleHtmlStr, rname)
			}
		}
	}

	this.Data["RoleHtml"] = template.HTML(roleHtml)
	this.Data["Info"] = info
	this.Layout = "layout.html"
	this.TplNames = "alimama/view.tpl"
	this.Render()
}

//下线商品
func (this *AlimamaController) Offline() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	id := this.GetString("id")
	if "" == id {
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

	err = models.OfflineAlimama(id, nowTime)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["succ"] = 1
		result["msg"] = "下线成功"
	}

	this.Data["json"] = result
	this.ServeJson()
}

//上线商品
func (this *AlimamaController) Online() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	id := this.GetString("id")
	if "" == id {
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

	err = models.OnlineAlimama(id, nowTime)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["succ"] = 1
		result["msg"] = "上线成功"
	}

	this.Data["json"] = result
	this.ServeJson()
}
