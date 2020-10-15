package controller

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hunter32292/go-server-example/pkg/dao"
	"github.com/hunter32292/go-server-example/pkg/models"
)

// UserData - The collection of Users retained in memory as a slice of structs
var UserData []*models.User

// SetupUserHandler - setup all the controller paths for Users
func SetupUserHandler(handler *http.ServeMux) {
	handler.HandleFunc("/user", Show)
	handler.HandleFunc("/user/create", Create)
	handler.HandleFunc("/user/update", Update)
	handler.HandleFunc("/user/replace", Replace)
	handler.HandleFunc("/user/delete", Delete)

	LoadData()
}

//LoadData - Setup Data For Users
func LoadData() {
	data, err := dao.FileLoadInData("data/MOCK_DATA.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for index, item := range records {
		if index == 0 {
			continue
		}
		UserData = append(UserData, &models.User{Id: index, First_name: item[1], Last_name: item[2], Email: item[3]})
	}
}

// Show - a User
func Show(w http.ResponseWriter, r *http.Request) {
	log.Println("Show User Data")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("["))
	for index, data := range UserData {
		payload, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(payload)
		if index != len(UserData)-1 {
			w.Write([]byte(","))
		}
	}
	w.Write([]byte("]"))
}

// Create - a New User
func Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create User Data")
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
	log.Println("Update User Data")
	io.WriteString(w, "Update a User")
}

// Replace - a current User
func Replace(w http.ResponseWriter, r *http.Request) {
	log.Println("Replace User Data")
	io.WriteString(w, "Replace a current User")
}

// Delete - a User
func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete User Data")
	io.WriteString(w, "Delete a User")
}
