package main

import (
	"fmt"
	"github.com/a-bali/nfexp/internal"
	"log"
	"net/http"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	nfexp.ParseArgs()
	http.Handle("/", nfexp.StaticFileServer())
	http.HandleFunc("/api/cmd", nfexp.ServeCmd)
	http.HandleFunc("/api/dns", nfexp.ServeDns)
	log.Printf("Launching web server on %s:%d", nfexp.WebHost, nfexp.WebPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", nfexp.WebHost, nfexp.WebPort), logRequest(http.DefaultServeMux)))
}
