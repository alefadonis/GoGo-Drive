package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var BaseDir = ""

func createDir() {
	homeDir, _ := os.UserHomeDir()
	BaseDir = filepath.Join(homeDir, "/gogo-drive")
	err := os.MkdirAll(BaseDir, 0755)
	if err != nil {
		log.Fatal("Failed to create the upload directory")
		return
	}
	log.Println("Base directory created at", BaseDir)
}

func main() {
	createDir()

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/files", ListFiles)
	// http.HandleFunc("/upload", UploadFile)

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		go func(w http.ResponseWriter, r *http.Request) {
			UploadFile(w, r)
		}(w, r)
	})

	http.HandleFunc("/download/", DownloadFile)
	http.HandleFunc("/delete/", DeleteFile)

	log.Println("Server running at port: 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
