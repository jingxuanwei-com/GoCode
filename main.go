package main

import (
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	// 加载模板（支持多层目录）
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	// 路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "pages/index.html", gin.H{
			"title": "Main website",
			"score": 60,
			"hobby": []string{"吃饭", "睡觉", "打豆豆"},
			"newsList": []interface{}{
				Article{
					Title:   "新闻标题",
					Content: "新闻内容1",
				},
				Article{
					Title:   "新闻标题2",
					Content: "新闻内容2",
				},
			},
			"newsList2": []string{},

			"news": Article{
				Title:   "标题",
				Content: "测试2",
			},
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
