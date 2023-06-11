package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/julienschmidt/httprouter"
)

const UploadDir = "./data"

type FileInfo struct {
	Name string `json:"name"`
	Size string  `json:"size"`
}


func handleError(err error, message string, w http.ResponseWriter, statusCode int) {
	if err != nil {
		http.Error(w , "Failed to retrieve the file", statusCode)
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	log.Println("[GET] /")
	fmt.Fprint(w, "Welcome to Go Go Drive!\n")
}

func listFiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("[GET] /files")
	files, err := ioutil.ReadDir(UploadDir)
	handleError(err, "Failed to read the directory", w, http.StatusInternalServerError)
	
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

	// Convert fileInfos to JSON
	fileInfosJSON, err := json.Marshal(fileInfos)
	handleError(err, "Failed to convert to JSON", w, http.StatusInternalServerError)
	
	// Set the response header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fileInfosJSON)
}

func uploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	log.Println("[POST] /upload")

	file, handler, err := r.FormFile("file")
	handleError(err, "Failed to retrieve the file", w, http.StatusBadRequest)
	defer file.Close()

	err = os.MkdirAll(UploadDir, os.ModePerm)
	handleError(err, "Failed to create the upload directory", w, http.StatusInternalServerError)

	filePath := filepath.Join(UploadDir, handler.Filename)

	// Create a new file on the server to save the uploaded file
	destinationFile, err := os.Create(filePath)
	handleError(err, "Failed to create the file on the server", w, http.StatusInternalServerError)
	defer destinationFile.Close()

	// Copy the uploaded file to the destination file on the server
	_, err = io.Copy(destinationFile, file)
	handleError(err, "Failed to save the uploaded file", w, http.StatusInternalServerError)
	
	// Return a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully!\n"))
}

func downloadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fileName := ps.ByName("filename")
	log.Println("[GET] /download/" + fileName)

	filePath := filepath.Join(UploadDir, fileName)

	// Check if the file exists
	_, err := os.Stat(filePath)
	handleError(err, "File not found", w, http.StatusNotFound)

	// Open the file
	file, err := os.Open(filePath)
	handleError(err, "Failed to open the file", w, http.StatusInternalServerError)
	defer file.Close()

	fileInfo, err := file.Stat()
	handleError(err, "Failed to get file information", w, http.StatusInternalServerError)

	// Set the appropriate headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Copy the file data to the response writer
	_, err = io.Copy(w, file)
	handleError(err, "Failed to write the file to the response", w, http.StatusInternalServerError)	
}

func handleRequests() {
	router := httprouter.New()
	router.GET("/", homePage)
	router.GET("/files", listFiles)
	router.POST("/upload", uploadFile)
	router.GET("/download/:filename", downloadFile)

	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main(){
	handleRequests()
}
