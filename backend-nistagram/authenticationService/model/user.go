package model


type User struct{
	Name string
}

func newUser() *User{
	u := User{Name: "brt"}
	return &u
}
