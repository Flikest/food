package storage

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	Log     *slog.Logger
	DB      *pgx.Conn
	Context context.Context
}

func InitStorage(s Storage) *Storage {
	return &Storage{
		Log:     s.Log,
		DB:      s.DB,
		Context: s.Context,
	}
}
