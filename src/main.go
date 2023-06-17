package main

import (
	"./api"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HandleRequests() {
	router := httprouter.New()
	router.GET("/", api.HomePage)
	router.GET("/files", api.ListFiles)
	router.POST("/upload", api.UploadFile)
	router.GET("/download/:filename", api.DownloadFile)
	router.DELETE("/delete/:filename", api.DeleteFile)

	log.Println("Server running at port: 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	HandleRequests()
}
