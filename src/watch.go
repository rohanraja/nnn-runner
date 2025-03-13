package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func WatchLogs(runID ...int) {
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
    stdoutFile := filepath.Join(runPath, "stdout.txt")
    stderrFile := filepath.Join(runPath, "stderr.txt")
    infoFile := filepath.Join(runPath, "info.yml")

    go printFileContents(stdoutFile)
    go printFileContents(stderrFile)

    for {
        status := checkStatus(infoFile)
        if status != "Running" {
            fmt.Printf("Status changed to %s. Stopping watch.\n", status)
            break
        }
        time.Sleep(1 * time.Second)
    }
}

func printFileContents(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file %s: %v\n", filePath, err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            time.Sleep(1 * time.Second)
            continue
        }
        fmt.Print(line)
    }
}

func checkStatus(infoFile string) string {
    file, err := os.Open(infoFile)
    if err != nil {
        fmt.Printf("Error opening file %s: %v\n", infoFile, err)
        return "Error"
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "Status:") {
            return strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file %s: %v\n", infoFile, err)
        return "Error"
    }

    return "Unknown"
}