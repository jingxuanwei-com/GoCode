package admin

import "github.com/gin-gonic/gin"

type ArticleController struct {
}

func (con ArticleController) Article(c *gin.Context) {
	c.String(200, "管理员文章列表")
}
