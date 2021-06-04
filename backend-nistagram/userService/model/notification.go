package model

import "time"

type Notification struct {
	Text string `json:"text"`
	Notifier RegularUser `json:"notifier"`
	NotificationDate time.Time `json:"notificationTime"`
	NotificationType NotificationType `json:"notificationType"`
}