package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Vehicle struct {
	VehicleID    int       `json:"VehicleID"`
	Make         string    `json:"Make"`
	Model        string    `json:"Model"`
	LicensePlate string    `json:"LicensePlate"`
	EntryTime    time.Time `json:"EntryTime"`
	ExitTime     time.Time `json:"ExitTime"`
	IsParked     bool      `json:"IsParked"`
}

type Parking struct {
	ParkingID int    `json:"ParkingID"`
	Name      string `json:"Name"`
	Location  string `json:"Location"`
	Capacity  int    `json:"Capacity"`
}

type Maintenance struct {
	MaintenanceID   int       `json:"MaintenanceID"`
	VehicleID       int       `json:"VehicleID"`
	MaintenanceType string    `json:"MaintenanceType"`
	StartTime       time.Time `json:"StartTime"`
	EndTime         time.Time `json:"EndTime"`
	IsCompleted     bool      `json:"IsCompleted"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(your-host:your-port)/your-database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
