package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: run_wrapper <runPath> <command> <args...>")
        return
    }

    runPath := os.Args[1]
    command := os.Args[2]
    args := os.Args[3:]

    stdoutFile := filepath.Join(runPath, "stdout.txt")
    stderrFile := filepath.Join(runPath, "stderr.txt")
    infoFile := filepath.Join(runPath, "info.yml")

    cmd := exec.Command(command, args...)
    cmd.Stdout, _ = os.OpenFile(stdoutFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    cmd.Stderr, _ = os.OpenFile(stderrFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    err := cmd.Start()
    if err != nil {
        fmt.Println("Error starting command:", err)
        return
    }

    cmd.Wait()
    endTime := time.Now().Format(time.RFC3339)
    status := "Completed"
    if cmd.ProcessState.ExitCode() != 0 {
        status = "Error"
    }
    cmdStr := fmt.Sprintf("%s %s", command, strings.Join(args, " "))
    infoContent := fmt.Sprintf("CMD: %s\nStartTime: %s\nEndTime: %s\nStatus: %s\n",
        cmdStr, time.Now().Format(time.RFC3339), endTime, status)
    os.WriteFile(infoFile, []byte(infoContent), 0644)
}