package routers

import (
	"beegodemo02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/:id", &controllers.ApiController{}, "get:Api")
}
