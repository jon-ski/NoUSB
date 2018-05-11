package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (s *server) routes() {
	s.router = mux.NewRouter()

	s.router.PathPrefix("/download/").Handler(http.StripPrefix("/download/", s.forceDownload(http.FileServer(http.Dir("./")))))
	s.router.HandleFunc("/downloadself/", s.handleDownloadFile(s.getExecutableName()))
	s.router.HandleFunc("/api/files/", s.handleAPIFiles())
	s.router.HandleFunc("/api/parentfolder", s.handleAPIGetParentFolder())
	s.router.HandleFunc("/zip/", s.forceDownload(s.handleZipDirectory()))
	s.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(s.uiBox)))
}

func (s *server) forceDownload(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment")
		h.ServeHTTP(w, r)
	}
}

// Meant to serve larger files
func (s *server) handleDownloadFile(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filename)
		if err != nil {
			w.Write([]byte("Error opening file"))
			return
		}
		defer file.Close()

		// Get content type of file
		fileHeader := make([]byte, 512)
		file.Read(fileHeader)
		fileContentType := http.DetectContentType(fileHeader)

		// Get file size
		fileStat, _ := file.Stat()
		fileSize := strconv.FormatInt(fileStat.Size(), 10)

		// Set headers
		w.Header().Set("Content-Disposition", "attachment; filename=%s"+filename)
		w.Header().Set("Content-Type", fileContentType)
		w.Header().Set("Content-Length", fileSize)

		// Send the file
		file.Seek(0, 0)
		io.Copy(w, file)
	}
}

func (s *server) handleAPIFiles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := fileWalk()
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(f, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

func (s *server) handleAPIGetParentFolder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(getParentFolder(), "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

func (s *server) handleZipDirectory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		zipFileName := getParentFolder() + ".zip"

		fmt.Print("Creating Archive... ")

		w.Header().Set("Content-Disposition", "attachment; filename=\""+zipFileName+"\"")

		zipWriter := zip.NewWriter(w)
		defer zipWriter.Close()

		info, err := os.Stat("./")
		if err != nil {
			fmt.Fprintf(w, "%\n", err)
			return
		}

		var baseDir string
		if info.IsDir() {
			baseDir = filepath.Base("./")
		}

		filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, "./"))
			}

			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			return err
		})

		if err != nil {
			fmt.Fprintf(w, "%\n", err)
			return
		}

		fmt.Print("Done.\n")
	}
}
