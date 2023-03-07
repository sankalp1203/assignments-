package zip

import (
	"archive/zip"
	"io"
	"log"
)

type Archiver interface {
	Archive(names []string, readers ...io.Reader) (outR io.Reader, err error)
}

type Zipper struct{}

func New() Archiver {
	return &Zipper{}
}

func (z *Zipper) Archive(s []string, readers ...io.Reader) (outR io.Reader, err error) {
	r, w := io.Pipe()
	go func() {
		defer w.Close()
		zipWriter := zip.NewWriter(w)
		for i, reader := range readers {
			name := s[i]
			f, err := zipWriter.Create(name)
			if err != nil {
				log.Printf("Error creating zip entry: %v", err)
				continue
			}
			_, err = io.Copy(f, reader)
			if err != nil {
				log.Printf("Error copying file into zip: %v", err)
			}
		}
		err := zipWriter.Close()
		if err != nil {
			log.Printf("Error closing zip file:%v", err)
		}
	}()
	return r, nil
}
