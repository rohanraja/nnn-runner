package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const runsDir = "runs"

func ensureDir(path string) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        os.Mkdir(path, 0755)
    }
}

func getNextRunID() int {
    files, err := os.ReadDir(runsDir)
    if err != nil {
        fmt.Println("Error reading runs directory:", err)
        return 1
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
    return maxID + 1
}

func createRunFiles(runID int, command string, args []string) (string, error) {
    runPath := filepath.Join(runsDir, strconv.Itoa(runID))
    ensureDir(runPath)

    stdoutFile := filepath.Join(runPath, "stdout.txt")
    stderrFile := filepath.Join(runPath, "stderr.txt")
    infoFile := filepath.Join(runPath, "info.yml")

    files := []string{stdoutFile, stderrFile, infoFile}
    for _, file := range files {
        f, err := os.Create(file)
        if err != nil {
            return "", err
        }
        f.Close()
    }

    cmdStr := fmt.Sprintf("%s %s", command, strings.Join(args, " "))
    infoContent := fmt.Sprintf("CMD: %s\nStartTime: %s\nEndTime: \nStatus: Running\n", cmdStr, time.Now().Format(time.RFC3339))
    os.WriteFile(infoFile, []byte(infoContent), 0644)

    return runPath, nil
}

func RunCommand(command string, args ...string) {
    ensureDir(runsDir)
    runID := getNextRunID()
    runPath, err := createRunFiles(runID, command, args)
    if err != nil {
        fmt.Println("Error creating run files:", err)
        return
    }

    wrapperPath := "run_wrap" // Adjust the path if necessary
    cmd := exec.Command(wrapperPath, append([]string{runPath, command}, args...)...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err = cmd.Start()
    if err != nil {
        fmt.Println("Error starting run_wrapper:", err)
        return
    }

    fmt.Printf("Process started with PID: %d\n", cmd.Process.Pid)
    fmt.Printf("Run ID: %d\n", runID)
}