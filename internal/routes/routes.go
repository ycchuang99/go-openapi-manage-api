package routes

import (
	"net/http"

	"github.com/ycchuang99/go-openapi-manage-api/internal/handlers"
	"github.com/ycchuang99/go-openapi-manage-api/internal/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/swagger-ui/", handlers.SwaggerUIHandler)

	// API routes with authentication
	mux.HandleFunc("/api/specs", middleware.AuthMiddleware(handlers.SpecsHandler))
	mux.HandleFunc("/api/specs/", middleware.AuthMiddleware(handlers.SpecHandler))
	mux.HandleFunc("/api/specs/validate", middleware.AuthMiddleware(handlers.ValidateSpecHandler))

	return mux
}
