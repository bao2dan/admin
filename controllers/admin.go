package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
)

type AdminController struct {
	beego.Controller
}

var (
	EMAILREG  string = `^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$`
	PASSWDREG string = `^[A-Za-z0-9_]+$`
)

//login system
func (this *AdminController) Login() {
	if !this.IsAjax() {
		uname := this.GetSession("uname")
		if nil != uname && "" != uname {
			this.Redirect("/", 302)
		}
		this.TplNames = "admin/login.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//get params
	p, err := this.getParams()
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//connect mongo db
	db, err := models.ConnectMgo(p["mgourl"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer db.Session.Close()

	//get admin info
	info, err := models.GetAdminInfo(db, p["collection"], p["uname"], p["passwd"], "0")
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//set session and return
	if getuname, ok := info["uname"]; ok && "" != getuname {
		this.SetSession("uname", p["uname"])
		result["succ"] = 1
		result["msg"] = "登陆成功"
	} else {
		result["msg"] = "登陆失败或用户不存在"
	}
	this.Data["json"] = result
	this.ServeJson()
}

//登陆
func (this *AdminController) Logout() {
	this.DelSession("uname")
	this.Redirect("/admin/login", 302)
}

//注册
func (this *AdminController) Register() {
	if !this.IsAjax() {
		this.TplNames = "admin/register.tpl"
		this.Render()
		return
	}

	//result map
	result := map[string]interface{}{"succ": 0, "msg": ""}

	//获取相关参数
	p, err := this.getParams()
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//connect mongo db
	db, err := models.ConnectMgo(p["mgourl"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer db.Session.Close()

	//判断是否已经存在该用户名
	info, err := models.GetAdminInfo(db, p["collection"], p["uname"], p["passwd"], "0")
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	if getuname, ok := info["uname"]; ok && "" != getuname {
		result["msg"] = errors.New("该用户名已经存在")
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//取用户名+安全码的md5作为token(激活时要验证它)
	h := md5.New()
	h.Write([]byte(p["uname"] + UNAME_SECURITY))
	token := hex.EncodeToString(h.Sum(nil))

	//保存注册信息
	err = models.InsertAdminInfo(db, p["collection"], p["uname"], p["passwd"], token)
	if nil == err {
		result["succ"] = 1
		result["msg"] = "注册成功"
	} else {
		result["msg"] = "注册失败或该用户名已存在"
	}
	this.Data["json"] = result
	this.ServeJson()
}

//获取管理员登陆或注册公共方法
func (this *AdminController) getParams() (p map[string]string, err error) {
	//get params
	uname := this.GetString("uname")
	passwd := this.GetString("passwd")
	regEmail := regexp.MustCompile(EMAILREG)
	isEmail := regEmail.MatchString(uname)
	regPasswd := regexp.MustCompile(PASSWDREG)
	isPasswd := regPasswd.MatchString(passwd)
	if "" == uname || !isEmail || "" == passwd {
		err = errors.New("用户名或密码错误")
		return p, err
	}
	if !isPasswd {
		err = errors.New("密码必须为字母、数字、下划线")
		return p, err
	}
	if len(passwd) < 8 {
		err = errors.New("密码至少需要8个字符")
		return p, err
	}

	//md5 encode
	h := md5.New()
	h.Write([]byte(passwd + PASSWD_SECURITY))
	passwd = hex.EncodeToString(h.Sum(nil))

	//get config for collection
	collection := beego.AppConfig.String("table_admin_user")
	mgourl_ir, _ := beego.GetConfig("string", "mgourlsomi")
	mgourl, _ := mgourl_ir.(string)
	if "" == mgourl {
		err = errors.New("Config mgourl is not exists")
		return p, err
	}
	p = map[string]string{
		"uname":      uname,
		"passwd":     passwd,
		"mgourl":     mgourl,
		"collection": collection,
	}
	return p, err
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
