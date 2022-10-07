package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser interface{}
}

// Prepare 判断是否登录
func (c *BaseController) Prepare() {
	loginuser := c.GetSession("loginuser")
	fmt.Println("loginuser--->", loginuser)
	if loginuser != nil {
		c.IsLogin = true
		c.LoginUser = loginuser
	} else {
		c.IsLogin = false
	}
	c.Data["IsLogin"] = c.IsLogin
}
