package api

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	c.String(200, "API接口")
}

func (con IndexController) User(c *gin.Context) {
	c.String(200, "用户信息")
}

func (con IndexController) Plist(c *gin.Context) {
	c.String(200, "用户列表")
}
