package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	u "server/pb"

	"google.golang.org/protobuf/proto"
)

type UserJSON struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	IsActive bool    `json:"isActive"`
	Score    float32 `json:"score"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	var user UserJSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func protoHandler(w http.ResponseWriter, r *http.Request) {
	var user u.User
	body, _ := io.ReadAll(r.Body)
	err := proto.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	data, _ := proto.Marshal(&user)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/proto", protoHandler)
	log.Println("Listening on Port :8080")
	http.ListenAndServe(":8080", nil)
}
