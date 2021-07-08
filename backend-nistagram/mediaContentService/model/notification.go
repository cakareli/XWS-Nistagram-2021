package model

type Notification struct {
	EventId string `bson:"eventId,omitempty"`
	UserId string `bson:"userId,omitempty"`
	ReceiversIds []string `bson:"receiversIds,omitempty"`
	NotificationType NotificationType `bson:"notificationType"`
}