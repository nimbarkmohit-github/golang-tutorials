package main

import (
	"flag"
	"fmt"
	"os"

	"../../client"
)

var (
	backendURIFlag = flag.String("backend", "http://localhost8080", "Backend API URL")
	helpFlag       = flag.Bool("help", false, "Display a helpful message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %v\n", err)
		os.Exit(2)
	}
}
