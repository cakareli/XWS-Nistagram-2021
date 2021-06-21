package dto

import "XWS-Nistagram-2021/backend-nistagram/userService/model"

type VerificationRequestDTO struct {
	UserId string `json:"_id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	ImageUrl string `json:"imageUrl"`
	VerificationType model.UserType `json:"verificationType"`
}