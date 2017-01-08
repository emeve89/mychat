package main

import (
	"flag"
	"os"

	"github.com/emeve89/mychat/lib"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address.")
	flag.Parse()

	if isHost {
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}
}
