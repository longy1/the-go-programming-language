package main

import (
	db2 "The.Go.Programming.Language/localpkg/db"
	"log"
	"net/http"
)

func main() {
	db := db2.Database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.List))
	mux.Handle("/price", http.HandlerFunc(db.Price))
	log.Fatal(http.ListenAndServe("localhost:2021", mux))
}
