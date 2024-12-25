package model

type User struct {
	Id       uint   `json:"id" gorm:"id"`
	Username string `json:"username" gorm:"username"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
