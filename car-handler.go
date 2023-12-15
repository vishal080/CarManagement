package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Car struct {
	ID    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
}

func getCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars := getCarsFromDB()

	carsJSON, err := json.Marshal(cars)
	if err != nil {
		log.Println("Error marshaling cars to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(carsJSON)
}

func getCarsFromDB() []Car {
	var cars []Car

	rows, err := db.Query("SELECT id, make, model FROM cars")
	if err != nil {
		log.Println("Error querying database:", err)
		return cars
	}
	defer rows.Close()

	for rows.Next() {
		var car Car
		err := rows.Scan(&car.ID, &car.Make, &car.Model)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		cars = append(cars, car)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error iterating rows:", err)
	}

	return cars
}
