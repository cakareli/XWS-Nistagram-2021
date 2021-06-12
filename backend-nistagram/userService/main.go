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

func initUserRepository(database *mongo.Database) *repository.UserRepository {
	return &repository.UserRepository{Database: database}
}

func initUserService(repository *repository.UserRepository) *service.UserService {
	return &service.UserService{UserRepository: repository}
}

func initUserHandler(service *service.UserService) *handler.UserHandler {
	return &handler.UserHandler{UserService: service}
}

func handleFunc(handler *handler.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", handler.Hello).Methods("GET")
	router.HandleFunc("/create-regular-user", handler.Create).Methods("POST")
	router.HandleFunc("/update-regular-user", handler.Update).Methods("POST")

	c := SetupCors()

	fmt.Println("Server running...")
	//http.Handle("/", c.Handler(router))
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), c.Handler(router)))

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

func main() {

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

	userDatabase := client.Database("user")

	userRepository := initUserRepository(userDatabase)
	userService := initUserService(userRepository)
	userHandler := initUserHandler(userService)
	handleFunc(userHandler)
}
