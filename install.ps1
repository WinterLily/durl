# Check for admin privileges
if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Warning "Please run as Administrator!"
    exit 1
}

# Hello message

Write-Host "🐍🐍 Thanks for using DURL! 🐍🐍"

# Configuration
$BinaryPath = ".\build\durl.exe"
$InstallDir = "$env:ProgramFiles\durl"
$BinaryInstallPath = "$InstallDir\durl.exe"

# Check if binary exists
if (-not (Test-Path $BinaryPath)) {
    Write-Error "Binary not found at $BinaryPath"
    exit 1
}

# Create installation directory
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

# Copy binary
Copy-Item $BinaryPath -Destination $BinaryInstallPath -Force

# Update PATH
$CurrentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($CurrentPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$CurrentPath;$InstallDir", "Machine")
}

Write-Host "durl has been installed successfully. Please restart your terminal to use durl from anywhere."
