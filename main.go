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

	// musicAppDatabase := client.Database("Music-app")

	// userCollection := musicAppDatabase.Collection("user")

	// newUser, err := userCollection.InsertOne(ctx, bson.D{
	// 	{Key: "name",Value : "Toluwalope"},
	// })

	// if err != nil {
	// 	log.Panic(err, "An error occured")
	// }

	// fmt.Print(newUser)
	// var user bson.D
	// errr := userCollection.FindOne(ctx, bson.D{{"name","Toluwalope"}}).Decode(&user)

	// if errr != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println(user)
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
