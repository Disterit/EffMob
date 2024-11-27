package repositroy

import "github.com/jmoiron/sqlx"

type Song interface {
}

type Repository struct {
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
