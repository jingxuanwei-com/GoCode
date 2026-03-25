package admin

import (
	"awesomeProject/model"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

func (con ArticleController) Index(c *gin.Context) {
	articleList := []model.Article{}
	model.DB.Preload("ArticleCate").Find(&articleList)
	c.JSON(200, articleList)

}

func (con ArticleController) Article(c *gin.Context) {
	c.String(200, "管理员文章列表")
}
