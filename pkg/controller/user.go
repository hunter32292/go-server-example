package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hunter32292/warmups/pkg/models"
)

// UserData - The collection of Users retained in memory as a slice of structs
var UserData []*models.User

// SetupUserHandler - setup all the controller paths for Users
func SetupUserHandler(handler *http.ServeMux) {
	handler.HandleFunc("/user/create", Create)
	handler.HandleFunc("/user/update", Update)
	handler.HandleFunc("/user/replace", Replace)
	handler.HandleFunc("/user/delete", Delete)
}

// Create - a New User
func Create(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	io.WriteString(w, "Create a New User")
	bytes, err := json.Marshal(newUser)
	if err != nil {
		fmt.Fprintf(w, "Failed to Marshal Data from Request %s", err)
	}
	log.Println("Receieved Data: ", bytes)
	UserData = append(UserData, newUser)
	fmt.Fprintln(w, newUser)
}

// Update - a User
func Update(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Update a User")
}

// Replace - a current User
func Replace(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Replace a current User")
}

// Delete - a User
func Delete(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Delete a User")
}
