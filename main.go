package main

import (
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/detail", detail)
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/edit", edit)
	mux.HandleFunc("/delete", delete)
	port := ":9001"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func detail(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "detail", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "create", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "edit", nil)
}

func delete(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "delete", nil)
}
