package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblogdemo/models"
	"myblogdemo/utils"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

// RegisterShow 注册页面
func (c *RegisterController) RegisterShow() {
	c.TplName = "register.html"
}

// Register 注册业务逻辑代码
func (c *RegisterController) Register() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//username: usertest
	//password: 123456
	//repassword: 123456

	//思路
	//1.从前端拿到数据
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	fmt.Println(username, password, repassword)

	//2.判断数据的合法性
	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		resp["code"] = 0
		resp["message"] = "用户名已经存在"
		return
	}
	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后:", password)

	//3.将数据存入数据库中
	user := models.User{
		Id:         0,
		Username:   username,
		Password:   password,
		Status:     0,
		Createtime: time.Now().Unix(),
	}

	_, err := models.InsertUser(user)
	//4.返回给前端
	if err != nil {
		resp["code"] = 0
		resp["message"] = "注册失败"
	} else {
		resp["code"] = 1
		resp["message"] = "注册成功"
		//c.Redirect("/", 302)
	}

}
