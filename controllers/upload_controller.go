package controllers

import (
	"fmt"
	"io"
	"myblogdemo/models"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UploadController) UploadAlbum() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	fmt.Println("fileupload...")
	//upload:(binary)
	fileData, fileHeader, err := c.GetFile("upload")
	if err != nil {
		c.responseErr(err)
		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
	fmt.Println("fileData:", fileData)

	now := time.Now()
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileType := "other"

	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		c.responseErr(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathSrt := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathSrt)
	if err != nil {
		c.responseErr(err)
		return
	}
	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		c.responseErr(err)
		return
	}
	if fileType == "img" {
		album := models.Album{
			Id:         0,
			Filepath:   filePathSrt,
			Filename:   fileName,
			Status:     0,
			Createtime: timeStamp,
		}
		models.InsertAlbum(album)
	}
	resp["code"] = 1
	resp["message"] = "上传成功"
}

func (c *UploadController) responseErr(err error) {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	resp["code"] = 0
	resp["message"] = err

}
