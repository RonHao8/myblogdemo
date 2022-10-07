package controllers

import (
	"fmt"
	"log"
	"myblogdemo/models"
)

type DeleteArticleController struct {
	BaseController
}

func (c *DeleteArticleController) DeleteArticle() {
	artID, _ := c.GetInt("id")
	fmt.Println("删除id:", artID)
	_, err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	c.Redirect("/", 302)
}
