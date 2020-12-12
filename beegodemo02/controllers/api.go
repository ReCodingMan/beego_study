package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	c.Ctx.WriteString("默认接口")
}

func (c *ApiController) Api() {
	//获取动态路由的值
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api接口----" + id)
}
