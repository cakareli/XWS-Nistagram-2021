package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository struct {
	Database *mongo.Database
}

func (repository *NotificationRepository) CreateNotification(notification *model.Notification) error {
	notificationCollection := repository.Database.Collection("notifications")
	_, err := notificationCollection.InsertOne(context.TODO(), &notification)
	if err != nil {
		return fmt.Errorf("notification is NOT created")
	}
	return nil
}