package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./html-files")))
	log.Fatal(http.ListenAndServe(":6090", nil))
}
