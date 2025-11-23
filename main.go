package main

import (
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

func main() {
	r := gin.Default()
	//html模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(200, "首页")
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": "true",
			"msg":     "你好gin",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "true",
			"msg":     "你好gin",
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		a := Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})

	//响应jsonp请求 http://0.0.0.0:8080/jsonp?callback=xxx
	r.GET("/jsonp", func(c *gin.Context) {
		a := Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})

	r.GET("xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"success": "true",
			"msg":     "你好gin，我是一个xml",
		})
	})

	r.GET("news", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"title": "我是一个后台数据",
		})
	})

	r.GET("goods", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"title": "我是一个商品页面",
		})
	})

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
