package controllers

import (
	"myblogdemo/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

func (c *AddArticleController) Get() {
	c.TplName = "write_article.html"
}

func (c *AddArticleController) RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AddArticleController) AddArticle() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//title: go语言
	//tags: golang
	//short: 1234
	//content: 11111111111
	//id:

	//思路
	//1.从前端获取数据
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")

	//2.将数据存入数据库中
	article := models.Article{
		Id:         0,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     c.GetSession("loginuser").(string),
		Createtime: time.Now().Unix(),
	}
	_, err := models.AddArticle(article)

	//3.返回给前端
	if err == nil {
		resp["code"] = 1
		resp["message"] = "ok"
	} else {
		resp["code"] = 0
		resp["message"] = "error"
	}
}
