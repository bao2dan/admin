package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"

	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
)

const (
	ACCOUNT_SECURITY string = "somi_admin_account_token" //账号安全码，用户注册激活等场景使用
	PASSWD_SECURITY  string = "somi_admin_passwd_token"  //密码安全码，密码加密入库
	EMAILREG         string = `^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$`
	PASSWDREG        string = `^[A-Za-z0-9_]+$`
)

//发送邮件
//mailto string 收件人
//subject string 邮件主题
//body string 邮件内容
//isHtml bool 邮件内容是否是html
func sendEmail(mailto, subject, body string, isHtml bool) (err error) {
	regEmail := regexp.MustCompile(EMAILREG)
	isEmail := regEmail.MatchString(mailto)
	if !isEmail {
		err = errors.New("mailto not is email")
		return err
	}

	myMail := beego.AppConfig.String("mail_name")
	myMailpasswd := beego.AppConfig.String("mail_passwd")
	myMailHost := beego.AppConfig.String("mail_host")
	myMailPort := beego.AppConfig.String("mail_port")

	config := `{"username":"` + myMail + `","password":"` + myMailpasswd + `","host":"` + myMailHost + `","port":` + myMailPort + `}`
	mail := utils.NewEMail(config)
	if "" == mail.Username || "" == mail.Password || "" == mail.Host || 0 == mail.Port {
		err = errors.New("email parse get params error")
		return err
	}

	mail.From = myMail
	mail.To = []string{mailto}
	mail.Subject = subject
	if isHtml {
		mail.HTML = body
	} else {
		mail.Text = body
	}

	mail.Send()
	return err
}

//md5加密
//param s string 要加密的字符串
func md5Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	s = hex.EncodeToString(h.Sum(nil))
	return s
}

//判断是否是超级管理员
//param account string 账号
func IsRoot(account string) (isRoot bool) {
	var controller *beego.Controller
	role := controller.GetSession("role")
	isRoot = false
	if "root" == role {
		isRoot = true
	}
	return isRoot
}
