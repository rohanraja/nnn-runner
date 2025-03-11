package main

import (
	"fmt"
	"os"
)

func CheckLogs() {
    logFile := getLogsPath()
	content, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}
	fmt.Println(string(content))
}
