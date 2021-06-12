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
	cors "github.com/rs/cors"
	"os"
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

	c := SetupCors()

	http.Handle("/", c.Handler(router))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(router))
	if err != nil {
		log.Println(err)
	}
}

func initDatabase() *gorm.DB {
	var database *gorm.DB
	err := godotenv.Load()
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))
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

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
}

func main() {
	database := initDatabase()
	authenticationRepository := initAuthenticationRepository(database)
	authenticationService := initAuthenticationService(authenticationRepository)
	authenticationHandler := initAuthentcationHandler(authenticationService)

	handleFunc(authenticationHandler)
}
