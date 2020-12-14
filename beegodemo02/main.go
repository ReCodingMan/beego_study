package main

import (
	_ "beegodemo02/routers"
	"beegodemo02/models"
	"github.com/astaxie/beego"
)



func main() {
	//beego.BConfig.WebConfig.TemplateLeft = "<<"
	//beego.BConfig.WebConfig.TemplateRight = ">>"

	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.Run()
}