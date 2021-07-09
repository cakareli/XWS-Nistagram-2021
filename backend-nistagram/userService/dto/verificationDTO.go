package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
)

type VerificationDTO struct {

	Id string `json:"_id"`
	UserId string `json:"userId"`
	Name string `json:"verificationName"`
	Surname string `json:"verificationSurname"`
	ImageUrl string `json:"verificationImageUrl"`
	VerificationType model.UserType `json:"verificationType"`
}
