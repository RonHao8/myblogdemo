package controllers

import (
	"fmt"
	"myblogdemo/models"
)

type HomeController struct {
	//beego.Controller
	BaseController
}

// HomeShow 主页
func (c *HomeController) HomeShow() {
	tag := c.GetString("tag")
	fmt.Println("tag=", tag)
	page, _ := c.GetInt("page")
	var artList []models.Article
	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		c.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		//var artList []models.Article
		artList, _ = models.FindArticleWithPage(page)
		//fmt.Println("artlist--->", artList)
		c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		c.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin-->", c.IsLogin, c.LoginUser)
	c.Data["Content"] = models.MakeHomeBlocks(artList, c.IsLogin)
	c.TplName = "home.html"
}
