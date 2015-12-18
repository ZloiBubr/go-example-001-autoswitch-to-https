package routers

import (
	"enterprise2020/controllers"
	"github.com/astaxie/beego"
	"enterprise2020/redirect"
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

