package lifecycle_repository

import (
	"database/sql"
)

// Repository implements lifecycle verifications over our database
type Repository struct {
	client *sql.DB
}

func NewRepository(client *sql.DB) Repository {
	return Repository{client: client}
}

func (r Repository) IsReady() bool {
	err := r.client.Ping()
	if err != nil {
		return false
	}
	return true
}
