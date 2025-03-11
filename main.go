package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  run <command>   - Run a command in background")
		fmt.Println("  logs            - Show logs")
		fmt.Println("  status          - Check process status")
		return
	}

	switch os.Args[1] {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: run <command> [args...]")
			return
		}
		RunCommand(os.Args[2], os.Args[3:]...)
	case "logs":
		CheckLogs()
	case "status":
		CheckStatus()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
