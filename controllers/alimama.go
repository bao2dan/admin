package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"

	"admin/models"

	"fmt"
	"html/template"
	"strconv"
)

type AlimamaController struct {
	beego.Controller
}

//阿里妈妈列表
func (this *AlimamaController) List() {
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplNames = "alimama/list.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//其他排序也必须按sort排序（如果有其他排序则为且的关系）
	if "4" != this.GetString("iSortCol_0") {
		this.Ctx.Input.Request.Form.Set("iSortCol_0", "4")
		this.Ctx.Input.Request.Form.Set("sSortDir_0", "desc")
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

	fileds := []string{"name", "_id", "fid", "level", "sort", "add_time", "update_time", ""}
	table := dateTableCondition(this.Ctx, fileds)

	rows := []interface{}{}
	list, count, err := models.AlimamaList(table)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		opHtmlStr := `<div class="action-buttons" catid="%s" name="%s" level="%s">
			                <a class="green addBtn" title="添加子阿里妈妈" href="javascript:void(0);">
			                    <i class="icon-circle bigger-130"></i>
			                </a>
			                <a class="green updateBtn" title="编辑" href="javascript:void(0);">
			                    <i class="icon-pencil bigger-130"></i>
			                </a>
			                <a class="red delBtn" title="删除" href="javascript:void(0);">
			                    <i class="icon-trash bigger-130"></i>
			                </a>
			            </div>`

		//递归并处理列表
		newlist := make([]map[string]interface{}, 0)
		newlist = this.recursionList(list, newlist, "0", "0")

		for _, row := range newlist {
			opHtml := template.HTML(fmt.Sprintf(opHtmlStr, row["_id"], row["name"], row["level"]))
			line := []interface{}{row["name"], row["_id"], row["fid"], row["level"], row["sort"], row["add_time"], row["update_time"], opHtml}
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

//递归并处理阿里妈妈列表的展示结构
func (this *AlimamaController) recursionList(list, newlist []map[string]interface{}, f, n string) []map[string]interface{} {
	intn, _ := strconv.Atoi(n)
	n = strconv.Itoa(intn + 1)
	prestr := "&nbsp;&nbsp;"
	for _, row := range list {
		level, _ := row["level"].(string)
		fid, _ := row["fid"].(string)
		if n == level && fid == f {
			//阿里妈妈ID的处理
			cid, _ := row["_id"].(bson.ObjectId)
			catId := cid.Hex()
			row["_id"] = catId

			//阿里妈妈名称的处理
			name, _ := row["name"].(string)
			for i := 0; i < intn; i++ {
				name = prestr + name
			}
			row["name"] = name

			newlist = append(newlist, row)
			newlist = this.recursionList(list, newlist, catId, level)
		}
	}
	return newlist
}

//添加阿里妈妈
func (this *AlimamaController) Add() {
	//父阿里妈妈ID
	fid := this.GetString("fid")
	fname := this.GetString("fname")
	flevel := this.GetString("flevel")
	if "" == fid {
		fid = "0"
	}
	if "" == flevel {
		flevel = "0"
	}
	if "" == fname {
		fname = "无"
	}
	level := "1"
	if "0" != flevel {
		flev, _ := strconv.Atoi(flevel)
		level = strconv.Itoa(flev + 1)
	}

	if !this.IsAjax() {
		this.Data["Fid"] = fid
		this.Data["Fname"] = fname
		this.Data["Flevel"] = flevel
		this.Layout = "layout.html"
		this.TplNames = "alimama/add.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//获取参数并校验
	name := this.GetString("name")
	sort := this.GetString("sort")

	hasErr := false
	if "" == fid {
		result["msg"] = "父阿里妈妈ID有误"
		hasErr = true
	}
	if "" == name {
		result["msg"] = "名称有误"
		hasErr = true
	}
	if "" == sort {
		result["msg"] = "排序有误"
		hasErr = true
	}
	if hasErr {
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

	//添加阿里妈妈
	err = models.AddAlimama(fid, level, name, sort, nowTime)
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

//编辑阿里妈妈
func (this *AlimamaController) Update() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//阿里妈妈ID
	catid := this.GetString("catid")
	if "" == catid {
		result["msg"] = "阿里妈妈ID有误"
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

	if !this.IsAjax() {
		//获取阿里妈妈信息
		info, err := models.GetAlimama(catid)
		if nil != err {
			result["msg"] = err.Error()
			this.Data["json"] = result
			this.ServeJson()
			return
		}

		//获取父阿里妈妈信息
		fid, _ := info["fid"].(string)
		//一级阿里妈妈不需要查找父阿里妈妈信息
		if "" != fid && "0" != fid {
			finfo, err := models.GetAlimama(fid)
			if nil != err {
				result["msg"] = err.Error()
				this.Data["json"] = result
				this.ServeJson()
				return
			}
			info["fname"] = finfo["name"]
			info["flevel"] = finfo["level"]
		} else {
			info["fname"] = "无"
			info["flevel"] = "0"
		}

		this.Data["Info"] = info
		this.Layout = "layout.html"
		this.TplNames = "alimama/update.tpl"
		this.Render()
		return
	}

	//获取参数并校验
	name := this.GetString("name")
	sort := this.GetString("sort")

	hasErr := false
	if "" == name {
		result["msg"] = "名称有误"
		hasErr = true
	}
	if "" == sort {
		result["msg"] = "排序有误"
		hasErr = true
	}
	if hasErr {
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//添加阿里妈妈
	err = models.UpdateAlimama(catid, name, sort, nowTime)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	result["succ"] = 1
	result["msg"] = "编辑成功"
	this.Data["json"] = result
	this.ServeJson()
	return
}

//删除阿里妈妈
func (this *AlimamaController) Del() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	catid := this.GetString("catid")
	if "" == catid {
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

	err = models.DelAlimama(catid)
	if nil != err {
		result["msg"] = err.Error()
	} else {
		result["succ"] = 1
		result["msg"] = "删除成功"
	}

	this.Data["json"] = result
	this.ServeJson()
}
