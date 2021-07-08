package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type NotificationHandler struct {
	NotificationService *service.NotificationService
}

func (handler *NotificationHandler) CreateNewNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var notification model.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.NotificationService.CreateNewNotification(notification)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *NotificationHandler) GetAllNotificationsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	notificationsDTOs, err := handler.NotificationService.GetAllNotificationsByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(notificationsDTOs)
	}
}