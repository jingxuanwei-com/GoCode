package router

import (
	"awesomeProject/controller/nannanwu"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {
	defaultRouter := r.Group("/")

	{
		defaultRouter.GET("", nannanwu.DefaultController{}.Index)
		defaultRouter.GET("/article", nannanwu.DefaultController{}.Article)
		defaultRouter.GET("/deletecookie", nannanwu.DefaultController{}.DeleteCookie)
		defaultRouter.GET("/success", nannanwu.DefaultController{}.Success)
		defaultRouter.GET("/error", nannanwu.DefaultController{}.Error)
	}
}
