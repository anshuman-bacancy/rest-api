package server

import (
	"database/sql"
	"log"
	"net/http"
)

// Db is Db object
var Db *sql.DB

// InitializeDatabase makes connection to database
func InitializeDatabase(host, password, dbname string) {
	connection := "postgres://postgres:"+password+"@"+host+"/"+dbname+"?sslmode=disable"
	Db, _ = sql.Open("postgres", connection)
	log.Println("Database connection successful..")
}

// InitializeFileServer serves static files 
func InitializeFileServer() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resources/", http.StripPrefix("/resources/assets", fs))
}