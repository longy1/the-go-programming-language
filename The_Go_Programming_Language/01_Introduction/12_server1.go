// a simple http echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
	}
}
