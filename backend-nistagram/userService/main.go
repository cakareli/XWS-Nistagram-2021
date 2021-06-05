package main

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello from USER SERVICE!")
	})

	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	user := model.User{
		Id: 1,
		Name: "Dimitrije",
		Surname: "Bulaja",
		Gender: model.Male,
	}

	userDatabase := client.Database("user")
	userCollection := userDatabase.Collection("users")

	userResult, err := userCollection.InsertOne(context.TODO(), bson.D{
		{Key: "user", Value: user},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*userResult)

	filterCursor, err := userCollection.Find(context.TODO(), bson.M{"user.surname": "Bulaja"})
	if err != nil {
		log.Fatal(err)
	}

	var usersFiltered []bson.M
	if err = filterCursor.All(context.TODO(), &usersFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(usersFiltered)


	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
