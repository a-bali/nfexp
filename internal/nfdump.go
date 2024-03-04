package nfexp

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func runNfdump(params ...string) (string, error) {
	params = append([]string{"-R", NfdumpDir}, params...)
	log.Printf("Running nfdump with args: [%s]", strings.Join(params, ","))
	cmd := exec.Command(NfdumpCmd, params...)

	output, error := cmd.Output()
	if error != nil {
		log.Println("Error running nfdump:", error)
	}
	return string(output), error
}

func getFirst(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	} else {
		return ""
	}
}

func ServeCmd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	format := "text"

	if len(r.Form["command"]) > 0 {
		params := strings.Split(r.Form["command"][0], " ")
		if len(r.Form["format"]) > 0 && len(r.Form["format"][0]) > 0 {
			format = r.Form["format"][0]
			params = append(params, "-o", format)
		}
		if len(r.Form["filter"]) > 0 {
			params = append(params, strings.Split(r.Form["filter"][0], " ")...)
		}

		if len(r.Form["save"]) > 0 && len(r.Form["save"][0]) > 0 {
			SaveReport(getFirst(r.Form["save"]),
				getFirst(r.Form["command"]),
				getFirst(r.Form["filter"]),
				getFirst(r.Form["format"]))
		}

		res, _ := runNfdump(params...)

		w.Header().Set("Content-Type", "application/"+format)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, res)
	}
}
