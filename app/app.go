package main

import (
	"log"
	"net/http"
	"github.com/go-express/router"
)

func main() {
	router := &router.Router{}
	router.Get("/hello", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello world</h1>"))
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}