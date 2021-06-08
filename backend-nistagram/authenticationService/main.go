package main

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/handler"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/repository"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

func initAuthenticationRepository(database *gorm.DB) *repository.AuthenticationRepository {
	return &repository.AuthenticationRepository{Database: database}
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
	router.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/update", handler.UpdateUser).Methods("POST")

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), router))
}

func initDatabase() *gorm.DB {
	var database *gorm.DB
	err := godotenv.Load()
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"postgres", "1234567", "auth-service", "5432")

	log.Print("Connecting to PostgreSQL DB...")
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&model.User{})

	return database
}

func main() {
	database := initDatabase()
	authenticationRepository := initAuthenticationRepository(database)
	authenticationService := initAuthenticationService(authenticationRepository)
	authenticationHandler := initAuthentcationHandler(authenticationService)

	handleFunc(authenticationHandler)
}
