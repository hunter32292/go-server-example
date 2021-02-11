package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	trace "github.com/hunter32292/go-server-example/pkg/tracer"
)

// NotFound - 404 Handler
func NotFound(w http.ResponseWriter, r *http.Request) {
	t := trace.GetGlobalTracer()
	span := t.StartSpan("404 not found")
	defer span.Finish()

	log.Println("Serve Not Found")
	file, err := os.Open("pages/404.html")
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Server Errror: %s", err)))
	}
	span.LogEvent("Read in 404 page")
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Server Errror: %s", err)))
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(bytes)
	span.LogEvent("Wrote 404 page to browser")
}

// HomeHandler - Home page redirect view and 404 handle
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := trace.GetGlobalTracer()
	span := t.StartSpan("Serve Home Page")
	defer span.Finish()

	if r.URL.Path != "/" {
		NotFound(w, r)
		return
	}
	log.Println("Serve Home Page")
	file, err := os.Open("pages/index.html")

	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Server Errror: %s", err)))
	}
	span.LogEvent("Read Home page from directory")
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Server Errror: %s", err)))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	span.LogEvent("Wrote Home page to browser")
}
