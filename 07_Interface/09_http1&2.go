package main

import (
	db2 "The.Go.Programming.Language/localpkg/db"
	"log"
	"net/http"
)

func main() {
	db := db2.Database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:2021", db))
}
