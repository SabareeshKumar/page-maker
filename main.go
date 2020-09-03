package main

import (
	"github.com/SabareeshKumar/page-maker/app"
	"log"
	"net/http"
)

const clientDir = "client/build/web"

func createComponentSkeleton(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got wrong request method %s for createComponentSkeleton",
			r.Method)
		http.NotFound(w, r)		
		return
	}
	app.CreateComponentSkeleton()
}

func main() {
	fileServer := http.FileServer(http.Dir(clientDir))
	http.Handle("/", fileServer)
	http.HandleFunc("/create_component_skeleton", createComponentSkeleton)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
