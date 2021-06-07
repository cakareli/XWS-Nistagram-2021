package dto

import "XWS-Nistagram-2021/backend-nistagram/authenticationService/model"

type LoginResponseDTO struct {
	Username string `json:"username"`
	Token string `json:"token"`
	Role model.UserRole `json:"role"`
}
