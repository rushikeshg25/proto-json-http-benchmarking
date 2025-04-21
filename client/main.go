package main

import (
	"bytes"
	u "client/pb"
	"encoding/json"
	"fmt"
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

func benchmarkJSON(user UserJSON) time.Duration {
	data, _ := json.Marshal(user)
	start := time.Now()
	resp, err := http.Post("http://localhost:8080/json", "application/json", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return time.Since(start)
}

func benchmarkProto(user u.User) time.Duration {
	data, _ := proto.Marshal(&user)
	start := time.Now()
	resp, err := http.Post("http://localhost:8080/proto", "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return time.Since(start)
}

func average(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

func main() {
	userProto := u.User{
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

	var jsonTimes []time.Duration
	var protoTimes []time.Duration

	// Run 3 benchmarks each
	for i := 0; i < 3; i++ {
		jsonTimes = append(jsonTimes, benchmarkJSON(userJson))
		protoTimes = append(protoTimes, benchmarkProto(userProto))
	}

	// Print result in table
	fmt.Println("\n📊 Benchmark Comparison")
	fmt.Println("-------------------------------------------")
	fmt.Printf("| %-10s | %-15s | %-15s |\n", "Round", "JSON", "PROTOBUF")
	fmt.Println("-------------------------------------------")
	for i := 0; i < 3; i++ {
		fmt.Printf("| %-10d | %-15s | %-15s |\n", i+1, jsonTimes[i], protoTimes[i])
	}
	fmt.Println("-------------------------------------------")
	fmt.Printf("| %-10s | %-15s | %-15s |\n", "Average", average(jsonTimes), average(protoTimes))
	fmt.Println("-------------------------------------------")
}
