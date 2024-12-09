package main

import (
	"calculator/internal/api"
	"calculator/internal/repository"
	"calculator/internal/service/calculator"
	"log"
	"net/http"
)

func main() {
	calcRepo := repository.New()

	calc := calculator.New(calcRepo)

	handler := api.New(calc)

	http.HandleFunc("/add", handler.Add)

	http.HandleFunc("/subtraction", handler.Subtraction)

	if err := http.ListenAndServe("localhost:9001", nil); err != nil {
		log.Fatal(err)
	}
}
