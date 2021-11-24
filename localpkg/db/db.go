package db

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var dbMutex sync.RWMutex

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

func (db Database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	// check new price
	newPriceRaw := r.URL.Query().Get("price")
	newPriceF64, err := strconv.ParseFloat(newPriceRaw, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q\n", r.URL.Query().Get("price"))
		http.Error(w, msg, http.StatusBadRequest)
	}
	// for concurrency safety
	dbMutex.Lock()
	defer dbMutex.Unlock()
	newPrice := dollars(newPriceF64)
	db[item] = newPrice
	fmt.Fprintf(w, "update %s price from %s to %s\n", item, price, newPrice)
}

func (db Database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if ok {
		msg := fmt.Sprintf("item: %q\n already exsits", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	// check new price
	newPriceRaw := r.URL.Query().Get("price")
	newPriceF64, err := strconv.ParseFloat(newPriceRaw, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q\n", r.URL.Query().Get("price"))
		http.Error(w, msg, http.StatusBadRequest)
	}
	// for concurrency safety
	dbMutex.Lock()
	defer dbMutex.Unlock()
	newPrice := dollars(newPriceF64)
	db[item] = newPrice
	fmt.Fprintf(w, "create item: %s, price: %s\n", item, newPrice)
}

func (db Database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	// for concurrency safety
	dbMutex.Lock()
	defer dbMutex.Unlock()
	delete(db, item)
	fmt.Fprintf(w, "delete item: %s, price: %s\n", item, price)
}

func StartServer() {
	db := Database{}
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	log.Fatal(http.ListenAndServe("localhost:2021", nil))
}
