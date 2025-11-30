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
type UserInfo struct {
	Username string `json:"username" form:"username"` // 用户名
	Password string `json:"password" form:"password"` // 密码
	Age      int    `json:"age" form:"age"`           // 年龄
}

// UserInfoXML 处理xml的结构体
type UserInfoXML struct {
	Username string `json:"username" xml:"username"` // 用户名
	Password string `json:"password" xml:"password"` // 密码
	Age      int    `json:"age" xml:"age"`           // 年龄
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

	r.GET("/test", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		//Get请求值传递
		//http://0.0.0.0:8088/test?username=%E9%9A%8F%E4%BE%BF&age=%E6%88%91%E6%98%AF&page=0
		c.JSON(200, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	// 渲染user.html模板
	r.GET("/user", func(c *gin.Context) {
		c.HTML(200, "pages/user", gin.H{
			"title": "用户页面",
		})
	})

	//Get post请求值传递
	r.POST("/addUser", func(c *gin.Context) {
		username := c.PostForm("username")
		passwd := c.PostForm("password")
		age := c.PostForm("age")
		page := c.DefaultPostForm("page", "1")

		//Post请求值传递
		//http://0.0.0.0:8088/test?username=%E9%9A%8F%E4%BE%BF&age=%E6%88%91%E6%98%AF&page=0
		c.JSON(200, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
			"password": passwd,
		})
	})

	//获取post到结构体
	r.POST("/addUser2", func(c *gin.Context) {
		var user UserInfo
		err := c.Bind(&user)
		if err != nil {
			return
		}
		c.JSON(200, user)
	})

	//获取post到xml转换为json结构体
	r.POST("/addUserXML", func(c *gin.Context) {
		var user UserInfoXML
		err := c.BindXML(&user)
		if err != nil {
			return
		}
		c.JSON(200, user)
	})

	//动态路由传值
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(200, "%v", cid)
	})

	// 启动HTTP服务器，监听在8081端口
	err := r.Run(":8088")
	if err != nil {
		return
	}
}
