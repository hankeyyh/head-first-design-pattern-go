package main

import (
	"bytes"
	"net/http"
	"os"
)

// TODO timeout
func downloadImg(filename, url string) error {
	// Download the image from the URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.0.0")
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	// Save the image to a file
	fo, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fo.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(rsp.Body)
	if err != nil {
		return err
	}
	_, err = fo.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}