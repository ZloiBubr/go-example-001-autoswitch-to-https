package main

import (
	_ "go-example-001-autoswitch-to-https/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

