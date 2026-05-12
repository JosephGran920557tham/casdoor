package routers

import (
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor/controllers"
)

func initBuildInfoRouter() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/get-build-info",
			beego.NSRouter("", &controllers.ApiController{}, "get:GetBuildInfo"),
		),
	)
	beego.AddNamespace(ns)
}

func init() {
	initBuildInfoRouter()
}
