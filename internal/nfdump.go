package nfexp

import (
	"cmp"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

var files []string

func runNfdump(params ...string) (string, error) {
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

func loadFiles(d string) {
	filepath.WalkDir(d, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if !d.IsDir() &&
			strings.HasPrefix(d.Name(), "nfcapd.") &&
			!strings.HasPrefix(d.Name(), "nfcapd.current") {
			files = append(files, s)
		}
		return nil
	})
	slices.SortFunc(files, func(a, b string) int { return cmp.Compare(a, b) })
}

func getFirstItem(arr []string, x []string, after bool) string {
	var i int
	if getFirst(x) != "" && getFirst(x) != "undefined" {
		d, _ := time.Parse("2006/01/02 15:04:05", getFirst(x))
		i = slices.IndexFunc(files, func(f string) bool {
			dd, _ := time.Parse("200601021504", f[strings.LastIndex(f, ".")+1:])
			if after {
				return dd.After(d)
			} else {
				return dd.Before(d)
			}
		})
		if i < 0 {
			return ""
		} else {
			return arr[i]
		}
	} else {
		return arr[0]
	}
}

func ServeCmd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Form data: ", r.Form)
	format := "text"

	if len(r.Form["command"]) > 0 {
		params := strings.Split(r.Form["command"][0], " ")
		if len(r.Form["format"]) > 0 && len(r.Form["format"][0]) > 0 {
			format = r.Form["format"][0]
			params = append(params, "-o", format)
		}
		if len(r.Form["filter"]) > 0 && len(r.Form["filter"][0]) > 0 {
			params = append(params, strings.Split(r.Form["filter"][0], " ")...)
		}

		if len(r.Form["save"]) > 0 && len(r.Form["save"][0]) > 0 {
			SaveReport(getFirst(r.Form["save"]),
				getFirst(r.Form["command"]),
				getFirst(r.Form["filter"]),
				getFirst(r.Form["format"]))
		}

		loadFiles(NfdumpDir)
		from := getFirstItem(files, r.Form["from"], true)
		slices.Reverse(files)
		to_file := getFirstItem(files, r.Form["to"], false)
		var to string
		if to_file != "" {
			to = filepath.Base(to_file)
		} else {
			to = ""
		}
		params = append(params, "-R", from+":"+to)

		res, _ := runNfdump(params...)

		w.Header().Set("Content-Type", "application/"+format)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, res)
	}
}
