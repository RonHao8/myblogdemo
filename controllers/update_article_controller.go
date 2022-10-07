package controllers

import (
	"fmt"
	"myblogdemo/models"
)

type UpdateArticleController struct {
	BaseController
}

func (c *UpdateArticleController) RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UpdateArticleController) GetArticle() {
	//思路
	//1.从前端获取id
	id, _ := c.GetInt("id")
	fmt.Println("id:", id)

	//2.根据id从数据库查询文章
	art := models.QueryArticleWithId(id)

	//3.将数据返回给前端
	c.Data["Title"] = art.Title
	c.Data["Tags"] = art.Tags
	c.Data["Short"] = art.Short
	c.Data["Content"] = art.Content
	c.Data["Id"] = art.Id

	c.TplName = "write_article.html"
}

func (c *UpdateArticleController) UpdateArticle() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//思路
	//title: 22222
	//tags: 1
	//short: 111
	//content: 222223333111111
	//id: 2
	//1.从前端获取数据
	id, _ := c.GetInt("id")
	fmt.Println("id:", id)
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")

	//2.修改数据库
	article := models.Article{
		Id:      id,
		Title:   title,
		Tags:    tags,
		Short:   short,
		Content: content,
	}
	_, err := models.UpdateArticle(article)

	//3.返回数据给前端
	if err == nil {
		resp["code"] = 1
		resp["message"] = "更新成功"
	} else {
		resp["code"] = 0
		resp["message"] = "更新失败"
	}
}
