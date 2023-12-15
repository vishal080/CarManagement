





package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	ID    int
	Make  string
	Model string
}

var db *sql.DB

func main() {
	// Connect to SQLite3 database
	http.HandleFunc("/getCars", getCarsHandler)
	var err error
	db, err = sql.Open("sqlite3", "./cars.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create cars table if it doesn't exist
	createTable()

	http.HandleFunc("/addCar", addCarHandler)
	http.HandleFunc("/updateCar", updateCarHandler)
	http.HandleFunc("/deleteCar", deleteCarHandler)
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS cars (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		make TEXT,
		model TEXT
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func addCarHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	make := r.Form.Get("make")
	model := r.Form.Get("model")

	_, err := db.Exec("INSERT INTO cars (make, model) VALUES (?, ?)", make, model)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func getCarsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT make, model FROM cars")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []Car
	for rows.Next() {
		var make, model string
		if err := rows.Scan(&make, &model); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, Car{Make: make, Model: model})
	}

	// Convert cars data to JSON
	carsJSON, err := json.Marshal(cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(carsJSON)
}

func updateCarHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	make := r.Form.Get("make")
	model := r.Form.Get("model")

	_, err = db.Exec("UPDATE cars SET make = ?, model = ? WHERE id = ?", make, model, id)
	if err != nil {
		log.Fatal(err)
	}

	



	http.Redirect(w, r, "/", http.StatusSeeOther)
}


// func updateCarHandler(w http.ResponseWriter, r *http.Request) {
//     r.ParseForm()
//     idStr := r.Form.Get("id")
//     id, err := strconv.Atoi(idStr)
//     if err != nil {
//         log.Fatal(err)
//     }

//     make := r.Form.Get("make")
//     model := r.Form.Get("model")

//     _, err = db.Exec("UPDATE cars SET make = ?, model = ? WHERE id = ?", make, model, id)
//     if err != nil {
//         log.Fatal(err)
//     }

//     http.Redirect(w, r, "/", http.StatusSeeOther)
// }



func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM cars WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML file
	http.ServeFile(w, r, "template/template.html")
}
