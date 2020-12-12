package controllers

import "C"
import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Article struct {
	Title string
	Content string
}

func (c *MainController) Get() {
	//1、绑定基本类型
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["num"] = 12
	c.Data["flag"] = true

	//2、绑定结构体
	article := Article{
		Title: "我是golang教程",
		Content: "我是内容",
	}
	c.Data["article"] = article

	//3、模版循环切片
	c.Data["sliceList"] = []string{"php", "java", "golang"}

	c.TplName = "index.tpl"
}
