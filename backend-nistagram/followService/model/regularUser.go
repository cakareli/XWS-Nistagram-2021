package model

type RegularUser struct {
	User
	IsDisabled bool `json:"IsDisabled"`
	BlockedUsers []int `json:"blockedUsers"`
	UserType UserType `json:"userType"`
	Followings []Following `json:"followings"`
	Followers []Follower `json:"followers"`
}