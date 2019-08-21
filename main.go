package main

import (
	"log"
	"net/http"
)

func main() {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, actions!"))
	})

	addr := ":8080"

	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal(err)
	}
}
