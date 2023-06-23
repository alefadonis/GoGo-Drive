package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var mutexUpload sync.Mutex

func UploadFile(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.Println("[POST] /upload")

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	log.Printf("1 -  %s", handler.Filename)

	filePath := filepath.Join(BaseDir, handler.Filename)

	log.Printf("2 -  %s", handler.Filename)

	mutexUpload.Lock()
	_, err = os.Stat(filePath)
	if err == nil {
		http.Error(w, fmt.Sprintf("File %s alredy exists in the base directory.", filePath), http.StatusNotFound)
		return
	}

	destinationFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create the file on the server", http.StatusInternalServerError)
		return
	}

	mutexUpload.Unlock()

	log.Printf("3 -  %s", handler.Filename)

	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, file)
	if err != nil {
		http.Error(w, "Failed to save the uploaded file", http.StatusInternalServerError)
		return
	}

	log.Printf("4 -  %s", handler.Filename)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully!\n"))

	runtime := time.Since(startTime).Seconds()
	log.Printf("Runtime upload %s: %.2fs\n", handler.Filename, runtime)
}
