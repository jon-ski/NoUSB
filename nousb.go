package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
)

func main() {
	// Get path of executable
	selfExecutable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	_, selfExecutableName := filepath.Split(selfExecutable)

	// Argument handling
	aPort := flag.String("p", "8080", "Port to serve on")
	aDownload := flag.Bool("d", false, "Downloads all files from a hostlan server")
	aAddress := flag.String("a", "", "IP Address of the server to download from")
	flag.Parse()

	// Pack static in binary
	box := packr.NewBox("./assets")

	// Download
	if *aDownload {
		if *aAddress == "" {
			fmt.Println("Must assign an address")
			return
		}
		downloadAll(*aAddress, "")
		return
	}

	// Serve
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("./"))))
	// When just using /downloadself/ without the binary's name in the url
	// the resulting download was nameless. So I included a dynamic handle
	// with a redirect to ensure the download has the correct name.extension
	http.HandleFunc("/downloadself/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/downloadself/"+selfExecutableName, http.StatusMovedPermanently)
	})
	http.HandleFunc("/downloadself/"+selfExecutableName, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, selfExecutable)
	})
	http.HandleFunc("/api/files/", handleAPIFiles)
	http.HandleFunc("/api/ip/", handleAPIIP)
	http.HandleFunc("/api/external/files/", handleAPIExternalFiles)
	http.HandleFunc("/api/external/downloadall/", handleAPIExternalDownloadAll)
	http.HandleFunc("/api/parentfolder/", handleAPIParentFolder)
	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(box)))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/ui/", http.StatusMovedPermanently)
	})

	host := ":" + *aPort
	fmt.Println("Hosting on port:", *aPort)
	log.Fatal(http.ListenAndServe(host, nil))
}
