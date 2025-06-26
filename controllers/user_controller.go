package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang.git/models"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang.git/repository"
	"github.com/gorilla/mux"
)



func CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.User

	_=json.NewDecoder(r.Body).Decode(&user)
	err := repository.CreateUser(&user)
	if err != nil{
		http.Error(w,err.Error(),500)
		return 
	}
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	var users []models.User
	err := repository.GetAllUsers(&users)
	if err != nil{
		http.Error(w,err.Error(),500)
		return 
	}
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request){
	var user models.User
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := repository.GetUserByID(id, &user)

	if err != nil{
		http.Error(w,err.Error(),500)
		return
	}
	json.NewEncoder(w).Encode(user)
}



