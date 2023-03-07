package downloader

import (
	"io"
)

type Downloader interface {
	Downloader(url string) (r io.Reader, err error)
}
