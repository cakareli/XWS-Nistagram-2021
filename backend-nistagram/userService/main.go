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

func initVerificationRequestRepository(database *mongo.Database) *repository.VerificationRequestRepository {
	return &repository.VerificationRequestRepository{Database: database}
}

func initVerificationRequestService(repository1 *repository.VerificationRequestRepository, repository2 *repository.RegularUserRepository ) *service.VerificationRequestService {
	return &service.VerificationRequestService{VerificationRequestRepository: repository1, RegularUserRepository: repository2}
}

func initVerificationRequestHandler(service *service.VerificationRequestService) *handler.VerificationRequestHandler {
	return &handler.VerificationRequestHandler{VerificationRequestService: service}
}

func handleFunc(userHandler *handler.RegularUserHandler, verificationRequestHandler *handler.VerificationRequestHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/register-regular-user", userHandler.Register).Methods("POST")
	router.HandleFunc("/update-regular-user", userHandler.UpdatePersonalInformations).Methods("PUT")
	router.HandleFunc("/update-profile-privacy", userHandler.UpdateProfilePrivacy).Methods("PUT")
	router.HandleFunc("/logged-user/{id}", userHandler.FindUserById).Methods("GET")
	router.HandleFunc("/by-username/{username}", userHandler.CreateRegularUserPostDTOByUsername).Methods("GET")
	router.HandleFunc("/regular-user-by-username/{username}", userHandler.FindRegularUserByUsername).Methods("GET")
	router.HandleFunc("/public-regular-users", userHandler.GetAllPublicRegularUsers).Methods("GET")
	router.HandleFunc("/search-public-regular-users/{searchInput}", userHandler.GetUserSearchResults).Methods("GET")
	router.HandleFunc("/by-users-ids", userHandler.FindUsersByIds).Methods("POST")
	router.HandleFunc("/verification-request", verificationRequestHandler.CreateVerificationRequest).Methods("POST")
	router.HandleFunc("/liked-and-disliked/{username}", userHandler.FindRegularUserLikedAndDislikedPosts).Methods("GET")
	router.HandleFunc("/update-liked-posts", userHandler.UpdateLikedPosts).Methods("POST")
	router.HandleFunc("/update-disliked-posts", userHandler.UpdateDislikedPosts).Methods("POST")
	router.HandleFunc("/save-post", userHandler.SavePost).Methods("PUT")
	router.HandleFunc("/saved-posts/{username}", userHandler.GetAllSavedPostsByUsername).Methods("GET")

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

	verificationRequestRepository := initVerificationRequestRepository(userDatabase)
	verificationRequestService := initVerificationRequestService(verificationRequestRepository, userRepository)
	verificationRequestHandler := initVerificationRequestHandler(verificationRequestService)

	handleFunc(userHandler, verificationRequestHandler)
}