package router

import (
	"awesomeProject/controller/api"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("", api.IndexController{}.Index)
		apiRouter.GET("/user", api.IndexController{}.User)
		apiRouter.GET("/plist", api.IndexController{}.Plist)
	}
}
