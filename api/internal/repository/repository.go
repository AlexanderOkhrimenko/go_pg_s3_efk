package repository

import (
	"github.com/jackc/pgx/v4"
	"go_pg_s3_efk/domain"
)

type Rest interface {
	InsertDBurl(tab domain.EncoreTab) (uint64, error)
}

type Repository struct {
	Rest
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Rest: NewCrudPostgres(db),
	}
}
