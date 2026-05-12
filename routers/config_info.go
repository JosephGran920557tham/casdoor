package routers

import (
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor/controllers"
)

func initConfigInfoRouter() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/config",
			beego.NSRouter("/get-config-info", &controllers.ApiController{}, "get:GetConfigInfo"),
		),
	)
	beego.AddNamespace(ns)
}

func init() {
	initConfigInfoRouter()
}
