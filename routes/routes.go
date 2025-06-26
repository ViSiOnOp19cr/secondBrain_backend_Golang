package routes


import (
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router){
	r.HandleFunc("/users",controllers.CreateUser).Methods("POST")
}