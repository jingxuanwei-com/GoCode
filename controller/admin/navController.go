package admin

import (
	"awesomeProject/model"

	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

func (con NavController) Index(c *gin.Context) {
	navList := []model.Nav{}
	model.DB.Find(&navList)
	c.JSON(200, navList)
}
