package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/service"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerificationRequestHandler struct {
	VerificationRequestService *service.VerificationRequestService
}

func (handler *VerificationRequestHandler) CreateVerificationRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var verificationRequestDto dto.VerificationRequestDTO
	err := json.NewDecoder(r.Body).Decode(&verificationRequestDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.VerificationRequestService.CreateVerificationRequest(verificationRequestDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *VerificationRequestHandler) GetAllVerificationRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	allVerificationRequestsDto, err := handler.VerificationRequestService.GetAllVerificationRequests()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allVerificationRequestsDto)
	}
}

func (handler *VerificationRequestHandler) DeleteVerificationRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	verificationId,err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.VerificationRequestService.DeleteVerificationRequest(verificationId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}