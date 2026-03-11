package nannanwu

import (
	"awesomeProject/model"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
	//继承BaseController
	BaseController
}

func (con DefaultController) Index(c *gin.Context) {
	//设置session
	session := sessions.Default(c)
	session.Set("username", "张三sessions")
	err := session.Save()
	if err != nil {
		return
	}

	fmt.Println(model.UnixToTime(1629788418))
	//设置cookie
	c.SetCookie("username", "张三", 3600, "/", "", false, false)
	c.HTML(200, "pages/index", gin.H{
		"title": "首页",
		"t":     1629788418,
	})
}
func (con DefaultController) Article(c *gin.Context) {
	//获取session
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "session=%v", username)
	////获取cookie
	//username, err := c.Cookie("username")
	//if err != nil {
	//	c.String(200, "获取cookie失败")
	//	return
	//}
	//c.String(200, "cookie="+username)

}

func (con DefaultController) DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "张三", -1, "/", "", false, false)
	_, err := c.Cookie("username")
	if err != nil {
		c.String(200, "删除cookie成功")
		return
	}
}

// Success 使用继承的方式调用BaseController的success方法
func (con DefaultController) Success(c *gin.Context) {
	BaseController{}.success(c)
}

// 使用继承的方式调用BaseController的error方法
func (con DefaultController) Error(c *gin.Context) {
	BaseController{}.error(c)
}
