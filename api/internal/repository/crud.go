package repository

import (
	"context"
	"fmt"
	"go_pg_s3_efk/domain"

	"github.com/jackc/pgx/v4"
)

type CrudPostgres struct {
	db *pgx.Conn
}

func NewCrudPostgres(db *pgx.Conn) *CrudPostgres {
	return &CrudPostgres{db: db}
}

func (r *CrudPostgres) InsertDBurl(tab domain.EncoreTab) (uint64, error) {
	var lastInsertId uint64
	err := r.db.QueryRow(context.Background(), "INSERT INTO encore_tab (url) VALUES ($1) returning id;", tab.Url).Scan(&lastInsertId)
	if err != nil {
		fmt.Println(err)
	}
	return lastInsertId, err
}

func (r *CrudPostgres) InsertDBurl2(tab domain.EncoreTab) (uint64, error) {
	var lastInsertId uint64
	err := r.db.QueryRow(context.Background(), "INSERT INTO encore_tab (url) VALUES ($1) returning id;", tab.Url).Scan(&lastInsertId)
	if err != nil {
		fmt.Println(err)
	}
	return lastInsertId, err
}
