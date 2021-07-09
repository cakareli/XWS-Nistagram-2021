package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

func (repository *NotificationRepository) FindAllNotifications() []bson.D{
	notificationsCollection := repository.Database.Collection("notifications")
	filterCursor, err := notificationsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var notificationsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &notificationsFiltered); err != nil {
		log.Fatal(err)
	}
	return notificationsFiltered
}