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
	"golang.org/x/crypto/bcrypt"
	helper "main.go/helpers"
)

type er struct {
	Msg string `json:"msg"`
}

var validate = validator.New()

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func verifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		check = false
		msg = "Incorrect password"
	}

	return check, msg
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

	userCollection := configs.OpenCollection(configs.DB, "user")

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

	token, _, _ := helper.GenerateAllTokens(newUser.Email, newUser.Firstname, newUser.Lastname)

	newUser.Verified = false

	result, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		json.NewEncoder(w).Encode(er{"An error occured"})
	}

	json.NewEncoder(w).Encode(response.RegisterResponse{newUser.Firstname, newUser.Lastname, newUser.Email, token})

	fmt.Println(result)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var foundUser models.User

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userCollection := configs.OpenCollection(configs.DB, "user")

	json.NewDecoder(r.Body).Decode(&user)

	err := userCollection.FindOne(ctx, bson.D{{"email", user.Email}}).Decode(&foundUser)

	if err != nil {
		json.NewEncoder(w).Encode(er{"Incorrect Eemail or password"})
		fmt.Println(err)
		return
	}

	passwordValid, msg := verifyPassword(user.Password, foundUser.Password)

	if passwordValid != true {
		json.NewEncoder(w).Encode(er{Msg: msg})
		return
	}

	token, _, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.Firstname,foundUser.Lastname)

	json.NewEncoder(w).Encode(response.RegisterResponse{foundUser.Firstname, foundUser.Lastname, foundUser.Email, token})

	fmt.Println(foundUser)
}
