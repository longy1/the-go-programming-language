package main

import (
	db2 "The.Go.Programming.Language/localpkg/db"
	"log"
	"net/http"
)

func main() {
	db := db2.Database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.List)
	mux.HandleFunc("/price", db.Price)
	log.Fatal(http.ListenAndServe("localhost:2021", mux))
}
