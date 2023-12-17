Simple REST API with Go
This project implements a straightforward RESTful API in Go for managing student records within a database.

Overview
The API is structured into several key components:

datastore: Defines the data access layer responsible for database interactions.
handler: Manages incoming HTTP requests, interacts with the datastore, and processes data.
model: Specifies the structure of a student and interfaces for CRUD operations.
main.go: Acts as the entry point to the application, initializing the server and setting up endpoints.
Setup
Prerequisites
Go installed on your machine (Installation Guide)
MySQL database instance
