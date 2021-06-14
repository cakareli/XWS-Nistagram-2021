package main

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/handler"
	"XWS-Nistagram-2021/backend-nistagram/followService/repository"
	"XWS-Nistagram-2021/backend-nistagram/followService/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func initFollowRepository() *repository.FollowRepository {
	return &repository.FollowRepository{}
}

func initFollowService(repository *repository.FollowRepository) *service.FollowService {
	return &service.FollowService{FollowRepository: repository}
}

func initFollowHandler(service *service.FollowService) *handler.FollowHandler {
	return &handler.FollowHandler{FollowService: service}
}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
}

func handleFunc(handler *handler.FollowHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", handler.Hello).Methods("GET")

	c := SetupCors()

	fmt.Println("Server is running...")

	http.Handle("/", c.Handler(router))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(router))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	authenticationRepository := initFollowRepository()
	authenticationService := initFollowService(authenticationRepository)
	authenticationHandler := initFollowHandler(authenticationService)

	handleFunc(authenticationHandler)
}
