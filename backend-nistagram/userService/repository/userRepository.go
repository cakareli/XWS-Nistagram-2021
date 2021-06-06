package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository struct {
	Database *mongo.Database
}

func (repository *UserRepository) Hello (){
	fmt.Printf("Hello from Repository");
}

func (repository *UserRepository) CreateRegularUser(user *model.RegularUser) error{
	regularUserCollection := repository.Database.Collection("regularUsers")
		_, err := regularUserCollection.InsertOne(context.TODO(), bson.D{
		{Key: "regularUser", Value: user},
	})
	if err != nil {
		return fmt.Errorf("Regular user is NOT created")
	}
	return nil
}

func (repository *UserRepository) ExistByUsername(username string) bool{
	regularUserCollection := repository.Database.Collection("regularUsers")
	filterCursor, err := regularUserCollection.Find(context.TODO(), bson.M{"regularUser.username": username})
	if err != nil {
		log.Fatal(err)
	}

	var userFiltered bson.D
	if err = filterCursor.All(context.TODO(), &userFiltered); err != nil {
		log.Fatal(err)
	}
	if len(userFiltered) != 0 {
		return true
	}
	return false
}