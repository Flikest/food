package storage

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	Log     *slog.Logger
	DB      *pgx.Conn
	Context context.Context
	RDB     *redis.Client
}

func InitStorage(s Storage) *Storage {
	return &Storage{
		Log:     s.Log,
		DB:      s.DB,
		Context: s.Context,
	}
}
