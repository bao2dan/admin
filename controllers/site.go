package controllers

import (
	"github.com/astaxie/beego"

	"admin/models"

	"errors"
	"strconv"
	"time"
)

type SiteController struct {
	beego.Controller
}

//登陆
func (this *SiteController) Login() {
	if !this.IsAjax() {
		this.TplNames = "site/login.tpl"
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
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	//获取账号信息
	info, err := models.LoginGetAdminInfo(p["account"], p["passwd"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//当前时间
	//nowTime := time.Now().Format("2006-01-02 15:04:05")

	//设置session并返回
	if role, ok := info["role"]; ok && "" != role {
		this.SetSession("account", p["account"])
		this.SetSession("role", role)
		//models.SetAdminLoginTime(mgoCon, p["account"], nowTime)
		result["succ"] = 1
		result["msg"] = "登陆成功"
	} else {
		result["msg"] = "登陆失败或账号不存在"
	}
	this.Data["json"] = result
	this.ServeJson()
}

//退出
func (this *SiteController) Logout() {
	this.DelSession("account")
	this.Redirect("/site/login", 302)
}

//注册
func (this *SiteController) Register() {
	if !this.IsAjax() {
		this.TplNames = "site/register.tpl"
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
	mgoCon, err := models.ConnectMgo(MGO_CONF)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer mgoCon.Close()

	//判断是否已经存在该账号
	info, err := models.GetAdminInfo(mgoCon, p["account"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	if getaccount, ok := info["account"]; ok && "" != getaccount {
		result["msg"] = "该账号已经存在"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//取账号+安全码+当天0点时间戳 的md5作为token(激活时要验证它)
	year, month, day := time.Now().Date()
	dayTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	dayTimeStr := strconv.FormatInt(dayTime, 10)
	token := md5Encode(p["account"] + ACCOUNT_SECURITY + dayTimeStr)

	//当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	//保存注册信息
	err = models.InsertAdminInfo(mgoCon, p["account"], p["passwd"], token, nowTime)
	if nil == err {
		//发送激活邮件
		this.sendActivateMail(p["account"], p["account"], token)
		result["succ"] = 1
		result["msg"] = "注册成功"
	} else {
		result["msg"] = "注册失败或该账号已存在"
	}
	this.Data["json"] = result
	this.ServeJson()
}

//激活
func (this *SiteController) Activate() {
	//获取相关参数
	gettoken := this.GetString("token")
	account := this.GetString("account")
	if "" == gettoken || "" == account {
		this.Ctx.WriteString("参数有误")
		return
	}

	//取账号+安全码+当天0点时间戳 的md5作为token(激活链接上带有它)
	year, month, day := time.Now().Date()
	dayTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	dayTimeStr := strconv.FormatInt(dayTime, 10)
	token := md5Encode(account + ACCOUNT_SECURITY + dayTimeStr)
	if gettoken != token {
		this.Ctx.WriteString("参数有误")
		return
	}

	//连接mongodb
	mgoCon, err := models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
	defer mgoCon.Close()

	//判断是否存在未激活的账号
	info, err := models.GetNotActivateAdmin(mgoCon, account)
	if nil != err {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
	if getaccount, ok := info["account"]; !ok || "" == getaccount {
		this.Ctx.WriteString("激活失败：该账号已经激活或不存在")
		return
	}

	//当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	//激活账号
	err = models.UnlockAdmin(mgoCon, account, nowTime)
	if nil == err {
		this.Redirect("/site/login", 302)
		return
	} else {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
}

//发送账号激活邮件
func (this *SiteController) sendActivateMail(account, mailto, token string) (err error) {
	activateUrl := "http://" + this.Ctx.Request.Host + "/site/activate?token=" + token + "&account=" + account
	subject := "SOMI管理员账号激活邮件"
	body := "<style type='text/css'>div.emailrow{color:#FF5511;line-height:40px;padding-left:36px;}</style>"
	body += "<h3 style='color:#FF5511;'>尊敬的用户：</h3>"
	body += "<div class='emailrow'>您好！</div>"
	body += "<div class='emailrow'>恭喜您即将成为 SOMI 网的管理员；</div>"
	body += "<div class='emailrow'>请牢记账号，并严格履行管理员的职责；</div>"
	body += "<div class='emailrow'>点击此链接 <a href='" + activateUrl + "' target='_blank'>激活</a> 您的账号 " + account + "（当天有效）；</div>"
	body += "<div class='emailrow'>（系统发送，请务回复）</div>"
	err = sendEmail(mailto, subject, body, true)
	return err
}

//获取账号登陆或注册公共方法
func (this *SiteController) getParams() (p map[string]string, err error) {
	//获取并校验参数
	account := this.GetString("account")
	passwd := this.GetString("passwd")
	isEmail := isMatch(account, EMAILREG)
	isPasswd := isMatch(passwd, PASSWDREG)
	if "" == account || !isEmail || "" == passwd {
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

	p = map[string]string{
		"account": account,
		"passwd":  passwd,
	}
	return p, err
}
