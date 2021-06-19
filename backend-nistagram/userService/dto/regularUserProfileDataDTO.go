package dto

type RegularUserProfileDataDTO struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Biography string `json:"biography"`
	WebSite string `json:"webSite"`
}