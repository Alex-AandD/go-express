package main

import (
	"log"
	"net/http"
	"github.com/go-express/router"
	_"github.com/go-express/env"
	_"github.com/go-express/handler"
	"github.com/go-express/request"
	
)

/*
func LoggingMdlw(next handler) handler {
	// do some stuff
	next()
	// maybe do some other stuff
}

[handlers]



*/

func main() {
	rtr := &router.Router{}
	rtr.Get("/hello", 
	func (w http.ResponseWriter, r *request.Request, n router.NextFunction) error {
		w.Write([]byte("<h1>Hello World</h1>"))
		n()
		return nil
	},
	func(w http.ResponseWriter, r *request.Request, n router.NextFunction) error {
		w.Write([]byte("<h1>Hello World Number 2</h1>"))
		n()
		return nil
	},
)
	log.Fatal(http.ListenAndServe(":8080", rtr))
}