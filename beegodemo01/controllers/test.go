package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
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

func (c *TestController) TGet() {
	id := c.GetString("id")
	c.Ctx.WriteString("测试 get 方法" + id)//直接页面返回数据
}

func (c *TestController) TPost() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id必须是整数")
		return
	}
	fmt.Printf("%v--%T\n", id, id)
	c.Ctx.WriteString("测试 post 方法" + strconv.Itoa(id))//直接页面返回数据
}

//定义一个结构体
type User struct {
	Username string `form:"username" json:"username2"`
	Password string `form:"password" json:"password2"`
	Hobby []string `form:"hobby" json:"hobby2"`
}

func (c *TestController) TForm() {
	u := User{}
	err := c.ParseForm(&u)
	if err != nil {
		c.Ctx.WriteString("post提交失败")
		return
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("解析post成功")
}

func (c *TestController) TGetData() {
	u := User{
		Username: "张三",
		Password: "123",
		Hobby: []string{"1","2","3"},
	}
	c.Data["json"] = u
	c.ServeJSON()
}
