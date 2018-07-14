package main

import (
	"io"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, Coders!")
}

func main() {
	http.Handle("/", http.HandlerFunc(handle))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
