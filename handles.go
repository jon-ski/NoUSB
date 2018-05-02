package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "broken")
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

// handleAPIExternalFiles takes a GET parameter "url" and attempts to
// return the /api/files/ from another NoUSB server.
// This is to combat a cross-origin request.
func handleAPIExternalFiles(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "err: %s", err)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprintf(w, "Could not retrieve url from form")
		return
	}
	b, err := getRequest(httpsStrip(url) + "/api/files/")
	if err != nil {
		fmt.Fprintf(w, "Could not retrieve files from external server")
		return
	}
	fmt.Fprintf(w, string(b))
}

func handleAPIExternalDownloadAll(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	url := r.FormValue("url")
	basePath := r.FormValue("basePath")
	if url == "" {
		fmt.Fprintf(w, "Could not retrieve url from form")
		return
	}
	err = downloadAll(url, basePath)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "Success!")
}

func handleAPIParentFolder(w http.ResponseWriter, r *http.Request) {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	d = strings.Replace(d, "/", "\\", -1)
	ds := strings.Split(d, "\\")
	d = ds[len(ds)-1] + "\\"
	fmt.Fprintf(w, d)
}
