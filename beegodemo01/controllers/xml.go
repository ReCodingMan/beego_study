package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)

type XmlController struct {
	beego.Controller
}

func (c *XmlController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Product struct {
	Content string `form:"content" xml:"content"`
	Title string `form:"title" xml:"title"`
}

func (c *XmlController) Xml() {
	p := Product{}
	//str := string(c.Ctx.Input.RequestBody)
	//beego.Info(str)

	var err error
	if e := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); e != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		fmt.Printf("%#v\n", p)
		c.Data["json"] = p
		c.ServeJSON()
	}
}
