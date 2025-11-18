package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//从服务器请求资源
	r.GET("/", func(c *gin.Context) {
		c.String(200, "这是一个get请求，用于取出资源")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "这是一个post--主要用于增加数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个put请求 主要用于编辑数据")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "这是一个delete请求 主要用于删除数据")
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
