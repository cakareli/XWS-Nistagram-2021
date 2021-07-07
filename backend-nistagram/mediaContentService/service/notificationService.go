package service

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"fmt"
)

type NotificationService struct {
	NotificationRepository *repository.NotificationRepository
}

func (service *NotificationService) CreateNewNotification(notification model.Notification) error {
	fmt.Println("Creating new notification")

	err := service.NotificationRepository.CreateNotification(&notification)
	if err != nil {
		return err
	}
	return nil
}