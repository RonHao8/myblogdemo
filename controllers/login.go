package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblogdemo/models"
	"myblogdemo/utils"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) LoginShow() {
	c.TplName = "login.html"
}

func (c *LoginController) RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *LoginController) Login() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	//username: testuser
	//password: 123456
	//思路
	//1.从前端获取数据
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	//2.根据用户名和密码从数据库判断是否正确
	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	//3.返回给前端
	if id > 0 {
		/*
			设置了session后会将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
			因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		*/
		c.SetSession("loginuser", username) //设置session

		resp["code"] = 1
		resp["message"] = "登陆成功"
	} else {
		resp["code"] = 0
		resp["message"] = "登陆失败"
	}

}
