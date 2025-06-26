package routes

import (
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {

	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
}

func RegisterContentRoutes(r *mux.Router) {

	r.HandleFunc("/content", controllers.CreateContent).Methods("POST")
	r.HandleFunc("/users/{userId}/content", controllers.GetAllContent).Methods("GET")
	r.HandleFunc("/content/{id}", controllers.GetContentByID).Methods("GET")
	r.HandleFunc("/content/{id}", controllers.UpdateContent).Methods("PUT")
	r.HandleFunc("/content/{id}", controllers.DeleteContent).Methods("DELETE")
}
