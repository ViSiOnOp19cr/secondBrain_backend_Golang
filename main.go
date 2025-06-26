package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/config"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/models"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/routes"
	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{}, &models.Content{})

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	routes.RegisterUserRoutes(api)
	routes.RegisterContentRoutes(api)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
