package storage

import "github.com/ycchuang99/go-openapi-manage-api/internal/models"

// Storage defines the interface for spec storage operations
type Storage interface {
	CreateSpec(spec *models.OpenAPISpec) error
	GetSpec(id string) (*models.OpenAPISpec, error)
	ListSpecs() ([]*models.OpenAPISpec, error)
	UpdateSpec(id string, spec *models.OpenAPISpec) error
	DeleteSpec(id string) error
}
