package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	addr   string
	uiBox  packr.Box
}

func (s *server) start() {
	s.routes()
	fmt.Printf("Serving on %s\n", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, s.router))
}

func (s *server) getExecutableName() string {
	name, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	_, name = filepath.Split(name)
	return name
}
