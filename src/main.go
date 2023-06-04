package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

const UploadDir = "./data"


func handleError(err error, message string, w http.ResponseWriter, statusCode int) {
	if err != nil {
		http.Error(w , "Failed to retrieve the file", statusCode)
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Welcome to Go Go Drive!\n")
}

func listFiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
}

func uploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
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

func handleRequests() {
	router := httprouter.New()
	router.GET("/", homePage)
	router.GET("/files", listFiles)
	router.POST("/upload", uploadFile)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main(){
	handleRequests()
}
