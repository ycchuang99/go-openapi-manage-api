package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OpenAPI Management API")
}

func SwaggerUIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Swagger UI - Coming Soon")
}
