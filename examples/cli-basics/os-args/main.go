package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(2)
	}
	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "Reminder CLI - CLI Basics"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("Hello and Welcome: %s\n", msg)
	case "help":
		fmt.Println("Help message")
	default:
		fmt.Printf("unknown command: %s\n", cmd)

	}
	// fmt.Println(os.Args)
}
