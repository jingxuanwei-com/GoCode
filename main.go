package main

import (
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

// Article 文章结构体，用于存储文章相关信息
type Article struct {
	Title   string // 文章标题
	Content string // 文章内容
}

// UnixToTime 将Unix时间戳转换为格式化的日期时间字符串
// 参数：
//
//	timestamp: Unix时间戳（秒）
//
// 返回值：
//
//	格式化后的日期时间字符串，格式为"2006-01-02 15:04:05"
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// main 函数是程序的入口点，初始化并启动Web服务器
func main() {
	// 创建默认的gin引擎
	r := gin.Default()

	// 注册模板函数，使模板可以使用自定义函数UnixToTime
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	// 加载模板文件，支持多层目录结构
	r.LoadHTMLGlob("templates/**/*")

	// 配置静态文件目录，将./static目录映射到URL路径/static
	r.Static("/static", "./static")

	// 设置路由
	// 根路由处理函数，渲染首页
	r.GET("/", func(c *gin.Context) {
		// 渲染pages/index.html模板，并传递数据
		c.HTML(200, "pages/index", gin.H{
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
			"data": 1764077331, // Unix时间戳示例
		})
	})

	// 启动HTTP服务器，监听在8081端口
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
