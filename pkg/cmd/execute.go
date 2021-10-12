package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/hunter32292/go-server-example/pkg/options"
	"github.com/hunter32292/go-server-example/pkg/server"
	trace "github.com/hunter32292/go-server-example/pkg/tracer"
)

func NewServerCommand() *cobra.Command {
	cfg := options.DefaultFlags()
	cleanFlagSet := pflag.NewFlagSet(cfg.Name, pflag.ContinueOnError)
	cleanFlagSet.SetNormalizeFunc(options.WordSepNormalizeFunc)

	cmd := &cobra.Command{
		Use:                cfg.Name,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			// Parse FlagSet if flags don't work, return failure and help message
			err := cleanFlagSet.Parse(args)
			if err != nil {
				cmd.Usage()
				os.Exit(1)
			}

			// Check if there are extra commands in the args
			cmds := cleanFlagSet.Args()
			if len(cmds) > 0 {
				cmd.Usage()
				os.Exit(1)
			}

			help, err := cleanFlagSet.GetBool("help")
			if err != nil {
				log.Fatal(err, "Failed to parse help message")
			}
			if help {
				cmd.Help()
				return
			}
			ctx := context.Background()
			start(ctx, cfg)

		},
	}

	cleanFlagSet.BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))

	return cmd
}

func setupLogger() {
	if len(os.Getenv("LOG_FILE")) > 0 {
		file, err := os.OpenFile(os.Getenv("LOG_FILE")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		log.SetOutput(file)
	}
	log.Println("Logger Created")
}

func setupSignalHandler(c chan os.Signal) {
	<-c
	log.Printf("Got terminate signal. Aborting...\n")
	server.Close() // Close currently running instance of server
	os.Exit(0)
}

func start(ctx context.Context, cfg *options.ServerFlags) {
	var waitgroup sync.WaitGroup

	trace.NewTraceConfig(cfg.Name)
	trace.CreateGlobalTracer()
	log.Println("Global tracer created...")

	setupLogger()
	waitgroup.Add(2) // Signal handler and server

	go func() {
		defer waitgroup.Done()
		server.RunServer(ctx, cfg)
	}()

	// Graceful Exit Scenario
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go setupSignalHandler(c)

	waitgroup.Wait()
}
