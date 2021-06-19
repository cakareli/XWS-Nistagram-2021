package dto

type NewFollowDTO struct {
	FollowerId string `json:"followerId"`
	FollowedId string `json:"followedId"`
	IsPrivate bool `json:"isPrivate"`
}