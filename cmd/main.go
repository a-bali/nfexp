package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web/")))
	port := ":9999"
	log.Fatal(http.ListenAndServe(port, nil))
}
