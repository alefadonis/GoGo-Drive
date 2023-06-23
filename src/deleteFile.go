package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var mutexDelete sync.Mutex

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	fileName := urlParts[len(urlParts)-1]

	DeleteInProgress.Store(0, true)

	log.Println("[DELETE] /delete/" + fileName)

	if r.Method != http.MethodDelete {
		http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
		return
	}

	if fileName == "" {
		http.Error(w, "Name of the file is empty", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(BaseDir, fileName)

	if _, ok := DeleteInProgress.Load(1); ok {
		http.Error(w, "Download in progress, cannot delete the file", http.StatusForbidden)
		return
	}

	mutexDelete.Lock()
	_, err := os.Stat(filePath)
	if err != nil {
		mutexDelete.Unlock()
		http.Error(w, fmt.Sprintf("File %s not found", fileName), http.StatusNotFound)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		mutexDelete.Unlock()
		http.Error(w, fmt.Sprintf("Error deleting the file: %v", err), http.StatusInternalServerError)
		return
	}

	DeleteInProgress.Store(0, false)
	mutexDelete.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File deleted successfully!\n"))
}
