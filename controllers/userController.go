package controllers

import (
	// "music-app-backend/configs"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"music-app-backend/configs"
	"music-app-backend/models"
	"music-app-backend/response"
	"net/http"
	"time"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"music-app-backend/helpers"
	"golang.org/x/crypto/bcrypt"
)

type er struct{
	Msg string `json:"msg"`
}

var validate = validator.New()

func hashPassword(password string) string{
	bytes, err  := bcrypt.GenerateFromPassword([]byte(password),14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	json.NewDecoder(r.Body).Decode(&newUser)

	validationErr := validate.Struct(newUser)

	if validationErr != nil {
		json.NewEncoder(w).Encode(validationErr.Error())
		return 
	}

	userCollection := configs.OpenCollection(configs.DB,"user")

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": newUser.Email})

	if err != nil {
		fmt.Println(err)
	} 

	if count > 0 {
		json.NewEncoder(w).Encode(er{"User already exists."})
		return
	}

	hashedPassword := hashPassword(newUser.Password)

	newUser.Password = hashedPassword

	token, _,_ := helper.GenerateAllTokens(newUser.Email,newUser.Firstname,newUser.Lastname)

	result,err := userCollection.InsertOne(ctx,newUser)


	if err != nil{
		json.NewEncoder(w).Encode(er{"An error occured"})
	}

	json.NewEncoder(w).Encode(response.RegisterResponse{newUser.Firstname,newUser.Lastname,newUser.Email,token})

	fmt.Println(result)




}

func Login(w http.ResponseWriter, r *http.Request){
	var newUser models.User
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userCollection := configs.OpenCollection(configs.DB,"user")

	json.NewDecoder(r.Body).Decode(&newUser)

	errr := userCollection.FindOne(ctx, bson.D{{"email",userM}}).Decode(&user)
}
