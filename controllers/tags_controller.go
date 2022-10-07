package controllers

import (
	"fmt"
	"myblogdemo/models"
)

type TagsController struct {
	BaseController
}

func (c *TagsController) GetTags() {
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))
	c.Data["Tags"] = models.HandleTagsListData(tags)
	c.TplName = "tags.html"
}
