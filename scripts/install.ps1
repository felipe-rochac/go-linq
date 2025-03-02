# Script to update Go on Windows

# 1. Get the latest Go version from the download page
$DownloadURL = Invoke-WebRequest -Uri "https://go.dev/dl/?mode=json" | ConvertFrom-Json | Where-Object { $_.os -eq "windows" -and $_.arch -eq "amd64" } | Select-Object -ExpandProperty version

write-host "Downloading Go version: $DownloadURL"

if (-not $DownloadURL) {
    Write-Error "Could not find the latest Go download URL."
    exit 1
}

$InstallerFileName = Split-Path $DownloadURL -Leaf
$InstallerPath = Join-Path $env:TEMP $InstallerFileName

# 2. Download the installer
Write-Host "Downloading Go installer: $InstallerFileName"
Invoke-WebRequest -Uri $DownloadURL -OutFile $InstallerPath

if (-not (Test-Path $InstallerPath)) {
    Write-Error "Failed to download the Go installer."
    exit 1
}

# 3. Run the installer silently
Write-Host "Running Go installer..."
Start-Process -FilePath $InstallerPath -ArgumentList "/s" -Wait -NoNewWindow

# 4. Verify the installation
Write-Host "Verifying Go installation..."
$GoVersion = go version
if ($LASTEXITCODE -eq 0) {
    Write-Host "Go updated successfully. Version: $GoVersion"
} else {
    Write-Error "Failed to verify Go installation."
    exit 1
}

# 5. Clean up the installer (optional)
Remove-Item $InstallerPath