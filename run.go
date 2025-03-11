package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const logDir = "logs"

func ensureLogDir() {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
	}
}

func RunCommand(command string, args ...string) {
	ensureLogDir()
	logFile := filepath.Join(logDir, "background.log")

	cmd := exec.Command(command, args...)
	cmd.Stdout, _ = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cmd.Stderr = cmd.Stdout

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	fmt.Printf("Process started with PID: %d\n", cmd.Process.Pid)
}
