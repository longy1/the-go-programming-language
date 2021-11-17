package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}(resp.Body)

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		filename, n, err := fetch(url)
		if err != nil {
			log.Printf("%s: %v", url, err)
			continue
		}
		log.Printf("get %d bytes from %s, write to %s", n, url, filename)
	}
}
