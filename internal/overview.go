package nfexp

import (
	"fmt"
	"net/http"
)

func ServeOverview(w http.ResponseWriter, r *http.Request) {
	x, _ := runNfdump("-g")
	w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, x)
}
