package db

import (
	"fmt"
	"net/http"
)

type Database map[string]dollars

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// ServeHTTP 1.0
//func (db Database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	for item, price := range db {
//		fmt.Fprintf(w, "%s: %s\n", item, price)
//	}
//}

// ServeHTTP 2.0
func (db Database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", item, price)
	default:
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w, "no such page: %s\n", r.URL)

		msg := fmt.Sprintf("no such page: %s\n", r.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}

func (db Database) List(w http.ResponseWriter, _ *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) Price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}
