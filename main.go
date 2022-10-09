package main

import (
	"fmt"

	"music-app-backend/configs"
	// "music-app-backend/routes"
	"music-app-backend/controllers"
	"music-app-backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	
	configs.ConnectDB()

	mux := mux.NewRouter()
	// create subrouter to check token status
	secure := mux.PathPrefix("/").Subrouter()
	secure.Use(middlewares.Authentication)
	mux.HandleFunc("/register", controllers.Register).Methods("POST")
	mux.HandleFunc("/login", controllers.Login).Methods("POST")


	secure.HandleFunc("/test", controllers.CreateStroe).Methods("GET")

	http.ListenAndServe(":8000", mux)

	fmt.Println("listening")
	



}
