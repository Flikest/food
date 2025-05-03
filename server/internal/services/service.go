package services

import "github.com/Flikest/food/internal/storage"

type Service struct {
	Storage *storage.Storage
}

func InitService(s *storage.Storage) *Service {
	return &Service{
		Storage: s,
	}
}
