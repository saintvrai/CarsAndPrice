package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

func getCars(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func deleteCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars{
		if item.ID == params["id"]{
			cars = append(cars[:index], cars [index+1:]...)
			break
		}
	}
}

func getCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,item := range cars {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return item
		}
	}

}
func main() {
	r := mux.NewRouter()

	cars = append(cars,Car{
		ID:"1",
		Name:     "Mark2",
		CarBrand: &CarBrand{Name: "Toyota"},
		CarColor: "Black",

	})

	cars = append(cars,Car{
		ID:"2",
		Name:     "President",
		CarBrand: &CarBrand{Name: "Nissan"},
		CarColor: "White",

	})

	}
	r.HandleFunc("/cars", getCars).Methods("GET")
	r.HandleFunc("/cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/cars", createCar).Methods("POST")
	r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/cars", deleteCar).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
