package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func fileWalk() ([]string, error) {
	var fs []string

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fs = append(fs, path)
		}
		return nil
	})

	return fs, err
}

func getParentFolder() string {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	d = strings.Replace(d, "/", "\\", -1)
	ds := strings.Split(d, "\\")
	d = ds[len(ds)-1] + "\\"
	return d
}
