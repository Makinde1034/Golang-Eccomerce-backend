package main

import (
	"fmt"

	"music-app-backend/configs"
	// "music-app-backend/routes"
	"github.com/gorilla/mux"
	"net/http"
	"music-app-backend/controllers"
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

	mux.HandleFunc("/", controllers.Register).Methods("POST")

	http.ListenAndServe(":8000", mux)

	fmt.Println("listening")



}
