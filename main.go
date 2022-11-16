package main

import (
	"fmt"

	"music-app-backend/configs"
	// "music-app-backend/routes"
	"music-app-backend/controllers"
	"music-app-backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	
	configs.ConnectDB()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	mux := mux.NewRouter()
	// create subrouter to check token status
	secure := mux.PathPrefix("/").Subrouter()
	secure.Use(middlewares.Authentication)
	mux.HandleFunc("/register", controllers.Register).Methods("POST")
	mux.HandleFunc("/login", controllers.Login).Methods("POST")
	mux.HandleFunc("/googleAuth", controllers.GoogleAuthentication).Methods("GET")
	mux.HandleFunc("/google/googlecallback", controllers.GoogleCallback).Methods("GET")
	mux.HandleFunc("/get-stores", controllers.GetStores).Methods("GET")
	mux.HandleFunc("/image-uploads", controllers.ImageUpload).Methods("POST")

	// secure routess
	secure.HandleFunc("/upload-product",controllers.AddProductToStore).Methods("POST")
	secure.HandleFunc("/ccreate-store", controllers.CreateStore).Methods("POST")

	http.ListenAndServe(":8000", c.Handler(mux))

	fmt.Println("listening")
	



}
