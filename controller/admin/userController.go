package admin

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "管理员用户列表")
}

func (con UserController) Show(c *gin.Context) {
	c.String(200, "管理员用户详情")
}
