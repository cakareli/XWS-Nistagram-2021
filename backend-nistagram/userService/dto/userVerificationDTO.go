package dto

import "XWS-Nistagram-2021/backend-nistagram/userService/model"

type UserVerificationDTO struct {
	UserId string `json:"_id"`
	VerificationType model.UserType `json:"verificationType"`
}
