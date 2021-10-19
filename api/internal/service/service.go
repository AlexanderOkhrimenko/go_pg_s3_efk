package service

import (
	"go_pg_s3_efk/domain"
	"go_pg_s3_efk/internal/repository"
)

type Rest interface {
	InsertDBurl(tab domain.EncoreTab) (uint64, error)
}

type Service struct {
	Rest
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Rest: NewRestService(repos.Rest),
	}
}
