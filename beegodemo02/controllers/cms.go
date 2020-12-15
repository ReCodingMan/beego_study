package controllers

import (
	"beegodemo02/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"io"
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

func (c *CmsController) ModelFunc() {
	//md5加密
	h := md5.New()
	io.WriteString(h, "123456") //e10adc3949ba59abbe56e057f20f883e
	fmt.Printf("%x\n", h.Sum(nil))

	date := []byte("123456")
	fmt.Printf("%x\n", md5.Sum(date))

	str := models.UnixToDate(1589898013)
	c.Ctx.WriteString("新闻add" + str)
}
