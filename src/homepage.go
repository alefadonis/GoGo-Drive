package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("[GET] /")
	fmt.Fprint(w, "Welcome to Go Go Drive!\n")
}
