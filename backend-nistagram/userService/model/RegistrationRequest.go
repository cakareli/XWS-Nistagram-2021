package model

import "time"

type RegistrationRequest struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Gender Gender `json:"gender"`
	BirthDate time.Time `json:"birthDate"`
	Biography string `json:"biography"`
	WebSite string `json:"webSite"`
}