package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {

	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5000", router))

}
