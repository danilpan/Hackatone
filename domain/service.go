package domain

import (
	"context"
	"log"

	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/domain/model"
	"github.com/madxiii/hackatone/domain/storage/db"
)

type Service interface {
	GetEstablishmentTypes(context.Context) ([]string, error)
	GetEstablishments(context.Context) ([]model.Establishment, error)
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

func (s service) GetEstablishments(ctx context.Context) ([]model.Establishment, error) {
	es, errGet := s.repo.GetEstablishments(ctx)
	if errGet != nil {
		return nil, errGet
	}

	establishments := make([]model.Establishment, 0, len(es))
	for _, e := range es {
		establishments = append(establishments, model.Establishment{
			Name:         e.Name,
			Address:      e.Address,
			Type:         e.TypeName,
			AverageCheck: e.AverageCheck,
			Rating:       e.Rating,
			ImagesURLs:   e.ImagesURLs,
		})
	}

	return establishments, nil
}
