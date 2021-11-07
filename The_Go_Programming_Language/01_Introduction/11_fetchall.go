// receive cli url args and GET response by goroutine, then counts running time
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	err = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f\t%d\t%s", secs, nbytes, url)
}
