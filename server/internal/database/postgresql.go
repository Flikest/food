package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDataBase(path string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), path)
	if err != nil {
		return nil, err
	}
	return db, err
}
