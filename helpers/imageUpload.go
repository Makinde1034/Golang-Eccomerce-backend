package helper

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/joho/godotenv"
)

func UploadImage(input interface{})(string, error){
	ctx, cancel := context.WithTimeout(context.Background(),20*time.Second)
	defer cancel()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("failed to load env")
	}

	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudApiKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudSecretKey := os.Getenv("CLOUDINARY_API_SECRET")
	cloudFolder := os.Getenv("CLOUDINARY_FOLDER")

	cld,err := cloudinary.NewFromParams(cloudName,cloudApiKey,cloudSecretKey)

	if err != nil {
		return "", err
	}

	uploadparam, err := cld.Upload.Upload(ctx,input,uploader.UploadParams{Folder: cloudFolder})

	if err != nil {
		return "",err
	}

	return uploadparam.SecureURL,nil

}