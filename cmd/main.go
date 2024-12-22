package main

import (
	"log"
	"net/http"

	"github.com/stuymedova/eval/internal/handler"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handler.CalcHandler)
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
