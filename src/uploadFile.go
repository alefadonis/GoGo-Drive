package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func UploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("[POST] /upload")

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = os.MkdirAll(UploadDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Failed to create the upload directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(UploadDir, handler.Filename)

	// Create a new file on the server to save the uploaded file
	destinationFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create the file on the server", http.StatusInternalServerError)
		return
	}

	defer destinationFile.Close()

	// Copy the uploaded file to the destination file on the server
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		http.Error(w, "Failed to save the uploaded file", http.StatusInternalServerError)
		return
	}

	// Return a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully!\n"))
}
