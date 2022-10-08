package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	helper "main.go/helpers"
)

type errorMsg struct{
	msg string
}


func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var header = r.Header.Get("x-access-token")
		if header == "" {
			json.NewEncoder(w).Encode("No token found")
			return
		}
		claims, err := helper.VerifyToken(header)

		if err != "" {
			json.NewEncoder(w).Encode("Invalid Authorization")

			
		}

		userId := r.WithContext(context.WithValue(r.Context(), "result", claims.Uid))
		next.ServeHTTP(w,userId)

	
	})
}  































































































