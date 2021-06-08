package model

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	UserRole UserRole `json:"role"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}