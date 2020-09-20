package main

import (
	"encoding/json"
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
	if err := app.CreateComponentSkeleton(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func createForm(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		log.Printf("Got wrong request method %s for createForm",
			req.Method)
		http.NotFound(w, req)
		return
	}
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	var form app.CreateFormSchema
	if err := decoder.Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := app.CreateForm(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir(clientDir))
	http.Handle("/", fileServer)
	http.HandleFunc("/create_component_skeleton", createComponentSkeleton)
	http.HandleFunc("/create_form", createForm)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
