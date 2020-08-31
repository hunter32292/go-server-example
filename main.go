package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/hunter32292/warmups/pkg/controller"
)

// Create Instance Vars
var (
	// Wait group for main go routines
	waitgroup = sync.WaitGroup{}
	host      = "localhost"
	port      = "8080"
	// certPath  = "./cert.pem"
	// keyPath   = "./key.pem"
	certPath = ""
	keyPath  = ""
)

func main() {
	log.Println("Starting ...")

	mux := http.NewServeMux() // Create Main Handler
	setupControllers(mux)     // Setup all controllers for server

	s := &http.Server{
		Addr:         host + ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	waitgroup.Add(1)
	if certPath != "" && keyPath != "" {
		go runTLS(s)
	} else {
		go run(s)
	}

	// Graceful Exit Scenario
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			log.Printf("Got %s signal. Aborting...\n", sig)
			s.Close() // Close currently running instance of server
		}
	}()

	waitgroup.Wait()
	os.Exit(0)
}

func runTLS(s *http.Server) {
	log.Fatal(s.ListenAndServeTLS(certPath, keyPath))
	waitgroup.Done()
}

func run(s *http.Server) {
	log.Fatal(s.ListenAndServe())
	waitgroup.Done()
}

func setupControllers(mux *http.ServeMux) {
	controller.SetupUserHandler(mux)
	mux.HandleFunc("/404", controller.NotFound)
	// Add more like so:
	// controller.SetupNAMEHandler(mux)
}
