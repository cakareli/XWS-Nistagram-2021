package model

type RegularUser struct {
	User
	IsDisabled bool `json:"isDisabled"`
	RegularUser []RegularUser `json:"regularUser"` //!!!!!!!!!!!!!
	UserType UserType `json:"userType"`
	LikedPosts []Post `json:"likedPosts"`
	DislikedPosts []Post `json:"dislikedPosts"`
	ProfilePrivacy ProfilePrivacy `json:"profilePrivacy"`
	Notifications []Notification `json:"notifications"`
	MediaContents []MediaContent `json:"mediaContents"`
	Collections []Collection `json:"collection"`
	Followings []Following `json:"followings"`
	Followers []Follower `json:"followers"`
}