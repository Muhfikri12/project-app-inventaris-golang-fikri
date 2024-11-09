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
	cat := initializer.InitCategory()

	defer app.DB.Close()
	defer cat.DB.Close()

	router := chi.NewRouter()

	route.ItemsRoute(router, app.ItemHandler)
	route.CategoryRoute(router, cat.Handler)

	log.Println("Server berjalan di port :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
