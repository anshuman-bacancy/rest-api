package handler

import (
	"fmt"
	"html/template"
	"net/http"
	userservice "services"

	"github.com/gorilla/mux"
)

var tpl *template.Template

// Home handles "/"
func Home(res http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseGlob("assets/*.html"))
	tpl.ExecuteTemplate(res, "index.html", nil) 
}

// GetUserHandler handles "/users/{email}"
func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	email := mux.Vars(req)["email"]
	user := userservice.GetUser(email)
	res.Write([]byte(fmt.Sprintf("%+v\n", user)))
}

// GetUsersHandler handles "/users"
func GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	allUsers := userservice.GetUsers()
	res.Write([]byte(fmt.Sprintf("%+v\n", allUsers)))
}

// AddUserHandler handles "/user/"
func AddUserHandler(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	email := req.FormValue("email")
	position := req.FormValue("position")

	userservice.AddUser(name, email, position)
}

// DeleteUserHandler handles "/user/{email}"
func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
	email := mux.Vars(req)["email"]
	userservice.DeleteUser(email)
}

// UpdateUserHandler handles "/user/{email}"
func UpdateUserHandler(res http.ResponseWriter, req *http.Request) {
	oldEmail := mux.Vars(req)["email"]
	name := req.FormValue("name")
	newEmail := req.FormValue("email")
	position := req.FormValue("position")

	userservice.UpdateUser(oldEmail, name, newEmail, position)
}