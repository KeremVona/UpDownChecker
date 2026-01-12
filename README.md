# Website Uptime Monitor

A simple CLI tool written in Go to monitor the uptime of websites and send email alerts if they go down.

## Features
- Monitors a list of websites defined in a CSV file.
- configurable check intervals per website.
- Sends email alerts via SMTP when a website is unreachable (non-200 status or network error).
- Cross-platform (Windows/Linux/macOS).

## Prerequisites
- [Go](https://go.dev/dl/) (1.21 or later recommended)

## Configuration

### 1. Targets (`sites.csv`)
Make a file named `sites.csv` with `url` and `interval` columns. The interval is in seconds.
```csv
url,interval
https://google.com,30
https://github.com,60
```
*Note: The tool handles CSVs made by Excel (with BOM).*

### 2. Email Settings (`config.json`)
Make a file named `config.json` with your SMTP details.
```json
{
  "smtp_host": "smtp.gmail.com",
  "smtp_port": 587,
  "smtp_user": "your-email@gmail.com",
  "smtp_pass": "your-app-password",
  "from_email": "your-email@gmail.com",
  "to_email": "admin@example.com"
}
```

## Build and Run

### Windows
```powershell
go build -o monitor.exe cmd/monitor/main.go
.\monitor.exe
```

### Linux / macOS
```bash
go build -o monitor cmd/monitor/main.go
./monitor
```

### Cross-compile for Windows from Linux
```bash
GOOS=windows GOARCH=amd64 go build -o monitor.exe cmd/monitor/main.go
```

## Usage
By default, the tool looks for `sites.csv` and `config.json` in the current directory. You can specify different paths:

```bash
./monitor --sites my-sites.csv --config my-config.json
```
