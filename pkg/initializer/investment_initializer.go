package initializer

import (
	"database/sql"
	"log"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/controller/api"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/database"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
)


type Investment struct {
	DB      *sql.DB
	Repo    *apirepository.InvestmentRepoDB
	Service *apiservice.InvestmentService
	Handler *api.InvestmentHandler
}

func InitInvestment() *Investment {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	// Inisialisasi repository
	invRepo := apirepository.NewInvestmentRepo(db)

	// Inisialisasi service
	invService := apiservice.NewInvestmentService(invRepo)

	// Inisialisasi handler
	invHandler := api.NewInvestmentHandler(invService)

	return &Investment{
		DB:      db,
		Repo:    invRepo,
		Service: invService,
		Handler: invHandler,
	}
}