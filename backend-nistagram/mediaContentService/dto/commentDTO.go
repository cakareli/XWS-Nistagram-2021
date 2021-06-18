package dto

type CommentDTO struct {
	Username string `json:"username"`
	PostId string `json:"postId"`
	Text string `json:"text"`
}
