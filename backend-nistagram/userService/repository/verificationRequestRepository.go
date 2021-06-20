package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerificationRequestRepository struct {
	Database *mongo.Database
}

func (repository *VerificationRequestRepository) CreateVerificationRequest(verificationRequest *model.VerificationRequest) error {
	verificationRequestCollection := repository.Database.Collection("verificationRequests")
	_, err := verificationRequestCollection.InsertOne(context.TODO(), &verificationRequest)
	if err != nil {
		return fmt.Errorf("verification request is NOT created")
	}
	return nil
}