package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dustin/go-humanize"
)

func ListFiles(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET] /files")
	files, err := ioutil.ReadDir(BaseDir)

	if err != nil {
		http.Error(w, "Failed to read the directory", http.StatusInternalServerError)
		return
	}

	var fileInfos []FileInfo
	for _, file := range files {
		if file.Name() == ".gitkeep" {
			continue
		}

		fileInfos = append(fileInfos, FileInfo{
			Name: file.Name(),
			Size: humanize.Bytes(uint64(file.Size())),
		})
	}

	fileInfosJSON, err := json.Marshal(fileInfos)

	if err != nil {
		http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fileInfosJSON)
}
