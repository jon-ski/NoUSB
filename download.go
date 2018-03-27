package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func download(f string, url string) (err error) {
	// Create directory if needed
	dir, _ := filepath.Split(f)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// Create the file
	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(filepath.ToSlash(url))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Success!
	return nil
}

func getRequest(url string) ([]byte, error) {
	// Request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 3/27/2018 may have broken this. need to test.
func downloadAll(url string) error {
	body, err := getRequest("/api/external/files/?url=" + url)

	var files []string
	err = json.Unmarshal(body, &files)
	if err != nil {
		return err
	}

	// Download/Create the files
	for _, file := range files {
		fmt.Println(file)
		err = download(file, httpsStrip(url)+"/download/"+file)
		if err != nil {
			return err
		}
	}

	return nil
}
