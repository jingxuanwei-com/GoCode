package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "db:g1VYgyrRzfd06M@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	// 自动创建所有不存在的表
	err = DB.AutoMigrate(&User{}, &Article{}, &Nav{}, &ArticleCate{})
	if err != nil {
		fmt.Println("数据表迁移失败:", err)
	} else {
		fmt.Println("数据表迁移成功")
	}
}
