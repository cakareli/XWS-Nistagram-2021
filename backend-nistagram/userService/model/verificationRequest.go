package model

type VerificationRequest struct {
	Id int `json:"id"`
	UserType UserType `json:"userType"`
	VerificationName string `json:"verificationName"`
	VerificationSurname string `json:"verificationSurname"`
	VerificationImage string `json:"verificationImage"`
	RegularUser RegularUser `json:"regularUser"`
}