package main

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/handler"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/repository"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initAuthenticationRepository() *repository.AuthenticationRepository {
	return &repository.AuthenticationRepository{}
}

func initAuthenticationService(repository *repository.AuthenticationRepository) *service.AuthenticationService {
	return &service.AuthenticationService{AuthenticationRepository: repository}
}

func initAuthentcationHandler(service *service.AuthenticationService) *handler.AuthenticationHandler {
	return &handler.AuthenticationHandler{AuthenticationService: service}
}

func handleFunc(handler *handler.AuthenticationHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", handler.Hello).Methods("GET")

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), router))
}

func main() {
	authenticationRepository := initAuthenticationRepository()
	authenticationService := initAuthenticationService(authenticationRepository)
	authenticationHandler := initAuthentcationHandler(authenticationService)
	handleFunc(authenticationHandler)
}
