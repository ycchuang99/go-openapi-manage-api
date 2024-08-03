package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/openapi", apiOpenapiHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func apiOpenapiHandler(w http.ResponseWriter, r *http.Request) {
	data := "do something i don't know"
	fmt.Fprintln(w, data)
}
