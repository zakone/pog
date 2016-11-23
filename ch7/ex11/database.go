// http://localhost:8000/create?item=sock&price=50
// http://localhost:8000/create?item=beer&price=10
// http://localhost:8000/update?item=beer&price=5.2
// http://localhost:8000/delete?item=beer
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"strconv"
)

func main() {
	db := database{}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


type dollars float32

var mu sync.Mutex

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item or price info\n")
	}
	if _, ok := db[item]; ok {
		mu.Lock()
		val, _ := strconv.ParseFloat(price, 32)
		db[item] = dollars(val)
		mu.Unlock()
		db.list(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item info\n")
	}
	if _, ok := db[item]; ok {
		mu.Lock()
		delete(db, item)
		mu.Unlock()
		db.list(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item or price info\n")
	}
	if _, ok := db[item]; !ok {
		mu.Lock()
		val, _ := strconv.ParseFloat(price, 32)
		db[item] = dollars(val)
		mu.Unlock()
		db.list(w, req)
	} else {
		w.WriteHeader(http.StatusForbidden) // 404
		fmt.Fprintf(w, "already has item: %q\n", item)
	}

}
