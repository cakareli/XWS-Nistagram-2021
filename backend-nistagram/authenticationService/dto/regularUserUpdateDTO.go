package dto

type RegularUserUpdateDTO struct {
	UserId string `json:"userId"`
	Email string `json:"email"`
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}