package controllers

import (
	"fmt"
	"myblogdemo/models"
	"myblogdemo/utils"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (c *ShowArticleController) ArticleShow() {
	//思路
	//1.从前端获取id
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println(":id", id)

	//2.从数据库中查询id的文章
	art := models.QueryArticleWithId(id)

	//3.将数据返回给前端
	c.Data["Title"] = art.Title
	//resp["Content"] = art.Content
	c.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)

	c.TplName = "show_article.html"
}
