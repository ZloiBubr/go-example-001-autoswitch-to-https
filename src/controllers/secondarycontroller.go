package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SecondaryController struct {
	beego.Controller
}

func (c *SecondaryController) Get() {
	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
	log.Info("SecondaryController called!")

	c.Data["Website"] = "enterprise2020.com"
	c.Data["Email"] = "siarhei.hladkou@gmail.com"
	c.TplNames = "secure.tpl"

}
