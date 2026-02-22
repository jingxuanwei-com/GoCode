package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判断是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL.Path)
	c.Set("user", "张三")

	//定义goroutine统计
	cCp := c.Copy()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("统计信息" + cCp.Request.URL.Path)
	}()
}
