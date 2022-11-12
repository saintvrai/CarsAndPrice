package main

import (
	"fmt"
	"github.com/gorilla/mux"
)

type Car struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	CarBrand *CarBrand `json:"carBrand"`
}
type CarBrand struct {
	Name string `json:"name"`
}

var cars []Car

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cars", getCars).Methods("GET")
	r.HandleFunc("/cars/{id}", getCars).Methods("GET")
	r.HandleFunc("/cars", createCar).Methods("POST")
	r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/cars", deleteCar).Methods("DELETE")
	fmt.Printf("Hello World")
	fmt.Printf("Hello World")
}
