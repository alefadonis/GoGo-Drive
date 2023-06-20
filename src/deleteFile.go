package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DeleteFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fileName := ps.ByName("filename")
	log.Println("[DELETE] /delete/" + fileName)

	if r.Method != http.MethodDelete {
		http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
		return
	}

	if fileName == "" {
		http.Error(w, "Name of the file is empty", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(UploadDir, fileName)
	_, err := os.Stat(filePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("File %s not found", fileName), http.StatusNotFound)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting the file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File deleted successfully!\n"))
}
