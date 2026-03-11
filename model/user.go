package model

type User struct {
	Id       int
	Name     string
	Age      int
	Email    string
	Password string
	AddTime  int
}

func (User) TableName() string {
	return "user"
}
