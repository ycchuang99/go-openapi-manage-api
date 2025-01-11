package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/go-openapi-manage-api/internal/models"
)

var (
	specs = make(map[string]*models.OpenAPISpec)
	mu    sync.RWMutex
)

func CreateSpec(spec *models.OpenAPISpec) error {
	mu.Lock()
	defer mu.Unlock()

	spec.ID = uuid.New().String()
	spec.CreatedAt = time.Now()
	spec.UpdatedAt = spec.CreatedAt
	specs[spec.ID] = spec
	return nil
}

func GetSpec(id string) (*models.OpenAPISpec, error) {
	mu.RLock()
	defer mu.RUnlock()

	spec, ok := specs[id]
	if !ok {
		return nil, errors.New("spec not found")
	}
	return spec, nil
}

func ListSpecs() ([]*models.OpenAPISpec, error) {
	mu.RLock()
	defer mu.RUnlock()

	var result []*models.OpenAPISpec
	for _, spec := range specs {
		result = append(result, spec)
	}
	return result, nil
}

func UpdateSpec(id string, spec *models.OpenAPISpec) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := specs[id]; !ok {
		return errors.New("spec not found")
	}

	spec.ID = id
	spec.UpdatedAt = time.Now()
	specs[id] = spec
	return nil
}

func DeleteSpec(id string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := specs[id]; !ok {
		return errors.New("spec not found")
	}

	delete(specs, id)
	return nil
}
