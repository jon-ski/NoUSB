package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, bundleIndex)
}

func handleAPIFiles(w http.ResponseWriter, r *http.Request) {
	f, err := fileWalk()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(b))
}

func handleAPIIP(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Fprintf(w, "%v\n", localAddr.IP)
}
