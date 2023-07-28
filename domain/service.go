package domain

import (
	"context"
	"log"

	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/domain/storage/db"
)

type Service interface {
	GetEstablishmentTypes(context.Context) ([]string, error)
}

type service struct {
	cfg  *configs.Configs
	lgr  *log.Logger
	repo db.Repo
}

func NewService(cfg *configs.Configs, lgr *log.Logger, repo db.Repo) Service {
	return service{
		cfg:  cfg,
		lgr:  lgr,
		repo: repo,
	}
}

func (s service) GetEstablishmentTypes(ctx context.Context) ([]string, error) {
	ets, errGet := s.repo.GetEstablishmentTypes(ctx)
	if errGet != nil {
		return nil, errGet
	}

	establishmentTypes := make([]string, 0, len(ets))
	for _, et := range ets {
		establishmentTypes = append(establishmentTypes, et.Name)
	}

	return establishmentTypes, nil
}
