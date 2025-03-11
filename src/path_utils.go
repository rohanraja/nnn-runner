package main

import (
    "path/filepath"
)

const logDir = "logs"

func getLogsPath() string {
    return filepath.Join(logDir, "run.log")
}