package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var mutexDownload sync.Mutex

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	fileName := urlParts[len(urlParts)-1]

	log.Println("[GET] /download/" + fileName)

	filePath := filepath.Join(BaseDir, fileName)

	mutexDownload.Lock()
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Error retriveing file", http.StatusNotFound)
		return
	}

	if _, ok := DeleteInProgress.Load(0); ok {
		mutexDownload.Unlock()
		http.Error(w, "Delete in progress, cannot download the file", http.StatusForbidden)
		return
	}

	DeleteInProgress.Store(1, true)
	file, err := os.Open(filePath)
	if err != nil {
		mutexDownload.Unlock()
		http.Error(w, "Failed to open the file", http.StatusInternalServerError)
		return
	}
	DeleteInProgress.Store(1, false)
	mutexDownload.Unlock()
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to write the file to the response", http.StatusInternalServerError)
		return
	}
}
