package main

import (
	"handler"
	"log"
	"net/http"
	"server"

	"github.com/gorilla/mux"
)

func main() {
	servername := "userserver"
	password := "password"
	host := "localhost"

	server.InitializeDatabase(host, password, servername)
	defer server.Db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", handler.Home).Methods("GET")
	router.HandleFunc("/users", handler.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{email}", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", handler.AddUserHandler).Methods("POST")
	router.HandleFunc("/user/{email}", handler.UpdateUserHandler).Methods("UPDATE")
	router.HandleFunc("/user/{email}", handler.DeleteUserHandler).Methods("DELETE")

	server.InitializeFileServer()

	log.Println("Server running at 8000")
	http.ListenAndServe(":8000", router)
}