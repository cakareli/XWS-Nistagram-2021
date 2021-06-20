package dto

type UpdatePostLikeAndDislikeDTO struct {
	Username string `json:"username"`
	PostId string `json:"postId"`
	IsAdd string `json:"isAdd"`
}
