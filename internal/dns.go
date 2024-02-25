package nfexp

import (
	"fmt"
	"net"
	"net/http"
)

func ServeDns(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip != "" {
		host, err := net.LookupAddr(ip)
		var res string
		if err == nil {
			res = host[0]
		} else {
			res = ip
		}
		w.Header().Set("Content-Type", "application/text")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, res)
	} else {
		http.Error(w, "Wrong request", 500)
		return
	}
}
