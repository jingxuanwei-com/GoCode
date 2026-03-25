package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	AddTime  int    `json:"add_time"`
}

func (User) TableName() string {
	return "user"
}
