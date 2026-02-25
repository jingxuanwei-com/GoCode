package router

import (
	"awesomeProject/controller/admin"
	"awesomeProject/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("/admin", middleware.InitMiddleware)
	{
		adminRouter.GET("", admin.IndexController{}.Index)

		adminRouter.GET("/user/add", admin.UserController{}.Add)

		adminRouter.GET("/plist", admin.IndexController{}.Plist)

		adminRouter.GET("/login", admin.IndexController{}.Login)

		adminRouter.GET("/article", admin.ArticleController{}.Article)

		adminRouter.GET("/user", admin.UserController{}.Index)

		adminRouter.POST("/user/doUpload", admin.UserController{}.DoUpload)

		adminRouter.GET("/user/edit", admin.UserController{}.Edit)

		adminRouter.POST("/user/doEdit", admin.UserController{}.DoEdit)

		adminRouter.GET("/user/:id", admin.UserController{}.Show)
	}
}
