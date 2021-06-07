package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
)

type RegularUserRegistrationDTO struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Role model.UserRole `json:"role"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}