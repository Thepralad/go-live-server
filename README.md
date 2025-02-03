# Go Live Server

A simple command-line live server built with Go that serves static files and automatically reloads the browser when files change.

## Features
- Serves static files over HTTP
- Auto-reloads the browser on file changes using WebSockets
- Supports custom ports and file paths

## Installation

### Clone the Repository
```bash
git clone https://github.com/your-username/go-live-server.git
cd go-live-server
```

### Install Dependencies
```bash
go mod tidy
```

### Run Locally
```bash
go run main.go -port=8080 -path=.
```

### Install Globally
```bash
go install
```
Then run it from anywhere:
```bash
go-live-server -port=8080 -path=/path/to/your/project
```

## Example Usage
- **Serve Current Directory on Default Port (8080):**
  ```bash
  go run main.go
  ```
- **Serve a Specific Directory on Port 3000:**
  ```bash
  go run main.go -port=3000 -path=/path/to/project
  ```

## License
MIT


