package nfexp

import (
	"log"
	"os/exec"
)

func runNfdump(params ...string) (string, error) {
	params = append(params, "-R", NfdumpDir, "-o", "json")
	log.Println("Running nfdump with args:", params)
	cmd := exec.Command(NfdumpCmd, params...)

	output, error := cmd.Output()
	if error != nil {
		log.Println("Error running nfdump:", error)
	}
	return string(output), error
}
