package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto); err != nil {
		log.Print(err)
	}
	for k, v := range r.Header {
		if _, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v); err != nil {
			log.Print(err)
		}
	}
	if _, err := fmt.Fprintf(w, "Host = %q\n", r.Host); err != nil {
		log.Print(err)
	}
	if _, err := fmt.Fprintf(w, "Remote Addr = %q\n", r.RemoteAddr); err != nil {
		log.Print(err)
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if _, err := fmt.Fprintf(w, "Form[%q] = %q\n", k, v); err != nil {
			log.Print(err)
		}
	}
}
