package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Car struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	CarBrand *CarBrand `json:"carBrand"`
	CarColor string    `json:"carColor"`
}
type CarBrand struct {
	Name string `json:"name"`
}

var cars []Car

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			json.NewEncoder(w).Encode(cars)
			break
		}
	}
}

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(100000))
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)

}

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			var car Car
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(car)
			return
		}
	}
}
func main() {
	r := mux.NewRouter()

	cars = append(cars, Car{
		ID:       "1",
		Name:     "Mark2",
		CarBrand: &CarBrand{Name: "Toyota"},
		CarColor: "Black",
	})

	cars = append(cars, Car{
		ID:       "2",
		Name:     "President",
		CarBrand: &CarBrand{Name: "Nissan"},
		CarColor: "White",
	})
	r.HandleFunc("/cars", getCars).Methods("GET")
	r.HandleFunc("/cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/cars", createCar).Methods("POST")
	r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/cars", deleteCar).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
