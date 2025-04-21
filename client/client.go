package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

type UserJSON struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	IsActive bool    `json:"isActive"`
	Score    float32 `json:"score"`
}

func benchmarkJSON(user UserJSON) {
	data, _ := json.Marshal(user)
	start := time.Now()
	resp, err := http.Post("http://localhost:8080/json", "application/json", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("JSON round trip:", time.Since(start))
}

func benchmarkProto(user User) {
	data, _ := proto.Marshal(&user)
	start := time.Now()
	resp, err := http.Post("http://localhost:8080/proto", "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("Proto round trip:", time.Since(start))
}

func main() {
	userProto := User{
		Id:       1,
		Name:     "John Doe",
		Email:    "john@doe.com",
		IsActive: true,
		Score:    100.0,
	}
	userJson := UserJSON{
		Id:       1,
		Name:     "John Doe",
		Email:    "john@doe.com",
		IsActive: true,
		Score:    100.0,
	}
	benchmarkJSON(userJson)
	benchmarkProto(userProto)
}
