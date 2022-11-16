package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/neox5/go-formdata"
	// helper "main.go/helpers"
	helper "main.go/helpers"
	"main.go/response"
)

type errorMsg struct{
	msg string
}


func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		/* Temporary fix for CORS error */

		fd, _err := formdata.Parse(r)
		fmt.Println(fd)
	
		if _err == formdata.ErrNotMultipartFormData {
			 log.Panic("unsupported file")
			  return
		}
	
		if _err != nil {
			json.NewEncoder(w).Encode(response.Error{"An error occured.",false})
		  log.Panic("Error while parsing file")
		  
		  return
		}
	
		token := fd.Get("token").First()
		fmt.Println(token)
		if token == "" {
			json.NewEncoder(w).Encode(response.Error{"No token found",false})
			return
		}
		claims, err := helper.VerifyToken(token) 

		if err != "" {
			json.NewEncoder(w).Encode(response.Error{"Invalid Authorization",false})
			return

			
		}

		userId := r.WithContext(context.WithValue(r.Context(), "result", claims.Uid))
		next.ServeHTTP(w,userId)

		/******************************************/
		
		
		// var header = r.Header.Get("Authorization")
		// fmt.Println(header)
		// if header == "" {
		// 	json.NewEncoder(w).Encode(response.Error{"No token found",false})
		// 	return
		// }
		// claims, err := helper.VerifyToken(header) 

		// if err != "" {
		// 	json.NewEncoder(w).Encode("Invalid Authorization")
		// 	return

			
		// }

		// userId := r.WithContext(context.WithValue(r.Context(), "result", claims.Uid))
		// next.ServeHTTP(w,userId)

	
	})
}  































































































