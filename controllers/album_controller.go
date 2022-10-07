package controllers

import (
	"github.com/astaxie/beego/logs"
	"myblogdemo/models"
)

type AlbumController struct {
	BaseController
}

func (c *AlbumController) AlbumShow() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Album"] = albums
	c.TplName = "album.html"
}
