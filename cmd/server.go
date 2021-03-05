package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/hunter32292/go-server-example/pkg/controller"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Create Instance Vars
var (
	Name = "example-server"
	// Wait group for main go routines
	waitgroup = sync.WaitGroup{}
	host      = "localhost"
	port      = "8080"
	// certPath  = "./cert.pem"
	// keyPath   = "./key.pem"
	certPath = ""
	keyPath  = ""

	Server *http.Server
)

func init() {
	Server = &http.Server{}
}

// StartServer - Start Server Server
func StartServer(server *http.Server) {

	if len(os.Getenv("LOG_FILE")) > 0 {
		file, err := os.OpenFile(os.Getenv("LOG_FILE")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		log.SetOutput(file)
	}
	log.Println("Starting ...")

	mux := http.NewServeMux() // Create Main Handler
	setupControllers(mux)     // Setup all controllers for server

	server.Addr = host + ":" + port
	server.Handler = mux
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second

	waitgroup.Add(1)
	if certPath != "" && keyPath != "" {
		go runTLS(server)
	} else {
		go run(server)
	}

	// Graceful Exit Scenario
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			log.Printf("Got %s signal. Aborting...\n", sig)
			server.Close() // Close currently running instance of server
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
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", controller.HomeHandler)
	mux.HandleFunc("/404", controller.NotFound)
	// Add more like so:
	// controller.SetupNAMEHandler(mux)
}

// StopServer - Shut down the Server
func StopServer(server *http.Server) error {
	return server.Close()
}

// GetServer - Return application server object
func GetServer() *http.Server {
	return Server
}
