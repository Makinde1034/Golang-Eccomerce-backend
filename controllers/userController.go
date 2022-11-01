package controllers

import (
	// "music-app-backend/configs"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Ok string `json:"ok"`
	
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

func JSONError(w http.ResponseWriter, err interface{}, code int) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(err)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newUser models.User

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	json.NewDecoder(r.Body).Decode(&newUser)

	validationErr := validate.Struct(newUser)

	if validationErr != nil {
		json.NewEncoder(w).Encode(validationErr.Error())
		return
	}

	if len(newUser.Password) < 5  {
		JSONError(w, er{"Password is too short","failed"},400)

		return
	}

	userCollection := configs.OpenCollection(configs.DB, "user")

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": newUser.Email})

	if err != nil {
		fmt.Println(err)
	}

	if count > 0 {
		// json.NewEncoder(w).Encode(er{"User already exists."})
		JSONError(w, er{"User already exists","failed"},400)
		return
	}

	hashedPassword := hashPassword(newUser.Password)

	newUser.Password = hashedPassword

	

	newUser.Verified = false
 
	result, err := userCollection.InsertOne(ctx, newUser)


	if err != nil {
		// json.NewEncoder(w).Encode(er{"An error occured"})
		log.Fatal(err)
		return
	}

	token, _, _ := helper.GenerateAllTokens(newUser.Email, newUser.Firstname, newUser.Lastname,result.InsertedID)

	json.NewEncoder(w).Encode(response.RegisterResponse{newUser.Firstname, newUser.Lastname, newUser.Email, token,result.InsertedID})

	msg := helper.Email(newUser.Email)

	fmt.Println(result.InsertedID,msg)

}

func Login(w http.ResponseWriter, r *http.Request) {   
	w.Header().Set("Access-Control-Allow-Origin", "*")  
	var user models.User
	var foundUser models.User

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userCollection := configs.OpenCollection(configs.DB, "user")

	json.NewDecoder(r.Body).Decode(&user)

	err := userCollection.FindOne(ctx, bson.D{{"email", user.Email}}).Decode(&foundUser)

	if err != nil {
		json.NewEncoder(w).Encode(er{"Incorrect Email or password","failed"})
		fmt.Println(err)
		return
	}

	passwordValid, msg := verifyPassword(user.Password, foundUser.Password)

	if passwordValid != true {
		// json.NewEncoder(w).Encode(er{Msg: msg})
		JSONError(w,er{Msg: msg,Ok:"failed"},403)
		return
	}

	token, _, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.Firstname,foundUser.Lastname,foundUser.ID)

	json.NewEncoder(w).Encode(response.RegisterResponse{foundUser.Firstname, foundUser.Lastname, foundUser.Email, token,foundUser.ID})

	fmt.Println(foundUser)
}


func GoogleAuthentication(w http.ResponseWriter, r *http.Request) {
	googleConfig := configs.GoogleAuthSetup()
	URL := googleConfig.AuthCodeURL("randomState")

	// redirect to google login page
	http.Redirect(w,r,URL, http.StatusSeeOther)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request){
	state:= r.URL.Query()["state"][0]
	if state != "randomState" {
		fmt.Println("state do not match")
		return
	}

	code := r.URL.Query()["code"][0]

	googleConfig := configs.GoogleAuthSetup()

	token,err := googleConfig.Exchange(context.Background(),code)

	if err != nil {
		fmt.Println("code exchange failed")
	}

	//fetch user details from google API
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Println("Err while fetching user details")
	}

	// PARSE USER DATA AS JSON

	userData,err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	fmt.Print(string(userData))
}