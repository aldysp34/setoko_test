package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/download-resi", downloadReceiptFileHandler).Methods("GET")

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	log.Println("Starting application")
	newRouter()
}
