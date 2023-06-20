package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HandleRequests() {
	router := httprouter.New()
	router.GET("/", HomePage)
	router.GET("/files", ListFiles)
	router.POST("/upload", UploadFile)
	router.GET("/download/:filename", DownloadFile)
	router.DELETE("/delete/:filename", DeleteFile)

	log.Println("Server running at port: 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	HandleRequests()
}
