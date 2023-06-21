# GoGo-Drive

Concurrent Project Lab.
This API provides endpoints to upload and download files.

## Planning

[Link to deliveries planning](https://docs.google.com/document/d/1yCEdzCiwON9m7ZTTcJZ1W4d2oVwjuQ9ezzgKpHltSDo/edit?usp=sharing).

## Dependencies

- [Docker](https://docs.docker.com/desktop/install/linux-install/)
- [Docker Compose](https://docs.docker.com/compose/install/linux/)

## Run

1. Clone the repository:
   ```sh
   git clone git@github.com:AlefAdonis/GoGo-Drive.git
   ```
2. Go to project directory:

   ```sh
   cd GoGo-Drive
   ```

3. Build & Run the server:
   ```sh
   docker compose up
   ```

## Endpoints

### Upload File

Endpoint: **POST /upload**

Uploads a file to the server.

If the file is not in the same directory, insert the full path to upload!

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

### Delete File

Endpoint: **DELETE /delete/:filename**

Deletes a file with the specified filename.

Example:

```sh
curl -X DELETE localhost:8081/delete/doc.txt
```
