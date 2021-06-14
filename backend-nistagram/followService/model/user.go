package model

import (

)

type User struct {
	Id string
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	UserRole UserRole `json:"userRole"`
}

