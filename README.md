# GoGo-Drive

Concurrent Project Lab.
This API provides endpoints to upload and download files.

## Dependencies

- github.com/julienschmidt/httprouter - HTTP request router
- github.com/google/uuid - UUID generation

## Setup

1. Clone the repository:
   ```sh
   git clone git@github.com:AlefAdonis/GoGo-Drive.git
   ```
2. Install dependencies

   ```
   go get github.com/julienschmidt/httprouter github.com/google/uuid
   ```

3. Build and run the Go server:

   ```sh
   cd GoGo-Drive
   ```

   ```sh
   go run src/main.go
   ```

## Endpoints

### Upload File

Endpoint: **POST /upload**

Uploads a file to the server.

Example:

```sh
curl -X POST -F "file=@/path/to/file" localhost:8081/upload
```

### List Files

Endpoint: **GET /files**

Lists all uploaded files.

Example:

```sh
curl localhost:8081/files
```

### Download File

Endpoint: **GET /download/:filename**

Downloads a file with the specified filename.

Example:

```sh
curl -OJ localhost:8081/download/doc.txt
```
