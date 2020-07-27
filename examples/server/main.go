package main

import (
	"log"
	"net/http"

	tib "github.com/mathyourlife/arthur/tib"
)

func main() {
	log.Println("starting server")

	http.HandleFunc("/", tib.ArthurHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("http server failed: %s", err)
	}

}
