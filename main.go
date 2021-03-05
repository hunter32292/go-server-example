package main

import (
	"runtime"

	"github.com/hunter32292/go-server-example/cmd"
)

// Setup system GOMAXPROCS based on the CPUs found on the system
func init() {
	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}
}

func main() {
	cmd.StartServer(cmd.GetServer())
}
