package main

import (
	"fmt"
	"github.com/a-bali/nfexp/internal"
	"log"
	"net/http"
)

func main() {
	nfexp.ParseArgs()
	http.Handle("/", nfexp.StaticFileServer())
	http.HandleFunc("/api/overview", nfexp.ServeOverview)
	log.Printf("Launching web server on %s:%d", nfexp.WebHost, nfexp.WebPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", nfexp.WebHost, nfexp.WebPort), nil))
}
