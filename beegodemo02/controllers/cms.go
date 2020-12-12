package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}

func (c *CmsController) Get() {
	c.Ctx.WriteString("默认接口")
}

func (c *CmsController) Cms() {
	//获取动态路由的值
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("cms详情----" + id)
}
