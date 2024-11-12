package initializer

import (
	"database/sql"
	"log"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/controller/api"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/database"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
)

type Authentication struct {
	DB      *sql.DB
	Repo    *apirepository.AuthRepo
	Service *apiservice.AuthService
	Handler *api.AuthHandler
}

func InitAuth() *Authentication {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	// Inisialisasi repository
	repo := apirepository.NewAuthRepo(db)

	// Inisialisasi service
	service := apiservice.NewAuthService(repo)

	// Inisialisasi handler
	handler := api.NewAuthHandler(service)

	return &Authentication{
		DB:      db,
		Repo:    repo,
		Service: service,
		Handler: handler,
	}
}