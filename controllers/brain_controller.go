package controllers
import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang.git/models"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang.git/repository"
	"github.com/gorilla/mux"
)

func CreateContent(w http.ResponseWriter, r *http.Request){
	var content models.Content
	_ = json.NewDecoder(r.Body).Decode(content)
	err := repository.CreateContent(content)
	if err != nil{
		http.Error(w,err.Error(),500)
	}
	json.NewEncoder(w).Encode(content)
}