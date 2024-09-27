package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type HelloHandler struct {
	env string
}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Got request %v\n", req)
	fmt.Fprintf(w, "hello from environment: \"%v\"\n", hh.env)
}

func health(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "OK")
}

func main() {
	var ENV = os.Getenv("ENV")
	if len(ENV) < 1 {
		log.Fatal("ERROR: please provide environment varialbe 'ENV'")
	}

	hh := HelloHandler{env: ENV}

	http.Handle("/hello", hh)
	http.HandleFunc("/health", health)

	const addr = ":8000"

	fmt.Printf("HTTP server listening on %v\n", addr)
	http.ListenAndServe(addr, nil)
}
