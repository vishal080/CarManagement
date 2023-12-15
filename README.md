# Car Garage Management Web App
This is a simple web application for managing cars in a garage. It allows users to add, update, and delete car information
## Features

- **Add Car**: Users can add cars to the garage by providing details like make, model, etc.
- **Update Car**: Edit existing car information.
- **Delete Car**: Remove cars from the garage.
- **Display Cars**: View the list of cars stored in the system.

## Tech Stack

- **Frontend**: HTML, CSS
- **Backend**: Golang
- Gofr
- **Database**: SQLite

## Setup Instructions

1. **Clone Repository:**
    ```bash
    git clone https://github.com/vishal080/CarManagement.git
    
   

2. **Database Setup:**
    - Ensure you have SQLite installed.
    - Run the application to auto-create the database (`cars.db`).

3. **Run the Application:**
    ```bash
    go run main.go
    ```
    Access the application at [http://localhost:8080]
## Usage

1. Access the web application in your browser.
2. Use the interface to:
   - Add cars by providing make, model, etc.
   - Update car information.
   - Delete cars from the garage.
   - View the list of cars stored in the system.

## Contributing

Contributions are welcome! Fork the repository, create a branch, make your changes, and submit a pull request.


## Flow Chart 

Start
|
|__ User opens Car Garage Management System
    |
    |__ User navigates to Cars in Garage
    |   |
    |   |__ System retrieves car data from the database
    |   |   |
    |   |   |__ Display cars list on the webpage
    |   |
    |   |__ User clicks Add Car button
    |       |
    |       |__ System prompts for car details
    |       |   |
    |       |   |__ User enters car details
    |       |   |
    |       |   |__ System adds the car to the database
    |       |
    |       |__ Display updated cars list
    |
    |__ User navigates to Contact Us
        |
        |__ Display contact information
        |
        |__ User can contact via provided details
|
End


## Author

- Vishal Chaudhary
   https://github.com/vishal080



