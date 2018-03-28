package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func download(f string, url string) (err error) {
	fmt.Println(f)

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

// 3/27/2018 may have broken this. need to test.
func downloadAll(url string, basePath string) error {
	// parse/modify/do stuff to basePath to make it usable
	if basePath != "" {
		basePath = strings.Replace(basePath, "/", "\\", -1)
		// basePath = strings.TrimSuffix(basePath, "\\")
		if !strings.HasSuffix(basePath, "\\") {
			basePath = basePath + "\\"
		}
	}

	// GET file list
	var files []string
	b, err := getRequest(httpsStrip(url) + "/api/files/")
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &files)
	if err != nil {
		return err
	}

	// GET parent url of server
	b, err = getRequest(httpsStrip(url) + "/api/parentfolder/")
	if err != nil {
		return err
	}
	parentFolder := string(b)

	// check if file exists so you don't overwrite something stupidly..
	for i := range files {
		files[i] = basePath + parentFolder + files[i]
		if _, err := os.Stat(files[i]); err == nil {
			return fmt.Errorf("NoUSB does not overwrite files...\n%s", files[i])
		}
	}

	// Download/Create the files
	for _, file := range files {
		err = download(file, httpsStrip(url)+"/download/"+file)
		if err != nil {
			return err
		}
	}

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
