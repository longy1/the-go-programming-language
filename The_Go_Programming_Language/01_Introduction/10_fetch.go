package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		err = resp.Body.Close()
		if err != nil {
		}
		fmt.Printf("%s", data)
	}
}
