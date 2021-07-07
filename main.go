package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	route := chi.NewRouter()

	route.Get("/", HomeController)
	route.Get("/records-transaction", RecordsController)

	http.ListenAndServe(":3000", route)

}

func HomeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hay, Welcome"))
}

func RecordsController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get All recors"))
}
