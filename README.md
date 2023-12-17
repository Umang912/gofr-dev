
# Simple REST API with Go

This project implements a straightforward RESTful API in Go for managing student records within a database.

# Overview

The API is structured into several key components:

datastore: Defines the data access layer responsible for database interactions.
handler: Manages incoming HTTP requests, interacts with the datastore, and processes data.
model: Specifies the structure of a student and interfaces for CRUD operations.
main.go: Acts as the entry point to the application, initializing the server and setting up endpoints.

# Setup
Prerequisites

Go installed on your machine (Installation Guide)
MySQL database instance






## Installation

1.Clone the repository:


```
git clone https://github.com/Umang912/gofr-dev.git

```
2.Navigate to the project directory:

```
cd gofr-dev

```
 
3.Create a .env file in the root directory and set up the required environment variables:

```
# Database Configuration
DB_HOST=localhost
DB_PORT=2001
DB_USER=root
DB_PASSWORD=password
DB_NAME=Students
DB_DIALECT=mysql

```

4.Run the application:

```
go run main.go

```

## API Endpoints

The API provides the following endpoints for managing student records:

GET /students/{id}: Retrieve a student by ID.

POST /students: Create a new student.

PUT /students/{id}: Update a student by ID.

DELETE /students/{id}: Delete a student by ID.