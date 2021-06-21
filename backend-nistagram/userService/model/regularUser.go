package model

type RegularUser struct {
	User `bson:",inline"`
	IsDisabled bool `bson:"isDisabled,omitempty"`
	BlockedUsers []int `bson:"regularUser,omitempty"`
	UserType UserType `bson:"userType,omitempty"`
	LikedPosts []string `bson:"likedPosts,omitempty"`
	DislikedPosts []string `bson:"dislikedPosts,omitempty"`
	ProfilePrivacy ProfilePrivacy `bson:"profilePrivacy,omitempty"`
	Notifications []Notification `bson:"notifications,omitempty"`
	MediaContents []MediaContent `bson:"mediaContents,omitempty"`
	Collections []Collection `bson:"collection,omitempty"`
	Followings []Following `bson:"followings,omitempty"`
	Followers []Follower `bson:"followers,omitempty"`
}