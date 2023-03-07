package main

import (
	"interface/downloader/web"
	zip "interface/zip"
	"io"
	"os"
)

func main() {

	url1 := "https://filesamples.com/samples/video/mp4/sample_1280x720_surfing_with_audio.mp4"
	url2 := "https://filesamples.com/samples/video/mp4/sample_960x400_ocean_with_audio.mp4"

	downloader := web.NewDownloader()
	zipper := zip.New()

	r1, _ := downloader.Downloader(url1)
	r2, _ := downloader.Downloader(url2)
	zipR, err := zipper.Archive([]string{"f1.mp4", "f2.mp4"}, r1, r2)
	if err != nil {
		panic(err)
	}
	zipW, err := os.Create("result.zip")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(zipW, zipR)
	if err != nil {
		panic(err)
	}

}
