package controllers

import (
	"github.com/astaxie/beego"
)

const httpAddr = "httpaddr"
const httpPort = "httpport"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = beego.AppConfig.String(httpAddr)
	c.Data["WebsiteHTTPS"] = beego.AppConfig.String(httpAddr) + ":" + beego.AppConfig.String(httpPort) + "/v1/main"
	c.Data["Email"] = "siarhei.hladkou@gmail.com"
	c.TplName = "index.tpl"

}
