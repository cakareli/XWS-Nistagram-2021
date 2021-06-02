package model

import (
	"time"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Gender Gender `json:"gender"`
	BirthDate time.Time `json:"birthDate"`
	UserRole UserRole `json:"userRole"`
	Biography string `json:"biography"`
	WebSite string `json:"webSite"`
}