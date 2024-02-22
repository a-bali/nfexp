package nfexp

import (
	"fmt"
	"net/http"
	"strings"
)

func ServeQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	if len(r.Form["command"]) > 0 {
		x, _ := runNfdump(strings.Split(r.Form["command"][0], " ")...)
		fmt.Fprintln(w, x)
	}
}
