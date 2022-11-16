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

	"github.com/neox5/go-formdata"
	"go.mongodb.org/mongo-driver/bson"
	helper "main.go/helpers"
)


func CreateStore(w http.ResponseWriter, r *http.Request) {
	var store models.Store	
	userId,ok := r.Context().Value("result").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	json.NewDecoder(r.Body).Decode(&store)


	// upload store image 
	formfile, _, err := r.FormFile("file")

	if(err !=nil){
		json.NewEncoder(w).Encode(response.Error{"failed",false})
		return
	}

	url, err := helper.UploadImage(formfile)

	if err != nil {
		json.NewEncoder(w).Encode(response.Error{"failed to upload image",false})
		fmt.Println(err)
		return
	}


	storeCollection := configs.OpenCollection(configs.DB, "store")

	store.Owner = userId
	store.Image = url

	result,err := storeCollection.InsertOne(ctx,store)

	if err != nil{
		json.NewEncoder(w).Encode("Failed to create store")
		return
	}

	fmt.Println(result)




	if ok {
		fmt.Println(userId,"herrrrre")
	}

	json.NewEncoder(w).Encode(store)


}


func GetStores(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
 	var stores []models.Store

	storeCollection := configs.OpenCollection(configs.DB,"store")

	cur,err := storeCollection.Find(context.Background(),bson.D{{}})

	if err != nil {
		json.NewEncoder(w).Encode(response.Error{"Failed to load stores",false})
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
	// fmt.Println(stores)

}                                                                    



func ImageUpload(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type","multipart/form-data")

	fd, err := formdata.Parse(r)
	fmt.Println(fd)

    if err == formdata.ErrNotMultipartFormData {
     log.Panic("unsupported file")
      return
    }
    if err != nil {
      log.Panic("Error while parsing file")
      return
    }

	from := fd.Get("file").First()

	fmt.Println(from)

	if fd.FileExists("file") {
		for _, file := range fd.GetFile("file") {
		  
		  fmt.Println(file.Filename)

		  reader, err := file.Open()

		  if err !=nil {
			log.Panic("error reading files")
		  }

		  	url, err := helper.UploadImage(reader)

			if err != nil {
				json.NewEncoder(w).Encode(response.Error{"failed to upload image",false})
				fmt.Println(err)
				return
			}

			json.NewEncoder(w).Encode(url)
		}
	  }

	

	
	// formfile, _, err := r.FormFile("file")

	// if(err !=nil){
	// 	json.NewEncoder(w).Encode(response.Error{"failed","failed"})  
	// }

	// fmt.Println(formfile) 

	// url, err := helper.UploadImage(file)

	// if err != nil {
	// 	json.NewEncoder(w).Encode(response.Error{"failed to upload image","failed"})
	// 	fmt.Println(err)
	// 	return
	// }

	// json.NewEncoder(w).Encode(url)

}