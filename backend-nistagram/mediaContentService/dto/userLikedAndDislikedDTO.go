package dto

type UserLikedAndDislikedDTO struct {
	LikedPostsIds []string `json:"likedPosts"`
	DislikedPostsIds []string `json:"dislikedPosts"`
}