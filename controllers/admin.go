package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"

	"errors"
	"regexp"
	"strconv"
	"time"
)

type AdminController struct {
	beego.Controller
}

//登陆
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

	//获取参数并校验
	p, err := this.getParams()
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//连接mongodb
	db, err := models.ConnectMgo(p["mgourl"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer db.Session.Close()

	//获取管理员信息
	info, err := models.LoginGetAdminInfo(db, p["collection"], p["uname"], p["passwd"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	//设置session并返回
	if role, ok := info["role"]; ok && "" != role {
		this.SetSession("uname", p["uname"])
		this.SetSession("role", role)
		models.SetAdminLoginTime(db, p["collection"], p["uname"], nowTime)
		result["succ"] = 1
		result["msg"] = "登陆成功"
	} else {
		result["msg"] = "登陆失败或用户不存在"
	}
	this.Data["json"] = result
	this.ServeJson()
}

//退出
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

	//连接mongodb
	db, err := models.ConnectMgo(p["mgourl"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer db.Session.Close()

	//判断是否已经存在该账号
	info, err := models.GetAdminInfo(db, p["collection"], p["uname"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	if getuname, ok := info["uname"]; ok && "" != getuname {
		result["msg"] = "该账号已经存在"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//取账号+安全码+当天0点时间戳 的md5作为token(激活时要验证它)
	year, month, day := time.Now().Date()
	dayTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	dayTimeStr := strconv.FormatInt(dayTime, 10)
	token := md5Encode(p["uname"] + UNAME_SECURITY + dayTimeStr)

	//当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	//保存注册信息
	err = models.InsertAdminInfo(db, p["collection"], p["uname"], p["passwd"], token, nowTime)
	if nil == err {
		//发送激活邮件
		this.sendActivateMail(p["uname"], p["uname"], token)
		result["succ"] = 1
		result["msg"] = "注册成功"
	} else {
		result["msg"] = "注册失败或该账号已存在"
	}
	this.Data["json"] = result
	this.ServeJson()
}

//激活
func (this *AdminController) Activate() {
	//获取相关参数
	gettoken := this.GetString("token")
	uname := this.GetString("uname")
	if "" == gettoken || "" == uname {
		this.Ctx.WriteString("参数有误")
		return
	}

	//取账号+安全码+当天0点时间戳 的md5作为token(激活链接上带有它)
	year, month, day := time.Now().Date()
	dayTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	dayTimeStr := strconv.FormatInt(dayTime, 10)
	token := md5Encode(uname + UNAME_SECURITY + dayTimeStr)
	if gettoken != token {
		this.Ctx.WriteString("参数有误")
		return
	}

	//获取mongo配置
	collection := beego.AppConfig.String("table_admin_user")
	mgourl_ir, _ := beego.GetConfig("string", "mgourlsomi")
	mgourl, _ := mgourl_ir.(string)
	if "" == mgourl {
		this.Ctx.WriteString("Config mgourl is not exists")
		return
	}

	//连接mongodb
	db, err := models.ConnectMgo(mgourl)
	if nil != err {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
	defer db.Session.Close()

	//判断是否存在未激活的账号
	info, err := models.GetNotActivateAdmin(db, collection, uname)
	if nil != err {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
	if getuname, ok := info["uname"]; !ok || "" == getuname {
		this.Ctx.WriteString("激活失败：该账号已经激活或不存在")
		return
	}

	//当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	//激活管理员
	err = models.UnlockAdmin(db, collection, uname, nowTime)
	if nil == err {
		this.Redirect("/admin/login", 302)
		return
	} else {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
}

//发送账号激活邮件
func (this *AdminController) sendActivateMail(uname, mailto, token string) (err error) {
	activateUrl := "http://" + this.Ctx.Request.Host + "/admin/activate?token=" + token + "&uname=" + uname
	subject := "SOMI管理员账号激活邮件"
	body := "<style type='text/css'>div.emailrow{color:#FF5511;line-height:40px;padding-left:36px;}</style>"
	body += "<h3 style='color:#FF5511;'>尊敬的用户：</h3>"
	body += "<div class='emailrow'>您好！</div>"
	body += "<div class='emailrow'>恭喜您即将成为 SOMI 网的管理员；</div>"
	body += "<div class='emailrow'>请牢记账号，并严格履行管理员的职责；</div>"
	body += "<div class='emailrow'>点击此链接 <a href='" + activateUrl + "' target='_blank'>激活</a> 您的账号 " + uname + "（当天有效）；</div>"
	body += "<div class='emailrow'>（系统发送，请务回复）</div>"
	err = sendEmail(mailto, subject, body, true)
	return err
}

//获取管理员登陆或注册公共方法
func (this *AdminController) getParams() (p map[string]string, err error) {
	//获取并校验参数
	uname := this.GetString("uname")
	passwd := this.GetString("passwd")
	regEmail := regexp.MustCompile(EMAILREG)
	isEmail := regEmail.MatchString(uname)
	regPasswd := regexp.MustCompile(PASSWDREG)
	isPasswd := regPasswd.MatchString(passwd)
	if "" == uname || !isEmail || "" == passwd {
		err = errors.New("账号或密码错误")
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

	//md5 加密
	passwd = md5Encode(passwd + PASSWD_SECURITY)

	//获取mongo配置
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
