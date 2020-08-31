package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// NotFound - 404 Handler
func NotFound(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("pages/404.html")
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Server Errror: %s", err)))
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Server Errror: %s", err)))
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(bytes)
}
