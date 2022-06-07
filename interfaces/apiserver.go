package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// func handler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", req.URL.Path)
// }

func serve() {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	// http.HandleFunc("/test", handler)
	// log.Fatal(http.ListenAndServe("localhost:8000", db))
	log.Fatal(http.ListenAndServe(":8000", db))
}
