package routers

import (
	"github.com/astaxie/beego"
	"myblogdemo/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:HomeShow")
	beego.Router("/register", &controllers.RegisterController{}, "get:RegisterShow;post:Register")
	beego.Router("/login", &controllers.LoginController{}, "get:LoginShow;post:Login")
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/article/add", &controllers.AddArticleController{}, "get:Get;post:AddArticle")
	beego.Router("/article/:id", &controllers.ShowArticleController{}, "get:ArticleShow")
	beego.Router("/article/update", &controllers.UpdateArticleController{}, "get:GetArticle;post:UpdateArticle")
	beego.Router("/article/delete", &controllers.DeleteArticleController{}, "get:DeleteArticle")
	beego.Router("/tags", &controllers.TagsController{}, "get:GetTags")
	beego.Router("/album", &controllers.AlbumController{}, "get:AlbumShow")
	beego.Router("/upload", &controllers.UploadController{}, "post:UploadAlbum")
	beego.Router("/aboutme", &controllers.AboutMeController{}, "get:AboutMe")
}
