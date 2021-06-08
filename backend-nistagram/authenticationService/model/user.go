package model

type User struct {
	Id int `json:"id"`
	UserId string `json:"userId"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	UserRole UserRole `json:"role"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}