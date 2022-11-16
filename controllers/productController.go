package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"music-app-backend/configs"
	"music-app-backend/response"
	"net/http"
	"time"

	"github.com/neox5/go-formdata"
	helper "main.go/helpers"
	"main.go/models"
)


func AddProductToStore(w http.ResponseWriter, r *http.Request) {    
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type","multipart/form-data")


	ctx,cancel := context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()

	// userId,ok := r.Context().Value("result").(string)
	
	// if !ok {
	// 	log.Panic("An error occured while uploading products")
	// 	json.NewEncoder(w).Encode(response.Error{"failed","An error occured while uploading products"})

	// }
	
	

	

	var urls []string

	fd, err := formdata.Parse(r)
	fmt.Println(fd)

    if err == formdata.ErrNotMultipartFormData {
     	log.Panic("unsupported file")
      	return
    }

    if err != nil {
		json.NewEncoder(w).Encode(response.Error{"An error occured.",false})
      log.Panic("Error while parsing file")
	  
      return
    }

	name := fd.Get("name").First()
	price := fd.Get("price").First()
	category := fd.Get("category").First()
	color := fd.Get("color").First()
	description := fd.Get("description").First()
	amount := fd.Get("amount").First()

	

	if fd.FileExists("file") {
		for _, file := range fd.GetFile("file") {
		  

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

			urls = append(urls, url)
			

			
		}
		
	}
	product := models.Product{name,price,category,color,description,amount,urls} 

	productCollection :=  configs.OpenCollection(configs.DB, "product")

	result,err := productCollection.InsertOne(ctx,product)

	if(err != nil){
		fmt.Println("An error occred whiled inserting document")
		return
	}

	fmt.Println(result)
	  
	json.NewEncoder(w).Encode(product)


}