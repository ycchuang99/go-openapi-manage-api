package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/yourusername/go-openapi-manage-api/internal/models"
	"github.com/yourusername/go-openapi-manage-api/internal/storage"
)

func SpecsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		specs, err := storage.ListSpecs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(specs)
	case http.MethodPost:
		var spec models.OpenAPISpec
		if err := json.NewDecoder(r.Body).Decode(&spec); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := storage.CreateSpec(&spec); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(spec)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SpecHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/specs/")

	switch r.Method {
	case http.MethodGet:
		spec, err := storage.GetSpec(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(spec)
	case http.MethodPut:
		var spec models.OpenAPISpec
		if err := json.NewDecoder(r.Body).Decode(&spec); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := storage.UpdateSpec(id, &spec); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(spec)
	case http.MethodDelete:
		if err := storage.DeleteSpec(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ValidateSpecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: Implement OpenAPI validation
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Validation endpoint - Coming soon")
}
