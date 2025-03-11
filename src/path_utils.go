package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const logDir = "logs"

const runsDir = "runs"

func getLogsPath() string {
    return filepath.Join(logDir, "run.log")
}


func getLastRunID() int {
    files, err := os.ReadDir(runsDir)
    if err != nil {
        fmt.Println("Error reading runs directory:", err)
        return 0
    }

    maxID := 0
    for _, file := range files {
        if file.IsDir() {
            id, err := strconv.Atoi(file.Name())
            if err == nil && id > maxID {
                maxID = id
            }
        }
    }
    return maxID
}

func getNextRunID() int {
    return getLastRunID() + 1
}