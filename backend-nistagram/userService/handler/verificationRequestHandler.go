package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/service"
	"encoding/json"
	"net/http"
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