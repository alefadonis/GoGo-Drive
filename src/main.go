package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type File struct {
	Name string `json: "Name"`
}

var Files []File

func listFiles(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Files)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to Go Go Drive!\n")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/files", listFiles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func populate(){
	Files = []File{
		File{Name: "log.txt"},
		File{Name: "index.html"},
	}
}

func main(){
	populate()
	handleRequests()
}