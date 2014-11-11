package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"admin/controllers"
	_ "admin/routers"

	"strings"
)

func main() {
	var filterDeal = func(ctx *context.Context) {
		//loginDeal(ctx)
		//authDeal(ctx)
	}

	//添加过滤处理（在执行Controller前）
	beego.InsertFilter("/*", beego.BeforeExec, filterDeal)
	//还需要多加一个，第一个不能截住"http://host:port"这样没有controller/method的请求（原因未知）
	beego.InsertFilter("/", beego.BeforeExec, filterDeal)
	beego.Run()
}

//判断是否已经登陆
func loginDeal(ctx *context.Context) {
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

//判断是否有权限
func authDeal(ctx *context.Context) {
	//controller/method
	url := ctx.Input.Url()
	if !strings.HasPrefix(url, "/site/") && "/" != url {
		role, ok := ctx.Input.Session("role").(string)
		if !ok || "" == role {
			ctx.Redirect(302, "/site/login")
		} else {
			if "root" == role {
				return
			}
			ok, err := controllers.IsAuth(role, url)
			if nil != err || !ok {
				ctx.Redirect(302, "/site/noauth")
			}
		}
	}
}
