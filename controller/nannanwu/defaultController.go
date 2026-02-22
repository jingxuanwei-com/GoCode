package nannanwu

import (
	"awesomeProject/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
	//继承BaseController
	BaseController
}

func (con DefaultController) Index(c *gin.Context) {
	fmt.Println(model.UnixToTime(1629788418))
	c.HTML(200, "pages/index", gin.H{
		"title": "首页",
		"t":     1629788418,
	})
}
func (con DefaultController) Article(c *gin.Context) {
	c.String(200, "文章详情")
}

// Success 使用继承的方式调用BaseController的success方法
func (con DefaultController) Success(c *gin.Context) {
	BaseController{}.success(c)
}

// 使用继承的方式调用BaseController的error方法
func (con DefaultController) Error(c *gin.Context) {
	BaseController{}.error(c)
}
