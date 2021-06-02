package model

type Follower struct {
	UserFollower RegularUser `json:"userFollower"`
	IsClose bool `json:"isClose"`
	IsAccepted bool `json:"isAccepted"`

}