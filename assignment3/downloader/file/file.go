package file

import (
	"interface/downloader"
	"io"
	"log"
	"os"
)

type FileSystemDownloader struct{}

func NewDownloader() downloader.Downloader {
	return &FileSystemDownloader{}
}

func (d *FileSystemDownloader) Downloader(url string) (r io.Reader, err error) {
	file, err := os.Open(url)
	if err != nil {
		return nil, err
	}
	r, w := io.Pipe()
	go func() {
		defer file.Close()
		defer w.Close()
		_, err := io.Copy(w, file)
		if err != nil {
			log.Printf("Error copying file:%v", err)
		}
	}()
	return r, nil
}
