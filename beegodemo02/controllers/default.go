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

	//4、模版中循环map
	userInfo := make(map[string]interface{})
	userInfo["username"] = "张三"
	userInfo["age"] = 20
	userInfo["sex"] = "男"
	c.Data["userInfo"] = userInfo

	//5、循环遍历结构体类型
	c.Data["articleList"] = []Article{
		{
			Title: "这是title",
			Content: "这是内容",
		},{
			Title: "这是title",
			Content: "这是内容",
		},{
			Title: "这是title",
			Content: "这是内容",
		},{
			Title: "这是title",
			Content: "这是内容",
		},
	}

	c.TplName = "index.tpl"
}
