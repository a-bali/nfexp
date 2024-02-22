package nfexp

import (
	"log"
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
