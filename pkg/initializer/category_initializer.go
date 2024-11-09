package initializer

import (
	"database/sql"
	"log"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/controller/api"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/database"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
)

type Categories struct {
	DB      *sql.DB
	Repo    *apirepository.CategoryRepoDB
	Service *apiservice.ServiceCategory
	Handler *api.CategoryHandler
}

func InitCategory() *Categories {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	// Inisialisasi repository
	catRepo := apirepository.NewCategoryRepo(db)

	// Inisialisasi service
	catService := apiservice.NewServiceCategory(catRepo)

	// Inisialisasi handler
	catHandler := api.NewCategoryHandler(catService)

	return &Categories{
		DB:      db,
		Repo:    catRepo,
		Service: catService,
		Handler: catHandler,
	}
}