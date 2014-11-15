package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"

	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
	"strconv"
	"time"
)

const (
	ACCOUNT_SECURITY string = "somi_admin_account_token"                                                //账号安全码，用户注册激活等场景使用
	PASSWD_SECURITY  string = "somi_admin_passwd_token"                                                 //密码安全码，密码加密入库
	EMAILREG         string = `^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$` //email正则
	PASSWDREG        string = `^[A-Za-z0-9_]+$`                                                         //设置的密码的正则
	MGO_CONF         string = "mgour"                                                                   //somi mongo连接串的配置名
)

type (
	M map[string]interface{}
)

var (
	//当前时间
	nowTime string = time.Now().Format("2006-01-02 15:04:05")
)

//发送邮件
//mailto string 收件人
//subject string 邮件主题
//body string 邮件内容
//isHtml bool 邮件内容是否是html
func sendEmail(mailto, subject, body string, isHtml bool) (err error) {
	if !isMatch(mailto, EMAILREG) {
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

//判断是否与正则匹配
//param s string  要判断的字符串
//param r string  正则
func isMatch(s, r string) (result bool) {
	reg := regexp.MustCompile(r)
	result = reg.MatchString(s)
	return result
}

//获取表格列表相关数据
func dateTableCondition(ctx *context.Context, fileds []string) (table M) {
	//获取skip和limit
	iDisplayStart := ctx.Input.Query("iDisplayStart")
	skip, _ := strconv.Atoi(iDisplayStart)
	iDisplayLength := ctx.Input.Query("iDisplayLength")
	limit, _ := strconv.Atoi(iDisplayLength)
	table = M{
		"iDisplayStart":  skip,
		"iDisplayLength": limit,
		"sSort":          "-_id",
		"sWhere":         M{},
	}

	//获取搜索条件
	sSearch := ctx.Input.Query("sSearch")
	regSearch := M{}
	if "" != sSearch {
		regSearch = M{"$regex": sSearch, "$options": "i"}
	}

	//获取排序
	iSort := ctx.Input.Query("iSortCol_0")
	sortNo, _ := strconv.Atoi(iSort)
	sortDir := ctx.Input.Query("sSortDir_0")

	//处理搜索条件
	sWhere := []interface{}{}
	for k, v := range fileds {
		if "" == v {
			continue
		}
		isSearch := ctx.Input.Query("bSearchable_" + strconv.Itoa(k))
		if "" != sSearch && "true" == isSearch {
			sWhere = append(sWhere, M{v: regSearch})
		}

		//排序
		if k == sortNo {
			sSort := "+" + v
			if "desc" == sortDir {
				sSort = "-" + v
			}
			table["sSort"] = sSort
		}
	}
	where := M{}
	if len(sWhere) > 0 {
		where = M{"$or": sWhere}
	}
	table["sWhere"] = where
	return table
}
