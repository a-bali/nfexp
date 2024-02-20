package nfexp

import (
	"flag"
)

var (
	NfdumpDir string
	NfdumpCmd string
	WebPort   int
	WebHost   string
	WebDir    string
)

func ParseArgs() {
	flag.StringVar(&NfdumpDir, "nfdumpdir", ".", "Directory where nfdump flows are located")
	flag.StringVar(&NfdumpCmd, "nfdumpcmd", "nfdump", "Path to nfdump binary")
	flag.IntVar(&WebPort, "port", 8080, "Port to bind for web interface")
	flag.StringVar(&WebHost, "host", "", "Host to bind for web interface")
	flag.StringVar(&WebDir, "webdir", "./web/", "Directory where static web assets are located")
	flag.Parse()
}
