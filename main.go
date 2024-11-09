package main

import (
	"log"
	"net/http"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/pkg/initializer"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/route"
	"github.com/go-chi/chi/v5"
)

func main() {
	app := initializer.InitApp()
	defer app.DB.Close()

	router := chi.NewRouter()

	route.ItemsRoute(router, app.ItemHandler)
	// route.Items(router, app.ItemHandler)
	// route.Items(router, app.ItemHandler)
	
	log.Println("Server berjalan di port ")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}