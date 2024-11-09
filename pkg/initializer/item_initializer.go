package initializer

import (
	"database/sql"
	"log"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/controller/api"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/database"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
)

type App struct {
	DB          *sql.DB
	ItemRepo    *apirepository.ItemRepoDB
	ItemService *apiservice.ItemService
	ItemHandler *api.ItemApiHandler
}

func InitApp() *App {
	// Koneksi database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	// Inisialisasi repository
	itemRepo := apirepository.NewItemRepo(db)

	// Inisialisasi service
	itemService := apiservice.NewItemService(itemRepo)

	// Inisialisasi handler
	itemHandler := api.NewItemHandler(&itemService)

	return &App{
		DB:          db,
		ItemRepo:    &itemRepo,
		ItemService: &itemService,
		ItemHandler: itemHandler,
	}
}