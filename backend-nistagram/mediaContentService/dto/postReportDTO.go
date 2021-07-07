package dto

type PostReportDTO struct {
	Text string `json:"text"`
	PostId string `json:"postId"`
	Username string `json:"username"`
}