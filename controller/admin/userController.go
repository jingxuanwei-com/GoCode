package admin

import (
	"awesomeProject/model"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//查询数据库
	userList := []model.User{}

	model.DB.Find(&userList)
	c.JSON(200, userList)
	//con.Success(c)
}

func (con UserController) Add(c *gin.Context) {
	user := model.User{
		Name:     "test",
		Age:      18,
		Email:    "test@example.com",
		Password: "password",
	}
	model.DB.Create(&user)
	c.JSON(200, user)
	fmt.Println(user)
	c.String(200, "用户添加成功")
}

func (con UserController) Edit(c *gin.Context) {
	user := model.User{Id: 1}
	model.DB.Find(&user)
	user.Age = 19
	user.Password = "6666"
	user.Name = "哈哈哈"
	model.DB.Save(&user)
	c.JSON(200, user)
}

func (con UserController) Delete(c *gin.Context) {
	user := model.User{Id: 1}
	model.DB.Delete(&user)
	c.String(200, "用户删除成功")
}

func (con UserController) Show(c *gin.Context) {
	c.String(200, "管理员用户详情")
}

func (con UserController) DoUpload(c *gin.Context) {
	// 获取表单中的用户名
	username := c.PostForm("username")

	// 获取上传的文件
	file, err := c.FormFile("face")
	if err != nil {
		c.JSON(400, gin.H{"error": "获取文件失败: " + err.Error()})
		return
	}

	// 创建上传目录
	uploadDir := "./upload"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": "创建目录失败: " + err.Error()})
		return
	}

	// 构建目标文件路径并保存文件
	dst := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{"error": "保存文件失败: " + err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}

func (con UserController) DoEdit(c *gin.Context) {
	username := c.PostForm("username")
	file1, err := c.FormFile("face1")
	if err != nil {
		c.JSON(400, gin.H{"error": "获取文件失败: " + err.Error()})
		return
	}
	uploadDir := "./upload"
	dst := filepath.Join(uploadDir, file1.Filename)
	if err := c.SaveUploadedFile(file1, dst); err != nil {
		c.JSON(500, gin.H{"error": "保存文件失败: " + err.Error()})
		return
	}

	file2, err := c.FormFile("face2")
	if err != nil {
		c.JSON(400, gin.H{"error": "获取文件失败: " + err.Error()})
		return
	}
	dst1 := filepath.Join(uploadDir, file2.Filename)
	if err := c.SaveUploadedFile(file2, dst1); err != nil {
		c.JSON(500, gin.H{"error": "保存文件失败: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
		"dst1":     dst1,
	})
}
