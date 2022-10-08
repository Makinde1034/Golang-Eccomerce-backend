package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateStroe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Got to create store")
	userId,ok := r.Context().Value("result").(string)

	if ok {
		fmt.Println(userId)
	}
}
