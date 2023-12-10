create database data;

CREATE TABLE Vehicle (
    VehicleID INT PRIMARY KEY AUTO_INCREMENT(1,1),
    Make VARCHAR(255),
    Model VARCHAR(255),
    LicensePlate VARCHAR(50),
    EntryTime DATETIME,
    ExitTime DATETIME,
    IsParked BIT
);


CREATE TABLE Parking (
    ParkingID INT PRIMARY KEY AUTO_INCREMENT(1,1),
    Name VARCHAR(100),
    Location VARCHAR(255),
    Capacity INT
);


CREATE TABLE Maintenance (
    MaintenanceID INT PRIMARY KEY AUTO_INCREMENT(1,1),
    VehicleID INT,
    MaintenanceType VARCHAR(100),
    StartTime DATETIME,
    EndTime DATETIME,
    IsCompleted BIT,
    FOREIGN KEY (VehicleID) REFERENCES Vehicle(VehicleID)
);
