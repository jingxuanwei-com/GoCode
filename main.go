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
		c.HTML(200, "index", gin.H{})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(200, "about", gin.H{})
	})

	r.GET("/articles", func(c *gin.Context) {
		c.HTML(200, "articles", gin.H{
			"Articles": []map[string]string{
				{"Title": "Gin框架入门", "Date": "2025-11-23"},
				{"Title": "Go语言最佳实践", "Date": "2025-11-20"},
			},
		})
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(200, "contact", gin.H{})
	})

	r.Run(":8080")
}
