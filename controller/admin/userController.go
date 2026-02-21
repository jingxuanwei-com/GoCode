package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	con.Success(c)
}

func (con UserController) Show(c *gin.Context) {
	c.String(200, "管理员用户详情")
}
