package main

import (
	"os"
	"path/filepath"
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
