package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET] /")
	fmt.Fprint(w, "Welcome to Go Go Drive!\n")
}
