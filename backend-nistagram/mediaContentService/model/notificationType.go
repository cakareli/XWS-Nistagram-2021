package model

type NotificationType int

const(
	PostNotification NotificationType = iota
	StoryNotification
	CommentNotification
)