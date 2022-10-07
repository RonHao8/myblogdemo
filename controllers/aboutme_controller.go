package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) AboutMe() {
	c.Data["wechat"] = "微信：123456789"
	c.Data["qq"] = "qq:123456789"
	c.Data["tel"] = "tel:123456789"
	c.TplName = "aboutme.html"
}
