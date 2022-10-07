package controllers

type ExitController struct {
	BaseController
}

// Get 退出
func (c *ExitController) Get() {
	//清除该用户登录状态的数据
	c.DelSession("loginuser")
	c.Redirect("/", 302)
}
