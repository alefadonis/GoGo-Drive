package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const BaseDir = "/gogo-drive"

func createDir() {
	homeDir, _ := os.UserHomeDir()
	nameDir := filepath.Join(homeDir, "/gogo-drive")
	err := os.MkdirAll(nameDir, 0755)
	if err != nil {
		log.Fatal("Failed to create the upload directory")
		return
	}
	log.Println("Base Directory created successfully!")
}

func main() {
	createDir()

	http.HandleFunc("/", HomePage)

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Server running at port: 8081")
	log.Fatal(http.ListenAndServe(":8081", http.FileServer(http.Dir("/usr/share/doc/gogo-drive"))))
}
