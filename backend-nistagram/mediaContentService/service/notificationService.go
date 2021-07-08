package service

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
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

func (service *NotificationService) GetAllNotificationsByUserId(userId string) ([]dto.NotificationDTO, error) {
	regularUserNotificationsDocuments := service.NotificationRepository.FindAllNotifications()

	regularUserNotifications := CreateNotificationsFromDocuments(regularUserNotificationsDocuments)

	var notificationDTOs []dto.NotificationDTO
	for i := 0; i < len(regularUserNotifications); i++ {
		alreadyInList := false
		if contains(regularUserNotifications[i].ReceiversIds, userId) {
			notificationDTO, err := CreateNotificationDTOFromNotification(regularUserNotifications[i])
			if err != nil {
				return nil, err
			}
			for j := 0; j < len(notificationDTOs); j++ {
				if notificationDTOs[j].MediaContentId == regularUserNotifications[i].EventId &&
					notificationDTOs[j].UserId == regularUserNotifications[i].UserId &&
					notificationDTOs[j].NotificationType == regularUserNotifications[i].NotificationType {
					alreadyInList = true
				}
			}
			if alreadyInList == false {
				notificationDTOs = append(notificationDTOs, *notificationDTO)
			}
		}
	}

	return notificationDTOs, nil
}

func CreateNotificationFromEvent(eventId string, userId string, receiversIds []string, notificationType model.NotificationType) *model.Notification {
	var notification model.Notification
	notification.EventId = eventId
	notification.UserId = userId
	notification.ReceiversIds = receiversIds
	notification.NotificationType = notificationType

	return &notification
}

func CreateNotificationDTOFromNotification(notification model.Notification) (*dto.NotificationDTO, error) {
	userDTO, err := GetUserByUserId(notification.UserId)
	if err != nil {
		return nil, err
	}

	var notificationDTO dto.NotificationDTO
	notificationDTO.MediaContentId = notification.EventId
	notificationDTO.NotificationType = notification.NotificationType
	notificationDTO.UserId = userDTO[0].UserId
	if notification.NotificationType == model.NotificationType(0) {
		notificationDTO.Text = "User " + userDTO[0].Username + " has a new post!"
	} else if notification.NotificationType == model.NotificationType(1) {
		notificationDTO.Text = "User " + userDTO[0].Username + " has a new story!"
	} else if notification.NotificationType == model.NotificationType(2) {
		notificationDTO.Text = "User " + userDTO[0].Username + " commented your post!"
	} else if notification.NotificationType == model.NotificationType(3) {
		notificationDTO.Text = "User " + userDTO[0].Username + " liked your post!"
	} else if notification.NotificationType == model.NotificationType(4) {
		notificationDTO.Text = "User " + userDTO[0].Username + " disliked your post!"
	} else {
		notificationDTO.Text = ""
	}

	return &notificationDTO, nil
}

func CreateNotificationsFromDocuments(NotificationDocuments []bson.D) []model.Notification {
	var notifications []model.Notification
	for i := 0; i < len(NotificationDocuments); i++ {
		var notification model.Notification
		bsonBytes, _ := bson.Marshal(NotificationDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &notification)
		notifications = append(notifications, notification)
	}
	return notifications
}

func GetUserByUserId(id string) ([]dto.UserDTO, error) {
	var userDTOs []dto.UserDTO
	var userIds []string
	userIds = append(userIds, id)
	postBody, _ := json.Marshal(userIds)
	requestUrl := fmt.Sprintf("http://%s:%s/by-users-ids", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&userDTOs)
	return userDTOs, nil
}