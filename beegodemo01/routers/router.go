package routers

import (
	"beegodemo01/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.TestController{})
    beego.Router("/testpost", &controllers.TestController{}, "get:TestPost")
    beego.Router("/tget", &controllers.TestController{}, "get:TGet")
    beego.Router("/tpost", &controllers.TestController{}, "post:TPost")
    beego.Router("/tform", &controllers.TestController{}, "post:TForm")
    beego.Router("/tgetdata", &controllers.TestController{}, "get:TGetData")
    beego.Router("/tput", &controllers.TestController{}, "put:TPut")
    beego.Router("/tdelete", &controllers.TestController{}, "delete:TDelete")
    beego.Router("/xml", &controllers.XmlController{}, "post:Xml")
}
