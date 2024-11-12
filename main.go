package main

import (
	"log"
	"net/http"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/pkg/initializer"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/route"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Inisialisasi semua handler
	app := initializer.InitApp()
	cat := initializer.InitCategory()
	inv := initializer.InitInvestment()
	auth := initializer.InitAuth()

	// Tutup koneksi database saat aplikasi dihentikan
	defer func() {
		app.DB.Close()
		cat.DB.Close()
		inv.DB.Close()
		auth.DB.Close()
	}()

	// Buat router baru
	router := chi.NewRouter()

	// Registrasi semua route
	route.RegisterRoutes(router, app.ItemHandler, cat.Handler, inv.Handler, auth.Handler)

	// Jalankan server
	log.Println("Server berjalan di port :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
