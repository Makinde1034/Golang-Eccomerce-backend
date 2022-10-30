package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"music-app-backend/configs"
	"music-app-backend/models"
	"music-app-backend/response"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateStore(w http.ResponseWriter, r *http.Request) {
	var store models.Store	
	userId,ok := r.Context().Value("result").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	json.NewDecoder(r.Body).Decode(&store)

	storeCollection := configs.OpenCollection(configs.DB, "store")

	store.Owner = userId

	result,err := storeCollection.InsertOne(ctx,store)

	if err != nil{
		json.NewEncoder(w).Encode("Failed to create store")
		return
	}

	fmt.Println(result)




	if ok {
		fmt.Println(userId,"herrrrre")
		json.NewEncoder(w).Encode(userId)
	}


}


func GetStores(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
 	var stores []models.Store

	storeCollection := configs.OpenCollection(configs.DB,"store")

	cur,err := storeCollection.Find(context.Background(),bson.D{{}})

	if err != nil {
		json.NewEncoder(w).Encode(response.Error{"Failed to load stores","failed"})
		return
	}

	for cur.Next(context.Background()) {
        //Create a value into which the single document can be decoded
		var store models.Store

        err := cur.Decode(&store)
        if err != nil {
            log.Fatal(err)
        }

        stores =append(stores,store)

    }

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	

	json.NewEncoder(w).Encode(stores)
	fmt.Println(stores)

}