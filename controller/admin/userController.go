package admin

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	con.Success(c)
}

func (con UserController) Add(c *gin.Context) {
	c.HTML(200, "admin/useradd", gin.H{})
}

func (con UserController) Show(c *gin.Context) {
	c.String(200, "管理员用户详情")
}

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")

	file, err := c.FormFile("face")
	if err != nil {
		c.JSON(400, gin.H{"error": "获取文件失败: " + err.Error()})
		return
	}

	uploadDir := "./upload"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": "创建目录失败: " + err.Error()})
		return
	}

	dst := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{"error": "保存文件失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}
