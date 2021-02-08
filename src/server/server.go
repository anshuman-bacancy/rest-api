package server

import (
	"database/sql"
	"log"
	"net/http"
	"utilities"

	_ "github.com/lib/pq"
)

// Db is global Db object
var Db *sql.DB
var dberror error

// InitializeDatabase makes connection to database
func InitializeDatabase(host, password, dbname string) {
	connection := "postgres://postgres:"+password+"@"+host+"/"+dbname+"?sslmode=disable"
	Db, dberror = sql.Open("postgres", connection)
	utilities.CheckError(dberror)
	log.Println("Database connection successful..")
}

// InitializeFileServer serves static files 
func InitializeFileServer() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resources/", http.StripPrefix("/resources/assets", fs))
}