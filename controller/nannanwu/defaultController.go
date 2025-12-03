package nannanwu

import "github.com/gin-gonic/gin"

type DefaultController struct {
	//继承BaseController
	BaseController
}

func (con DefaultController) Index(c *gin.Context) {
	c.String(200, "首页")
}
func (con DefaultController) Article(c *gin.Context) {
	c.String(200, "文章详情")
}

// 使用继承的方式调用BaseController的success方法
func (con DefaultController) Success(c *gin.Context) {
	BaseController{}.success(c)
}

// 使用继承的方式调用BaseController的error方法
func (con DefaultController) Error(c *gin.Context) {
	BaseController{}.error(c)
}
