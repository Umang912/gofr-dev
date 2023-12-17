package main1

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	password := "*****" 

	// Database connection parameters
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(localhost:3306)/data", password))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	fmt.Println("Connected to the database!")

	// Perform a sample query on Vehicle table
	rows, err := db.Query("SELECT * FROM Vehicle")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process the query result
	for rows.Next() {
		var vehicleID int
		var make, model, licensePlate string
		var entryTime, exitTime string
		var isParked bool
		if err := rows.Scan(&vehicleID, &make, &model, &licensePlate, &entryTime, &exitTime, &isParked); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("VehicleID: %d, Make: %s, Model: %s, LicensePlate: %s, EntryTime: %s, ExitTime: %s, IsParked: %t\n", vehicleID, make, model, licensePlate, entryTime, exitTime, isParked)
	}
}
