package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func CheckLogs(runID ...int) {
    var id int
    if len(runID) > 0 {
        id = runID[0]
    } else {
        id = getLastRunID()
        if id == 0 {
            fmt.Println("No runs found.")
            return
        }
    }

    runPath := filepath.Join(runsDir, strconv.Itoa(id))
    files := []string{"stdout.txt", "stderr.txt", "info.yml"}

    for _, file := range files {
        filePath := filepath.Join(runPath, file)
        content, err := os.ReadFile(filePath)
        if err != nil {
            fmt.Printf("Error reading %s: %v\n", file, err)
            continue
        }
        fmt.Printf("Contents of %s:\n%s\n", file, string(content))
    }
}