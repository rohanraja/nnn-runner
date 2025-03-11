package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckLogs() {
	logFile := filepath.Join("logs", "background.log")
	content, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}
	fmt.Println(string(content))
}
