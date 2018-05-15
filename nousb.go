/*
NoUSB
This is a simple LAN file-share tool


Compile with packr
*/
package main

import (
	"flag"

	"github.com/gobuffalo/packr"
)

func main() {
	// Argument handling
	aPort := flag.String("p", "8080", "Port to serve on")
	flag.Parse()

	s := server{
		uiBox: packr.NewBox("./assets"),
		addr:  ":" + *aPort,
	}
	s.start()
}
