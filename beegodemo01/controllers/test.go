package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Data["title"] = "标题"
	c.Data["num"] = 12
	c.TplName = "test.tpl"
}

func (c *TestController) TestPost() {
	c.Ctx.WriteString("新闻列表")//直接页面返回数据
}
