package helper

import (
	// "net/http"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/mongo/address"
	// "golang.org/x/text/message"
)

func Email(reciever string) string{

	msg := ""
	

	err := godotenv.Load(".env")

	if err != nil {
		log.Panic(err)
	}

	fromEmail := os.Getenv("fromEmail")
	password := os.Getenv("password")
	toEmail := []string{reciever}
	host := "smtp.gmail.com"
	port := "587"
	address := host+":"+port
	subject := "Please verify your account"
	body := "First email"
	message := []byte(subject + body)
	auth := smtp.PlainAuth("",fromEmail,password,host)

	_err := smtp.SendMail(address,auth,fromEmail,toEmail,message)

	if _err != nil {
		fmt.Println(_err)
		msg = "Somthing went wrong."
		return msg
	}
	msg = "Email sent successfully"

	return msg

	
}