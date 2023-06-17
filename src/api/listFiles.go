package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dustin/go-humanize"
	"github.com/julienschmidt/httprouter"
)

func ListFiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("[GET] /files")
	files, err := ioutil.ReadDir(types.UploadDir)

	if err != nil {
		http.Error(w, "Failed to read the directory", http.StatusInternalServerError)
		return
	}

	var fileInfos []types.FileInfo
	for _, file := range files {
		if file.Name() == ".gitkeep" {
			continue
		}

		fileInfos = append(fileInfos, types.FileInfo{
			Name: file.Name(),
			Size: humanize.Bytes(uint64(file.Size())),
		})
	}

	// Convert fileInfos to JSON
	fileInfosJSON, err := json.Marshal(fileInfos)

	if err != nil {
		http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fileInfosJSON)
}
