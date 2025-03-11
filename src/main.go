package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage:")
        fmt.Println("  run <command>   - Run a command in background")
        fmt.Println("  logs [runID]    - Show logs for a specific run ID or the latest run if no ID is provided")
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
        if len(os.Args) > 2 {
            runID, err := strconv.Atoi(os.Args[2])
            if err != nil {
                fmt.Println("Invalid run ID:", os.Args[2])
                return
            }
            CheckLogs(runID)
        } else {
            CheckLogs()
        }
    case "status":
        CheckStatus()
    default:
        RunCommand(os.Args[1], os.Args[2:]...)
    }
}