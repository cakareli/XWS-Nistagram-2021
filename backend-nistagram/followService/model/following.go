package model

type Following struct {
	UserFollowing RegularUser `json:"UserFollowing"`
	IsMuted bool `json:"isMuted"`
	Notifying bool `json:"notifying"`
}