package main

import (
	"log"
	"net/http"

	"github.com/joshbetz/htttp"
)

func main() {
	router := htttp.New()

	router.Get("/", func(r *http.Request, p htttp.Params) *htttp.Response {
		return htttp.Res("Hello World")
	})

	router.Get("/:name", func(r *http.Request, p htttp.Params) *htttp.Response {
		msg := "Hello " + p["name"]
		return htttp.Res(msg)
	})

	log.Fatal(http.ListenAndServe(":3000", router))
}
