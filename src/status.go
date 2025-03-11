package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckStatus() {
	logFile := filepath.Join("logs", "background.log")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		fmt.Println("No logs found, process may not have started.")
		return
	}

	fmt.Println("Process is running, check logs for output.")
}
