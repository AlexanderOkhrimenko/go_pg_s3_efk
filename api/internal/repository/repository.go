package repository

import (
	"go_pg_s3_efk/domain"

	"github.com/jackc/pgx/v4"
)

type Rest interface {
	InsertDBurl(tab domain.EncoreTab) (uint64, error)
	InsertDBurl2(tab domain.EncoreTab) (uint64, error)
}

type Repository struct {
	Rest
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Rest: NewCrudPostgres(db),
	}
}
