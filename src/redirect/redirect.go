package redirect

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

var (
	https_baseUrl string
	//isProd        bool
)

const httpsConfigKey = "https_base"

func Https(ctx *context.Context) {
	//if isProd {
		if ctx.Input.Request.TLS == nil {
			ctx.Redirect(http.StatusTemporaryRedirect, https_baseUrl+ctx.Input.Request.URL.String())
		}
	//}
}

func init() {
	//isProd = beego.RunMode == "prod"
	//if isProd {
		https_baseUrl = beego.AppConfig.String(httpsConfigKey)
		if len(https_baseUrl) == 0 {
			panic(fmt.Errorf("missing configuration '%s'", httpsConfigKey))
		}
	//}
}