package nfexp

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

type Report struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Filter  string `json:"filter"`
	Format  string `json:"format"`
}

type Reports map[string]Report

var reports Reports

func loadReports() {
	log.Println("Loading reports from: " + ReportFile)
	reports = make(Reports)
	_, err := toml.DecodeFile(ReportFile, &reports)
	if err != nil {
		log.Println("Unable to load reports: " + err.Error())
	}
}

func saveReports() {
	log.Println("Saving reports to: " + ReportFile)
	f, err := os.Create(ReportFile)
	if err != nil {
		log.Println(err)
	}
	if err := toml.NewEncoder(f).Encode(reports); err != nil {
		log.Println(err)
	}
	if err := f.Close(); err != nil {
		log.Println(err)
	}
}

func SaveReport(name string, command string, filter string, format string) {
	loadReports()
	log.Println("Saving report: " + name)
	reports[uuid.New().String()] = Report{
		Name:    name,
		Command: command,
		Filter:  filter,
		Format:  format,
	}
	saveReports()
}

func DeleteReport(id string) {
	loadReports()
	delete(reports, id)
	saveReports()
}

func ServeReports(w http.ResponseWriter, r *http.Request) {
	var code int
	var data []byte
	var err interface{}

	loadReports()
	r.ParseForm()
	method := r.Method

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, OPTIONS, DELETE")

	switch method {
	case "GET":
		data, err = json.Marshal(reports)
		if err != nil {
			code = http.StatusInternalServerError
		} else {
			code = http.StatusOK
		}
	case "OPTIONS":
		code = http.StatusOK
	case "DELETE":
		id := r.URL.Query().Get("uuid")
		if len(id) > 0 {
			DeleteReport(id)
			code = http.StatusOK
		} else {
			code = http.StatusInternalServerError
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}
