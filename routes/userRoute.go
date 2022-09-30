package routes

import (
	"github.com/gorilla/mux"
	"music-app-backend/controllers"
)


func RegisterUserRoute(){
	mux := mux.NewRouter()

	mux.HandleFunc("/register",controllers.Register)
}