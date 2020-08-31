package main

import (
	"log"
	"net/http"
)

const clientDir = "client/build/web"

func main() {
	fileServer := http.FileServer(http.Dir(clientDir))
	log.Fatal(http.ListenAndServe(":8080", fileServer))
}
