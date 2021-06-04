package model

type NotificationType int

const(
	Like NotificationType = iota
	Dislike
	Follow
	Message
)