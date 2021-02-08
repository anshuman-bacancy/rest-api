package userservice

import (
	"database/sql"
	"fmt"
	"models"
	"server"
	"utilities"

	"github.com/google/uuid"
)

// GetUser return []model.User from database
func GetUser(email string) models.User {
	getUser := "SELECT * from users where email = $1"
	var user models.User
	row := server.Db.QueryRow(getUser, email)
	rowErr := row.Scan(&user.Id, &user.Name, &user.Email, &user.Position)
	switch rowErr {
		case sql.ErrNoRows:
			fmt.Println("No rows found...")
		case nil:
			fmt.Println("Received: ", user)
		default:
			panic(rowErr)
	}

	return user
}

// GetUsers return []model.User
func GetUsers() []models.User {
	users := "SELECT * from users"
	rows, err := server.Db.Query(users)
	utilities.CheckError(err)

	defer rows.Close()

	var allUsers []models.User

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Position)
		utilities.CheckError(err)
		allUsers = append(allUsers, user)
	}
	return allUsers
}

// AddUser inserts model.User into Database
func AddUser(name, email, position string) {
	id := uuid.NewString()
	insUser := `INSERT INTO users (id, name, email, position) VALUES ($1, $2, $3, $4)`
  _, err := server.Db.Exec(insUser, id, name, email, position)
	utilities.CheckError(err) 
}

// DeleteUser deletes model.User from database
func DeleteUser(email string) {
	delUser := "DELETE from users where email = $1"
	_, err := server.Db.Exec(delUser, email)
	utilities.CheckError(err)
}

// UpdateUser updates []model.User based on email
func UpdateUser(oldEmail, name, newEmail, position string) {
	updUser := "UPDATE users set name=$1, email=$2, position=$3 where email=$4"
	_, err := server.Db.Exec(updUser,  name, newEmail, position, oldEmail)
	utilities.CheckError(err)
}