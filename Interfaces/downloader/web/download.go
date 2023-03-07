package web

import (
	"fmt"
	"interface/downloader"
	"io"
	"log"
	"net/http"
)

type WebDownloader struct{}

func NewDownloader() downloader.Downloader {
	return &WebDownloader{}
}

func (d *WebDownloader) Downloader(url string) (r io.Reader, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	r, w := io.Pipe()
	go func() {
		defer resp.Body.Close()
		defer w.Close()
		_, err := io.Copy(w, resp.Body)
		if err != nil {
			log.Printf("Error copying response %v", err)
		}
	}()
	return r, nil
}
