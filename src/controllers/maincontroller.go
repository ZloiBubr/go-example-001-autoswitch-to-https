package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "127.0.0.1"
	c.Data["WebsiteHTTPS"] = "127.0.0.1:8080/v1/main"
	c.Data["Email"] = "siarhei.hladkou@gmail.com"
	c.TplNames = "index.tpl"

}
