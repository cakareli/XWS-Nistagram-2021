package service

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"XWS-Nistagram-2021/backend-nistagram/userService/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerificationRequestService struct {
	VerificationRequestRepository *repository.VerificationRequestRepository
	RegularUserRepository *repository.RegularUserRepository
}

func (service *VerificationRequestService) CreateVerificationRequest(verificationRequestDto dto.VerificationRequestDTO) error {
	fmt.Println("Creating new verification request")

	id, _ := primitive.ObjectIDFromHex(verificationRequestDto.UserId)
	_, err1 := service.RegularUserRepository.FindUserById(id)
	if err1 != nil {
		return err1
	}

	var verificationRequest = createVerificationRequestFromVerificationRequestDto(&verificationRequestDto)
	err2 := service.VerificationRequestRepository.CreateVerificationRequest(verificationRequest)
	if err2 != nil {
		return err2
	}
	return nil
}

func createVerificationRequestFromVerificationRequestDto(verificationRequestDto *dto.VerificationRequestDTO) *model.VerificationRequest{
	var verificationRequest model.VerificationRequest
	verificationRequest.UserId = verificationRequestDto.UserId
	verificationRequest.Name = verificationRequestDto.Name
	verificationRequest.Surname = verificationRequestDto.Surname
	verificationRequest.ImageUrl = verificationRequestDto.ImageUrl
	verificationRequest.VerificationType = verificationRequestDto.VerificationType

	return &verificationRequest
}