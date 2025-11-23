package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静态资源映射
	r.Static("/static", "./static")

	// 加载模板（支持多层目录）
	r.LoadHTMLGlob("templates/**/*")

	// 路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "pages/index.html", gin.H{
			"Title": "首页",
		})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(200, "pages/about.html", gin.H{
			"Title": "关于我们",
		})
	})

	r.GET("/articles", func(c *gin.Context) {
		c.HTML(200, "pages/articles.html", gin.H{
			"Title": "文章列表",
			"Articles": []map[string]string{
				{"Title": "Gin框架入门", "Date": "2025-11-23"},
				{"Title": "Go语言最佳实践", "Date": "2025-11-20"},
			},
		})
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(200, "pages/contact.html", gin.H{
			"Title": "联系我们",
		})
	})

	r.Run(":8080")
}
