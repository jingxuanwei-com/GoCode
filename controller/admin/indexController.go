package admin

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	c.String(200, "管理员用户列表")
}

func (con IndexController) Plist(c *gin.Context) {
	c.String(200, "管理员用户详情")
}

func (con IndexController) Login(c *gin.Context) {
	c.String(200, "管理员登录页面")
}
