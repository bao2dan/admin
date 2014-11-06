package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	_ "admin/routers"

	"strings"
)

func main() {
	//判断是否已登陆
	var FilterIsLogined = func(ctx *context.Context) {
		//controller/method
		url := ctx.Input.Url()
		account, ok := ctx.Input.Session("account").(string)
		if !strings.HasPrefix(url, "/site/") {
			if !ok || "" == account {
				ctx.Redirect(302, "/site/login")
			}
		} else if strings.HasPrefix(url, "/site/login") {
			if ok && "" != account {
				ctx.Redirect(302, "/")
			}
		}

	}
	//添加判断是否已登陆方法验证（在执行Controller前）
	beego.InsertFilter("/*", beego.BeforeExec, FilterIsLogined)
	//还需要多加一个，第一个不能截住"http://host:port"这样没有controller/method的请求（原因未知）
	beego.InsertFilter("/", beego.BeforeExec, FilterIsLogined)
	beego.Run()
}
