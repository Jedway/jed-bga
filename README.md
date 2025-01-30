# Jed-BGA (Basic Go API)

▶️ This is a simple public API written in Go that returns some information in JSON format.

▶️ It uses the following Go packages/libraries: [encoding/json](https://pkg.go.dev/encoding/json "encoding/json"), [fmt](https://pkg.go.dev/fmt "Go to fmt"), [net/http](https://pkg.go.dev/net/http "Go to net/http"), [time](https://pkg.go.dev/time "Go to time")

### Features

- Returns my registered email address.
- Returns the current datetime in ISO 8601 format.
- Returns the github URL fo this project's codebase.

## Setup Instructions

### Prerequisites

- Go 1.20 or higher

### Running Locally

1. Clone this repo:

   ```
   git clone https://github.com/Jedway/jed-bga.git

   ```
2. Navigate to the project directory:

   ```
   cd jed-bga
   ```
3. Run the server:

   ```
   go run main.go
   ```
4. Access the API at port 8080 by either using:
   Curl:

   ```
   curl http://localhost:8080
   ```

   Or opening ``http://localhost:8080`` on your browser

### Deployment

The API is deployed and hosted on [Railway](https://railway.com/ "Railway.com"). You can access it at:

[https://jed-bga-production.up.railway.app/](https://jed-bga-production.up.railway.app/ "Go to live deployment")

### Endpoint

```
GET /
```

### Response Format

```json
   {
     "email": "jedotoekong@gmail.com",
     "current_datetime": "2025-01-30T09:30:00Z",
     "github_url": "https://github.com/Jedway/jed-bga.git"
   }
```

### Example Usage


```bash
   curl https://jed-bga-production.up.railway.app/
```

### Backlink

This project uses Go. Learn more about talented Go developers at [HNG Tech](https://hng.tech/hire/golang-developers "Hire Go devs")
