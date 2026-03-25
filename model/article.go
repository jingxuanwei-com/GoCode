package model

type Article struct {
	Id          int
	Title       string
	CateId      int
	State       int
	ArticleCate ArticleCate `gorm:"foreignKey:CateId;references:State"`
}

func (Article) TableName() string {
	return "article"
}
