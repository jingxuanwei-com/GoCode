package router

import (
	"awesomeProject/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("", admin.IndexController{}.Index)

		adminRouter.GET("/plist", admin.IndexController{}.Plist)

		adminRouter.GET("/login", admin.IndexController{}.Login)

		adminRouter.GET("/article", admin.ArticleController{}.Article)

		adminRouter.GET("/user", admin.UserController{}.Index)

		adminRouter.GET("/user/:id", admin.UserController{}.Show)
	}
}
