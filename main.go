package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hunter32292/go-server-example/pkg/cmd"
	"github.com/spf13/cobra"
)

func init() {
	// Set new rand source to time now
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	command := cmd.NewServerCommand()
	log.Println("Starting...")
	code := Run(command)
	log.Println("Ending...")
	os.Exit(code)
}

func Run(command *cobra.Command) int {
	err := command.Execute()
	if err != nil {
		return 1
	}
	return 0
}
