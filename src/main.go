package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var BaseDir = ""
var UploadChannel chan int

func PrintNumGoRoutines() {
	for {
		numGoRoutines := runtime.NumGoroutine()
		if numGoRoutines > 2 {
			log.Printf("Numero do goroutines %d/n", numGoRoutines)
		}
		time.Sleep(1 * time.Second)
	}
}
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
	go PrintNumGoRoutines()
	createDir()

	mux := http.NewServeMux()

	UploadChannel = make(chan int)

	go func() { mux.HandleFunc("/", HomePage) }()
	go func() { mux.HandleFunc("/files", ListFiles) }()

	go func() { mux.HandleFunc("/upload", UploadFile) }()

	go func() { mux.HandleFunc("/download/", DownloadFile) }()

	go func() { mux.HandleFunc("/delete/", DeleteFile) }()

	log.Println("Server running at port: 8081")
	err := http.ListenAndServe(":8081", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
