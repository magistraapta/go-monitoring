package user

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
