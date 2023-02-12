package main

import (
	"log"
	"net/http"
	"github.com/go-express/router"
	"github.com/go-express/env"
	"github.com/go-express/handler"
	
)


func main() {
	rtr := &router.Router{}
	env := &env.Env{}
	rtr.Get("/hello", handler.NewHandler(env, func(w http.ResponseWriter, r *http.Request) error {
		w.Write([]byte("<h1>Hello World</h1>"))
		return nil
	}))
	log.Fatal(http.ListenAndServe(":8080", rtr))
}