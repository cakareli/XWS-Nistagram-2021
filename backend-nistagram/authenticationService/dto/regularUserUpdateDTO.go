package dto

type RegularUserUpdateDTO struct {
	UserId string `json:"_id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}