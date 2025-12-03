package main

import (
	"awesomeProject/router"

	"github.com/gin-gonic/gin"
)

// main 函数是程序的入口点，初始化并启动Web服务器
func main() {
	// 创建默认的gin引擎
	r := gin.Default()

	// 加载模板文件，支持多层目录结构
	r.LoadHTMLGlob("templates/**/*")

	// 配置静态文件目录，将./static目录映射到URL路径/static
	r.Static("/static", "./static")

	router.AdminRouterInit(r)

	router.ApiRouterInit(r)

	router.DefaultRouterInit(r)

	// 启动HTTP服务器，监听在8081端口
	err := r.Run(":8088")
	if err != nil {
		return
	}
}
