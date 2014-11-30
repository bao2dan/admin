package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"

	"admin/models"

	"errors"
	"fmt"
	"html/template"
	"strconv"
	"time"
)

type SiteController struct {
	beego.Controller
}

//默认首页
func (this *SiteController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "index.tpl"
	this.Render()
}

//获取顶部导航栏
func (this *SiteController) Menu() {
	//result map
	result := map[string]interface{}{"succ": 0, "msg": "", "menu": "", "account": ""}

	//获取账号
	account := this.GetSession("account")
	result["account"] = account

	//获取角色并验证
	role, ok := this.GetSession("role").(string)
	if !ok || "" == role {
		result["msg"] = "角色不能为空"
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	//获取导航配置
	aMenu, bMenu, urlInfo, erra := models.GetMenuConfig()
	if nil != erra {
		result["msg"] = "获取导航配置失败"
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//获取权限配置
	var auth []string
	var err error
	if "root" != role {
		auth, err = models.GetAuthConfig(role)
		if nil != err {
			result["msg"] = "获取权限配置失败"
			this.Data["json"] = result
			this.ServeJson()
			return
		}
	}

	//跳转链接的[与downHtml平级的]
	hrefHtml := `<li><a href="%s">%s<span class="menu-text"> %s </span></a></li>`
	//下拉式的
	downHtml := `<li>
                    <a href="#" class="dropdown-toggle">%s<span class="menu-text"> %s </span><b class="arrow icon-angle-down"></b></a>
                    <ul class="submenu">
                        %s
                    </ul>
                </li>`
	//子跳转链接的[是downHtml的子li]
	sonHrefHtml := `<li><a href="%s"><i class="icon-double-angle-right"></i><span class="menu-text"> %s </span></a></li>`

	//进行一、二级导航的权限判断，并返回有权限的一、二级导航字符串
	var getAuthId = func(role, naid, nahref, naname, naimg string, bMenu map[string][]string, auth []string) (strHtml string) {
		downli := ""
		if nb, ok := bMenu[naid]; ok {
			for _, nbid := range nb {
				if utils.InSlice(nbid, auth) || "root" == role {
					if "" != nahref {
						strHtml += fmt.Sprintf(hrefHtml, nahref, naimg, naname)
						break
					} else {
						if nbinfo, ok := urlInfo[nbid]; ok {
							downli += fmt.Sprintf(sonHrefHtml, nbinfo[0], nbinfo[1])
						}
					}
				}
			}
		}
		if "" == nahref && "" != downli {
			strHtml += fmt.Sprintf(downHtml, naimg, naname, downli)
		}
		return strHtml
	}

	//循环拼接
	menuStr := ""
	for _, na := range aMenu {
		naid := na[0]
		nahref := na[1]
		naname := na[2]
		naimg := na[3]
		naStr := getAuthId(role, naid, nahref, naname, naimg, bMenu, auth)
		menuStr += naStr
	}

	menu := template.HTML(menuStr)

	result["succ"] = 1
	result["msg"] = "成功"
	result["menu"] = menu
	this.Data["json"] = result
	this.ServeJson()
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
	info, err := models.GetLoginAdmin(p["account"], p["passwd"])
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}

	//设置session并返回
	if role, ok := info["role"]; ok && "" != role {
		this.SetSession("account", p["account"])
		this.SetSession("role", role)
		models.SetAdminLoginTime(p["account"], nowTime)
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
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		result["msg"] = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	}
	defer models.MgoCon.Close()

	//判断是否已经存在该账号
	info, err := models.GetAdmin(p["account"])
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

	//保存注册信息
	err = models.AddAdminInfo(p["account"], p["passwd"], token, nowTime)
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

	var err error

	//连接mongodb
	models.MgoCon, err = models.ConnectMgo(MGO_CONF)
	if nil != err {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
	defer models.MgoCon.Close()

	//判断是否存在未激活的账号
	info, err := models.GetNotActivateAdmin(account)
	if nil != err {
		this.Ctx.WriteString("激活失败：" + err.Error())
		return
	}
	if getaccount, ok := info["account"]; !ok || "" == getaccount {
		this.Ctx.WriteString("激活失败：该账号已经激活或不存在")
		return
	}

	//激活账号
	err = models.UnlockAdmin(account, nowTime)
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

//没有权限
func (this *SiteController) NoAuth() {
	this.TplNames = "site/noauth.tpl"
	this.Render()
}

//获取账号登陆或注册公共方法
func (this *SiteController) getParams() (p map[string]string, err error) {
	//获取并校验参数
	account := this.GetString("account")
	passwd := this.GetString("passwd")
	if "" == account || !isEmail(account) || "" == passwd {
		return p, errors.New("账号或密码错误")
	}
	if !isPasswd(passwd) {
		return p, errors.New("密码必须为字母、数字、下划线")
	}
	if len(passwd) < 8 {
		return p, errors.New("密码至少需要8个字符")
	}

	//md5 加密
	passwd = md5Encode(passwd + PASSWD_SECURITY)

	p = map[string]string{
		"account": account,
		"passwd":  passwd,
	}
	return p, err
}

//判断是否有权限
func IsAuth(role, url string) (has bool, err error) {
	//获取导航配置
	_, _, urlInfo, erra := models.GetMenuConfig()
	if nil != erra {
		return has, errors.New("获取导航配置失败")
	}

	//获取权限配置
	auth, erra := models.GetAuthConfig(role)
	if nil != erra {
		return has, errors.New("获取权限配置失败")
	}

	//循环判断
	for nbid, nbinfo := range urlInfo {
		if url == nbinfo[0] {
			if utils.InSlice(nbid, auth) {
				return true, err
			}
		}
	}

	return has, err
}
