package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/go-openapi-manage-api/internal/models"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreateSpec(spec *models.OpenAPISpec) error {
	spec.ID = uuid.New().String()
	spec.CreatedAt = time.Now()
	spec.UpdatedAt = spec.CreatedAt

	_, err := s.db.Exec(`
		INSERT INTO specs (id, name, version, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, spec.ID, spec.Name, spec.Version, spec.Content, spec.CreatedAt, spec.UpdatedAt)

	return err
}

func (s *PostgresStorage) GetSpec(id string) (*models.OpenAPISpec, error) {
	spec := &models.OpenAPISpec{}
	err := s.db.QueryRow(`
		SELECT id, name, version, content, created_at, updated_at
		FROM specs WHERE id = $1
	`, id).Scan(&spec.ID, &spec.Name, &spec.Version, &spec.Content, &spec.CreatedAt, &spec.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return spec, err
}

func (s *PostgresStorage) ListSpecs() ([]*models.OpenAPISpec, error) {
	rows, err := s.db.Query(`
		SELECT id, name, version, content, created_at, updated_at
		FROM specs ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specs []*models.OpenAPISpec
	for rows.Next() {
		spec := &models.OpenAPISpec{}
		err := rows.Scan(&spec.ID, &spec.Name, &spec.Version, &spec.Content, &spec.CreatedAt, &spec.UpdatedAt)
		if err != nil {
			return nil, err
		}
		specs = append(specs, spec)
	}
	return specs, nil
}

func (s *PostgresStorage) UpdateSpec(id string, spec *models.OpenAPISpec) error {
	spec.UpdatedAt = time.Now()
	_, err := s.db.Exec(`
		UPDATE specs 
		SET name = $1, version = $2, content = $3, updated_at = $4
		WHERE id = $5
	`, spec.Name, spec.Version, spec.Content, spec.UpdatedAt, id)
	return err
}

func (s *PostgresStorage) DeleteSpec(id string) error {
	_, err := s.db.Exec("DELETE FROM specs WHERE id = $1", id)
	return err
}
