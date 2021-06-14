package main

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/handler"
	"XWS-Nistagram-2021/backend-nistagram/userService/repository"
	"XWS-Nistagram-2021/backend-nistagram/userService/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func initUserRepository(database *mongo.Database) *repository.RegularUserRepository {
	return &repository.RegularUserRepository{Database: database}
}

func initUserService(repository *repository.RegularUserRepository) *service.RegularUserService {
	return &service.RegularUserService{RegularUserRepository: repository}
}

func initUserHandler(service *service.RegularUserService) *handler.RegularUserHandler {
	return &handler.RegularUserHandler{RegularUserService: service}
}

func handleFunc(handler *handler.RegularUserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/register-regular-user", handler.Register).Methods("POST")
	router.HandleFunc("/update-regular-user", handler.Update).Methods("PUT")

	c := SetupCors()

	http.Handle("/", c.Handler(router))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(router))
	if err != nil {
		log.Println(err)
	}
}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
}

func initDatabase() *mongo.Database{
	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	userDatabase := client.Database("user")

	return userDatabase
}

func main() {
	userDatabase := initDatabase()
	userRepository := initUserRepository(userDatabase)
	userService := initUserService(userRepository)
	userHandler := initUserHandler(userService)

	handleFunc(userHandler)
}