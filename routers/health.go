package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/casdoor/casdoor/controllers"
)

func initHealthRouter() {
	ns := web.NewNamespace("/api",
		web.NSNamespace("/health",
			web.NSRouter("/", &controllers.HealthController{}, "get:GetHealthStatus"),
		),
		web.NSNamespace("/ping",
			web.NSRouter("/", &controllers.HealthController{}, "get:Ping"),
		),
	)
	web.AddNamespace(ns)
}

func init() {
	initHealthRouter()
}
