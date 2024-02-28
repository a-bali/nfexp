package nfexp

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
)

type Reports map[string]struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Filter  string `json:"filter"`
	Format  string `json:"format"`
}

var reports *Reports

func loadReports() {
	reports = new(Reports)
	_, err := toml.DecodeFile(ReportFile, &reports)
	if err != nil {
		log.Println("Unable to load reports: " + err.Error())
	}
}

func ServeReports(w http.ResponseWriter, r *http.Request) {
	var code int
	var data []byte
	var err interface{}

	loadReports()
	r.ParseForm()
	method := r.Method

	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch method {
	case "GET":
		data, err = json.Marshal(reports)
		if err != nil {
			code = http.StatusInternalServerError
		} else {
			code = http.StatusOK
		}
	case "POST":
		code = http.StatusNotFound
	case "DELETE":
		code = http.StatusNotFound
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}
