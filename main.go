// package main

// import (
//     "context"
//     "fmt"
//     "html/template"
//     "net/http"
//     "time"

//     "go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
// )

// type Car struct {
//     ID        int
//     Make      string
//     Model     string
//     Status    string
//     CreatedAt time.Time
// }

// type Welcome struct {
//     Sale string
//     Time string
//     Cars []Car // Store cars in a slice
// }

// var cars []Car // In-memory storage for cars

// func main() {
//     // MongoDB connection
//     client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
//     if err != nil {
//         panic(err.Error())
//     }
//     defer func() {
//         if err = client.Disconnect(context.Background()); err != nil {
//             panic(err.Error())
//         }
//     }()

//     // Ping the MongoDB server
//     err = client.Ping(context.Background(), nil)
//     if err != nil {
//         panic(err.Error())
//     }

//     fmt.Println("Connected to MongoDB!")

//     // Go HTTP server setup
//     welcome := Welcome{"Sale Begins Now", time.Now().Format(time.Stamp), cars}
//     template := template.Must(template.ParseFiles("template/template.html"))
	

//     http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         if sale := r.FormValue("sale"); sale != "" {
//             welcome.Sale = sale
//         }
//         welcome.Cars = cars // Update cars data in welcome struct
//         if err := template.ExecuteTemplate(w, "template.html", welcome); err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//     })

//     http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
//         if r.Method == http.MethodPost {
//             r.ParseForm()
//             make := r.FormValue("make")
//             model := r.FormValue("model")

//             newID := len(cars) + 1

//             newCar := Car{newID, make, model, "Available", time.Now()}
//             cars = append(cars, newCar)

//             // Insert the newly added car to MongoDB
//             collection := client.Database("carDB").Collection("cars")
//             _, err := collection.InsertOne(context.Background(), newCar)
//             if err != nil {
//                 panic(err.Error())
//             }

//             http.Redirect(w, r, "/", http.StatusFound)
//             return
//         }
//         http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//     })

//     fmt.Println(http.ListenAndServe(":8000", nil))
// }
package main

import (
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

func updateCarHandler(w http.ResponseWriter, r *http.Request) {
	// Handle update logic here
	// ...
}

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
