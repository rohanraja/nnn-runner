param (
    [string]$DownloadPath = "."
)

# Define the repository owner and name
$repoOwner = "rohanraja"
$repoName = "nnn-runner"

# Define the GitHub API URL for the latest release
$apiUrl = "https://api.github.com/repos/$repoOwner/$repoName/releases/latest"

# Fetch the latest release information
$response = Invoke-RestMethod -Uri $apiUrl -Headers @{ "User-Agent" = "PowerShell" }

# Extract the tag name and assets
$tagName = $response.tag_name
$assets = $response.assets

# Define the binary names for Windows
$nnnBinary = "nnn_amd64_win.exe"
$runWrapBinary = "run_wrap_amd64_win.exe"

# Find the download URLs for the binaries
$nnnUrl = $assets | Where-Object { $_.name -eq $nnnBinary } | Select-Object -ExpandProperty browser_download_url
$runWrapUrl = $assets | Where-Object { $_.name -eq $runWrapBinary } | Select-Object -ExpandProperty browser_download_url

# Ensure the download path exists
if (-Not (Test-Path -Path $DownloadPath)) {
    New-Item -ItemType Directory -Path $DownloadPath -Force
}

# Download the binaries using wget
Write-Output "Downloading $nnnBinary to $DownloadPath..."
Invoke-Expression "wget $nnnUrl -OutFile $DownloadPath\nn.exe"

Write-Output "Downloading $runWrapBinary to $DownloadPath..."
Invoke-Expression "wget $runWrapUrl -OutFile $DownloadPath\run_wrap.exe"

Write-Output "Installation complete."