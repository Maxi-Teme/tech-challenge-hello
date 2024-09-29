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

func (helloHandler HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Got request %v\n", req)
	fmt.Fprintf(w, "hello from environment: \"%v\"\n", helloHandler.env)
}

func health(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "OK")
}

func main() {
	var PORT = os.Getenv("PORT")
	if len(PORT) < 1 {
		log.Fatal("ERROR: please provide environment variable 'PORT'")
	}
	var ENV = os.Getenv("ENV")
	if len(ENV) < 1 {
		log.Fatal("ERROR: please provide environment varialbe 'ENV'")
	}

	helloHandler := HelloHandler{env: ENV}

	http.Handle("/hello", helloHandler)
	http.HandleFunc("/health", health)

	var addr = fmt.Sprintf(":%v", PORT)

	fmt.Printf("HTTP server listening on %v\n", addr)
	http.ListenAndServe(addr, nil)
}
