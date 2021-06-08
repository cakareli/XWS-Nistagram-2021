package dto

import "XWS-Nistagram-2021/backend-nistagram/authenticationService/model"

type LoginResponseDTO struct {
	UserId string `json:"userId"`
	Token string `json:"token"`
	Role model.UserRole `json:"role"`
}
