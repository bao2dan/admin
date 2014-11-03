package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

type AdminController struct {
	beego.Controller
}

var (
	EMAILREG string = `^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$`
)

//login system
func (this *AdminController) Login() {
	if !this.IsAjax() {
		this.TplNames = "admin/login.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//get params
	uname := this.GetString("uname")
	passwd := this.GetString("passwd")
	reg := regexp.MustCompile(EMAILREG)
	isUname := reg.MatchString(uname)
	if "" == uname || !isUname || "" == passwd {
		result["msg"] = "username or password is error"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//md5 encode
	h := md5.New()
	h.Write([]byte(passwd + PASSWD_TOKEN))
	passwd = hex.EncodeToString(h.Sum(nil))

	//get config for collection
	collname := beego.AppConfig.String("table_admin_user")
	mgourl_ir, _ := beego.GetConfig("string", "mgourlsomi")
	mgourl, _ := mgourl_ir.(string)
	if "" == mgourl {
		result["msg"] = "Config mgourl is not exists"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//connect mongo db
	db, err := models.ConnectMgo(mgourl)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer db.Session.Close()

	//get admin info
	info, err := models.AGetAdminInfo(db, collname, uname, passwd)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//set session and return
	if getuname, ok := info["uname"]; ok && "" != getuname {
		this.SetSession("uname", uname)
		result["succ"] = 1
		result["msg"] = "login success"
	}
	this.Data["json"] = result
	this.ServeJson()
}

//logout system
func (this *AdminController) Logout() {
	this.DelSession("uname")
	this.Redirect("/admin/login", 302)
}

func (this *AdminController) Register() {
	this.Ctx.WriteString("xxxxxxxxxxxccccccccccccccccccc")
}

func (this *AdminController) Getpwd() {
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
