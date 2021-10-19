package service

import (
	"go_pg_s3_efk/domain"
	"go_pg_s3_efk/internal/repository"
)

type RestService struct {
	repo repository.Rest
}

func (r *RestService) InsertDBurl(tab domain.EncoreTab) (uint64, error) {
	return r.repo.InsertDBurl(tab)
}

func NewRestService(repo repository.Rest) *RestService {
	return &RestService{repo: repo}
}
