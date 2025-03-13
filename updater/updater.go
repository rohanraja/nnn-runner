package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
    repoOwner = "rohanraja"
    repoName  = "nnn-runner"
    apiURL    = "https://api.github.com/repos/%s/%s/releases/latest"
)

type Release struct {
    TagName string `json:"tag_name"`
    Assets  []Asset `json:"assets"`
}

type Asset struct {
    Name               string `json:"name"`
    BrowserDownloadURL string `json:"browser_download_url"`
}

func main() {
    url := fmt.Sprintf(apiURL, repoOwner, repoName)
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error fetching latest release: %v\n", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error fetching latest release: %s\n", resp.Status)
        return
    }

    var release Release
    if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
        fmt.Printf("Error decoding release JSON: %v\n", err)
        return
    }

    fmt.Printf("Latest release: %s\n", release.TagName)
    for _, asset := range release.Assets {
        if isValidAsset(asset.Name) {
            fmt.Printf("Downloading %s...\n", asset.Name)
            if err := downloadFile(asset.BrowserDownloadURL, asset.Name); err != nil {
                fmt.Printf("Error downloading %s: %v\n", asset.Name, err)
                return
            }
            fmt.Printf("Downloaded %s\n", asset.Name)
        }
    }
}

func isValidAsset(assetName string) bool {
    os := runtime.GOOS
    arch := runtime.GOARCH

    if strings.Contains(assetName, os) && strings.Contains(assetName, arch) {
        return true
    }
    return false
}

func downloadFile(url, filename string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("error downloading file: %s", resp.Status)
    }

    out, err := os.Create(filepath.Join(".", filename))
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}