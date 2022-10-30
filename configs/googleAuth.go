package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleAuthSetup() *oauth2.Config{

	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("An error occured while loading env file")
	}



	conf := &oauth2.Config{
		ClientID:    os.Getenv("googleClientId"),
		ClientSecret: os.Getenv(("googleClientSecret")),
		RedirectURL:  "https://localhost:8000/google/googlecallback",
		Scopes: []string{
			"https://www.googleapis.com/auth/bigquery",
			"https://www.googleapis.com/auth/blogger",
		},
		Endpoint: google.Endpoint,
	
	}
	return conf
}	