package helper

import (
	"fmt"
	"net/smtp"
)

// "net/http"

// "os"
// "github.com/joho/godotenv"
// "go.mongodb.org/mongo-driver/mongo/address"
// "golang.org/x/text/message"


func Email(body string) string{

	response := ""

	

	msg := "Subject : Verification\n my email body"

	auth := smtp.PlainAuth("","makinde1034@gmail.com","jmlgqmlzabktgcbp","smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587",auth,"makinde1034@gmail.com",[]string{"makinde1034@gmail.com"},[]byte(msg))

	if err != nil {
		fmt.Println(err)
		return "An error occured"
	}
	response = "email sent"
	return  response


}