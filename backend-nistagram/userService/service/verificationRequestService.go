package service

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"XWS-Nistagram-2021/backend-nistagram/userService/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

func (service *VerificationRequestService) GetAllVerificationRequests() ([]dto.VerificationDTO, error){
	allVerification,err := service.VerificationRequestRepository.GetAllVerificationRequests()
	if err != nil {
		return nil, err
	}

	allVerificationModel := createVerificationRequestFromDocuments(allVerification)

	allVerificationDto := createVerificationRequestDtoFromVerificationRequest(allVerificationModel)

	return allVerificationDto,nil
}

func (service *VerificationRequestService) DeleteVerificationRequest(id primitive.ObjectID) error{
	err := service.VerificationRequestRepository.DeleteVerificationRequest(id)
	if err != nil {
		return err
	}
	return nil
}

func createVerificationRequestFromDocuments(VerificationRequestDocuments []bson.D) []model.VerificationRequest {
	var verificationRequests []model.VerificationRequest
	for i := 0; i < len(VerificationRequestDocuments); i++ {
		var verificationRequest model.VerificationRequest
		bsonBytes, _ := bson.Marshal(VerificationRequestDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &verificationRequest)
		fmt.Println("U documentsu: ",verificationRequest.Id)
		verificationRequests = append(verificationRequests, verificationRequest)
	}
	return verificationRequests
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

func createVerificationRequestDtoFromVerificationRequest(verificationRequestsModel []model.VerificationRequest) []dto.VerificationDTO{
	var verificationRequestsDto []dto.VerificationDTO
	for i := 0; i < len(verificationRequestsModel); i++{
		var verificationRequestDto dto.VerificationDTO
		fmt.Println("Id u pomocnoj(obican): ",verificationRequestsModel[i].Id)
		fmt.Println("Id u pomocnoj(Hex): ",verificationRequestsModel[i].Id.Hex())
		verificationRequestDto.Id = verificationRequestsModel[i].Id.Hex()
		verificationRequestDto.UserId = verificationRequestsModel[i].UserId
		verificationRequestDto.Name = verificationRequestsModel[i].Name
		verificationRequestDto.Surname = verificationRequestsModel[i].Surname
		verificationRequestDto.ImageUrl = verificationRequestsModel[i].ImageUrl
		verificationRequestDto.VerificationType = verificationRequestsModel[i].VerificationType
		verificationRequestsDto = append(verificationRequestsDto, verificationRequestDto)
	}
	return verificationRequestsDto
}