package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func DownloadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fileName := ps.ByName("filename")
	log.Println("[GET] /download/" + fileName)

	filePath := filepath.Join(UploadDir, fileName)

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Failed to open the file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Failed to get file information", http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Copy the file data to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to write the file to the response", http.StatusInternalServerError)
		return
	}
}
