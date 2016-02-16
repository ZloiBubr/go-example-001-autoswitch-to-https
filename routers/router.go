package routers

import (
	"github.com/zloibubr/go-example-001-autoswitch-to-https/controllers"
	"github.com/astaxie/beego"
	"github.com/zloibubr/go-example-001-autoswitch-to-https/redirect"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "127.0.0.1" {
				return true
			}
			return false
		}),
		beego.NSBefore(redirect.Https),
		beego.NSNamespace("/main",
			beego.NSRouter("/",
				&controllers.SecondaryController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

