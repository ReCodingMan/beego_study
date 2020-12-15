package routers

import (
	"beegodemo02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/:id", &controllers.ApiController{}, "get:Api")
    beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{}, "get:Cms")
    beego.Router("/cms/func", &controllers.CmsController{}, "get:ModelFunc")
}
